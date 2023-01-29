package main

import (
	"fmt"
	"sort"
)

type ServerSlice []Server
type Server struct {
	Weight int
}

//只要实现了下面三个方法就可以传入sort方法使用
func (p ServerSlice) Len() int           { return len(p) }
func (p ServerSlice) Less(i, j int) bool { return p[i].Weight < p[j].Weight }
func (p ServerSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	ss := ServerSlice{
		Server{Weight: 4},
		Server{Weight: 3},
		Server{Weight: 6},
	}
	sort.Sort(ss)
	for _, v := range ss {
		fmt.Println(v.Weight)
	}
}
