// Copyright 2017. All rights reserved.
// This file is part of ops-backend project
// Created by duguying on 2017/11/10.

package directed_graph

import "fmt"

type Af struct {
	isAf  bool
	g     *Graph
	n     int
	start int
	end   int

	finalChains [][]string
	theStack    *Stack

	tempList       *ArrayList
	counterExample string
}

func (a *Af) GetResult() (af bool, chains [][]string) {
	a.n = a.g.GetN()
	a.theStack = NewStack()

	if !a.isConnectable(a.start, a.end) {
		a.isAf = false
		a.counterExample = "节点之间没有通路"
	} else {
		for j := 0; j < a.n; j++ {
			a.tempList = NewArrayList()
			for i := 0; i < a.n; i++ {
				a.tempList.Add(0)
			}
			a.g.GetVertexList()[j].SetAllVisitedList(a.tempList)
		}

		a.isAf = a.af(a.start, a.end)
	}

	return a.isAf, a.finalChains
}

func (a *Af) af(start int, end int) bool {
	a.g.GetVertexList()[start].SetWasVisited(true)
	a.theStack.Push(start)

	for !a.theStack.Empty() {
		v := a.getAdjUnvisitedVertex(a.theStack.Peak().(int))
		if v == -1 {
			a.tempList = NewArrayList()
			for j := 0; j < a.n; j++ {
				a.tempList.Add(0)
			}
			a.g.GetVertexList()[a.theStack.Peak().(int)].SetAllVisitedList(a.tempList)
			a.theStack.Pop()
		} else {
			a.theStack.Push(v)
		}

		if !a.theStack.Empty() && end == a.theStack.Peak() {
			a.g.GetVertexList()[end].SetWasVisited(false)
			a.finalChains = append(a.finalChains, a.asChain())
			a.theStack.Pop()
		}
	}

	return a.isAf
}

func (a *Af) isConnectable(start int, end int) bool {
	queue := NewArrayList()
	visited := NewArrayList()
	queue.Add(start)
	for !queue.IsEmpty() {
		for j := 0; j < a.n; j++ {
			if a.g.GetAdjMatrix()[start][j] == 1 && !visited.Contains(j) {
				queue.Add(j)
			}
		}
		if queue.Contains(end) {
			return true
		} else {
			visited.Add(queue.Get(0))
			queue.Remove(0)
			if !queue.IsEmpty() {
				start = queue.Get(0).(int)
			}
		}
	}
	return false
}

func (a *Af) getAdjUnvisitedVertex(v int) int {
	al := a.g.GetVertexList()[v].GetAllVisitedList()
	for j := 0; j < a.n; j++ {
		if a.g.GetAdjMatrix()[v][j] == 1 && al.Get(j).(int) == 0 && !a.theStack.Contains(j) {
			a.g.GetVertexList()[v].SetVisited(j)
			return j
		}
	}
	return -1
}

func (a *Af) printTheStack() {
	chain := a.theStack.AsChain()
	for idx, value := range chain {
		fmt.Printf("%s", a.g.GetVertexList()[value.(int)].Label)
		if len(chain)-1 != idx {
			fmt.Printf("-->")
		}
	}
	fmt.Println()
}

func (a *Af) asChain() []string {
	var c []string
	chain := a.theStack.AsChain()
	for _, value := range chain {
		c = append(c, a.g.GetVertexList()[value.(int)].Label)
	}

	return c
}

func NewAf(graph *Graph, start int, end int) *Af {
	return &Af{
		g:     graph,
		start: start,
		end:   end,
	}
}
