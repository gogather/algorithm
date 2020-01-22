// Copyright 2019. All rights reserved.
// This file is part of gogather project
// I am coding in Tencent
// Created by rainesli on 2020/1/22.

package smooth_weighted_round_robin

import "github.com/gogather/json"

type Node struct {
	Weight          int    `json:"weight"`
	ServerName      string `json:"server_name"`
	currentWeight   int
	effectiveWeight int
}

func NewNode(serverName string, weight int) *Node {
	return &Node{
		Weight:        weight,
		ServerName:    serverName,
		currentWeight: weight,
	}
}

func (n *Node) String() string {
	c, _ := json.Marshal(n)
	return string(c)
}
