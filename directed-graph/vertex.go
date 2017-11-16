// Copyright 2017. All rights reserved.
// This file is part of ops-backend project
// Created by duguying on 2017/11/10.

package directed_graph

type Vertex struct {
	wasVisited     bool  // 是否遍历过
	Label          string // 节点名称
	AllVisitedList *ArrayList
}

func (v *Vertex) SetAllVisitedList(allVisitedList *ArrayList) {
	v.AllVisitedList = allVisitedList
}

func (v *Vertex) GetAllVisitedList() *ArrayList {
	return v.AllVisitedList
}

func (v *Vertex) GetWasVisited() bool {
	return v.wasVisited
}

func (v *Vertex) SetWasVisited(wasVisited bool) {
	v.wasVisited = wasVisited
}

func (v *Vertex) GetLabel() string {
	return v.Label
}

func (v *Vertex) SetLabel(label string) {
	v.Label = label
}

func (v *Vertex) SetVisited(j int) {
	l := *v.AllVisitedList
	l.Set(j, 1)
}

func NewVertex(lab string) *Vertex {
	return &Vertex{
		Label:          lab,
		wasVisited:     false,
		AllVisitedList: NewArrayList(),
	}
}
