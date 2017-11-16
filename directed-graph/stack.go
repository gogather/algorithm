// Copyright 2017. All rights reserved.
// This file is part of ops-backend project
// Created by duguying on 2017/11/10.

package directed_graph

import (
	"container/list"
)

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	l := list.New()
	return &Stack{l}
}

func (stack *Stack) Push(value interface{}) {
	stack.l.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.l.Back()
	if e != nil {
		stack.l.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.l.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.l.Len()
}

func (stack *Stack) Empty() bool {
	return stack.l.Len() == 0
}

func (stack *Stack) Contains(e interface{}) bool {
	for ele := stack.l.Back(); ele != nil; {
		if ele.Value == e {
			return true
		}
		ele = ele.Prev()
	}
	return false
}

func (stack *Stack) AsChain() []interface{} {
	var chain []interface{}
	for ele := stack.l.Front(); ele != nil; {
		chain = append(chain, ele.Value)
		ele = ele.Next()
	}
	return chain
}
