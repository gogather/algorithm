// Copyright 2019. All rights reserved.
// This file is part of gogather project
// I am coding in Tencent
// Created by rainesli on 2020/1/22.

package smooth_weighted_round_robin

import (
	"fmt"
	"testing"
)

func TestNewSmoothWeightedRoundRobin(t *testing.T) {
	n1 := NewNode("A", 10)
	n2 := NewNode("B", 10)
	n3 := NewNode("C", 59)
	swrr := NewSmoothWeightedRoundRobin([]*Node{n1, n2, n3})

	a_cnt := 0
	b_cnt := 0
	c_cnt := 0
	for i := 0; i < 1000; i++ {
		node := swrr.SelectNode()
		fmt.Println(node)
		switch node.ServerName {
		case "A":
			{
				a_cnt++
				break
			}
		case "B":
			{
				b_cnt++
				break
			}
		case "C":
			{
				c_cnt++
				break
			}
		}
	}
	fmt.Println("a:", a_cnt, "b:", b_cnt, "c:", c_cnt)
}
