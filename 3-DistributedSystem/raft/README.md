# Raft

## Raft的结构
初始化Raft的时候，除了给raft做基本的赋值之外，还要额外启动两个goroutine
ticker函数，作用也很简单，计时并且按时间触发leader election或者append entry
而applier则是负责将command应用到state machine

## ticker
1. ticker会以心跳为周期不断检查状态。如果当前是Leader就会发送心跳包，而心跳包是靠appendEntries()发送空log，而不是额外的函数，这一点在论文和student guide都有强调。
2. 如果发现选举超时，这时候就会出发新一轮leader election

## leader election
1. 启动新一轮leader election时，首先要将自己转为candidate状态，并且给自己投一票。然后向所有peer请求投票。RequestVote RPC的参数和返回值需要按照Figure 2实现。
2. 除了要满足论文的Figure 2中*Rules for Servers*的要求之外，要注意当candidate收到半数以上投票之后就可以进入leader状态，而这个状态转变会更新nextIndex[]和matchIndex[]，并且再成为leader之后要立刻发送一次心跳。我们希望状态转变只发生一次，这里我使用了go的sync.Once。简单的使用bool flag也同样可以达成目的，只不过可读性没有这么直观。

## RequestVote
1. 另一方面，任何服务器收到RequestVote RPC之后，要实现Figure 2中*RequestVote RPC Receiver implementation*的逻辑，同时也要满足*Rules for Servers*

## AppendEntry
1. 完成了leader election之后，leader会立刻触发一次心跳包，随后在每个心跳周期发送心跳包，来阻止新一轮leader election。
Figure 2中*Rules for Servers*的*Leaders*部分将心跳称为`initial empty AppendEntries RPCs (heartbeat)`，将包含log的RPC称为`AppendEntries RPC with log entries starting at nextIndex`。这种描述听起来像是用了两段不同的代码。
2. 而实际上因为这里的心跳有两种理解：每个心跳周期，发送一次AppendEntries RPC，当这个RPC不包含log时，这个包被称为心跳包。所以也有可能发生这么一种情况：触发了一次心跳，但是带有log（即心跳周期到了，触发了一次AppendEntries RPC，但是由于follower落后了，所以这个RPC带有一段log，此时这个包就不能称为心跳包）。
3. 实践中，我在每个心跳周期和收到新的command之后各会触发一次AppendEntries RPC。然而仔细读论文后发现，论文中并没有只说了心跳会触发AppendEntries RPC，并没有说收到客户端的指令之后应该触发AppendEntries RPC。
4. 我甚至认为在理论上AppendEntries可以完全交给heartbeat周期来触发，即收到command后，并不立刻发送AppendEntries，而是等待下一个心跳。这种方法可以减少RPC的数量，并且通过了连续1000次测试。但是代价就是每条command的提交周期变长。
5. Leader在AppendEntries中会并行地给所有server发送消息，然后根据返回的消息更新nextIndex和matchIndex，这部分需要按照论文5.3节来实现。
6. fast rollback优化:实现这种优化需要在返回消息中额外加入XTerm, XIndex, XLen三个字段。
7. 当peer收到AppendEntry RPC的时候，需要实现Figure 2中*AppendEntry RPC Receiver implementation* + *Rules for Servers*。具体哪些相关，我已经加在注释里了。论文里的步骤必须严格遵守，不要自由发挥。这一点想必大家在debug的时候都深有体会……
8. 完成AppendEntry RPC之后，Leader需要提交已有的日志条目，这一点在论文5.3 & 5.4有文字叙述。但是具体在什么时候提交，需要自己去把握。仔细看Figure 2的话，其实这部分对应*Rules for Servers*中Leader部分的最后一小段

## applier
1. student guide中提到应该使用一个`a dedicated “applier”`来专门处理日志commit的事情。所以按TA说的来，并且按照作业要求使用applyCond。这里可能会触发student guide所说的`The four-way deadlock`

## Start
1. 最后是start函数，它会接受客户端的command，并且应用raft算法。前面也说了，每次start并不一定要立刻触发AppendEntry。理论上如果每次都触发AppendEntry，而start被调用的频率又超高，Leader就会疯狂发送RPC。如果不主动触发，而被动的依赖心跳周期，反而可以造成batch operation的效果，将QPS固定成一个相对较小的值。当中的trade-off需要根据使用场景自己衡量。

