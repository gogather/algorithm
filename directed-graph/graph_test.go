// Copyright 2017. All rights reserved.
// This file is part of ops-backend project
// Created by duguying on 2017/11/10.

package directed_graph

import (
	"fmt"
	"testing"
)

func TestGraph_AddVertex(t *testing.T) {

	gn := NewGraphNetwork()
	gn.AddRelation("A", "D").AddRelation("A", "C").AddRelation("B", "A")
	gn.AddRelation("C", "A").AddRelation("C", "D").AddRelation("C", "E")
	gn.AddRelation("D", "G").AddRelation("E", "B").AddRelation("F", "C")
	gn.AddRelation("G", "F")

	ch, _ := gn.Done().GetChains("D", "A")
	fmt.Println(ch)
}
