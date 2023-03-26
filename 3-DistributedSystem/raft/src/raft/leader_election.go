package raft

import (
	"math/rand"
	"sync"
	"time"
)

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	term := rf.currentTerm
	isleader := rf.state == Leader
	return term, isleader
}

func (rf *Raft) resetElectionTimer() {
	t := time.Now()
	electionTimeout := time.Duration(150+rand.Intn(150)) * time.Millisecond
	rf.electionTime = t.Add(electionTimeout)
}

func (rf *Raft) setNewTerm(term int) {
	if term > rf.currentTerm || rf.currentTerm == 0 {
		rf.state = Follower
		rf.currentTerm = term
		rf.votedFor = -1
		DPrintf("[%d]: set term %v\n", rf.me, rf.currentTerm)
		rf.persist()
	}
}

// Part3：leader election
// ticker会以心跳为周期不断检查状态。如果当前是Leader就会发送心跳包，而心跳包是靠appendEntries()发送空log，而不是额外的函数，这一点在论文和student guide都有强调。
// 如果发现选举超时，这时候就会出发新一轮leader election。先看leader election的实现
// 启动新一轮leader election时，首先要将自己转为candidate状态，并且给自己投一票。然后向所有peer请求投票
// 除了要满足论文的Figure 2中*Rules for Servers*的要求之外，要注意当candidate收到半数以上投票之后就可以进入leader状态，而这个状态转变会更新nextIndex[]和matchIndex[]，并且再成为leader之后要立刻发送一次心跳。我们希望状态转变只发生一次，这里我使用了go的sync.Once。简单的使用bool flag也同样可以达成目的，只不过可读性没有这么直观
func (rf *Raft) leaderElection() {
	rf.currentTerm++
	rf.state = Candidate
	rf.votedFor = rf.me
	rf.persist()
	rf.resetElectionTimer()
	term := rf.currentTerm
	voteCounter := 1
	lastLog := rf.log.lastLog()
	DPrintf("[%v]: start leader election, term %d\n", rf.me, rf.currentTerm)
	args := RequestVoteArgs{
		Term:         term,
		CandidateId:  rf.me,
		LastLogIndex: lastLog.Index,
		LastLogTerm:  lastLog.Term,
	}

	var becomeLeader sync.Once
	for serverId, _ := range rf.peers {
		if serverId != rf.me {
			go rf.candidateRequestVote(serverId, &args, &voteCounter, &becomeLeader)
		}
	}
}
