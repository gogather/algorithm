// Copyright 2017. All rights reserved.
// This file is part of ops-backend project
// Created by duguying on 2017/11/10.

package directed_graph

import (
	"encoding/json"
	"fmt"
	"strings"
)

type GraphNetwork struct {
	g         *Graph
	labels    []string
	relations [][2]string
}

func (gn *GraphNetwork) AddRelation(from string, to string) *GraphNetwork {
	gn.relations = append(gn.relations, [2]string{from, to})
	if !gn.hasLabel(from) {
		gn.labels = append(gn.labels, from)
	}
	if !gn.hasLabel(to) {
		gn.labels = append(gn.labels, to)
	}
	return gn
}

func (gn *GraphNetwork) Done() *GraphNetwork {
	gn.g = NewGraph(len(gn.labels))
	for _, label := range gn.labels {
		gn.g.AddVertex(label)
	}
	for _, relation := range gn.relations {
		from, _ := gn.getIdx(relation[0])
		to, _ := gn.getIdx(relation[1])
		gn.g.AddEdge(from, to)
	}
	return gn
}

func (gn *GraphNetwork) GetGraph() *Graph {
	return gn.g
}

func (gn *GraphNetwork) getIdx(label string) (int, error) {
	for idx, item := range gn.labels {
		if label == item {
			return idx, nil
		}
	}
	return 0, fmt.Errorf("not exist")
}

func (gn *GraphNetwork) hasLabel(label string) bool {
	exist, _ := contain(label, gn.labels)
	return exist
}

func NewGraphNetwork() *GraphNetwork {
	return &GraphNetwork{
		labels:    []string{},
		relations: [][2]string{},
	}
}

func (gn *GraphNetwork) GetChains(from string, to string) ([]Chain, error) {
	fromIdx, err := gn.getIdx(from)
	if err != nil {
		return nil, fmt.Errorf("label [%s] %s", from, err.Error())
	}
	toIdx, err := gn.getIdx(to)
	if err != nil {
		return nil, fmt.Errorf("label [%s] %s", to, err.Error())
	}
	af := NewAf(gn.g, fromIdx, toIdx)
	_, chains := af.GetResult()
	return gn.loadAsChains(chains), nil
}

func (gn *GraphNetwork) loadAsChains(rawChains [][]string) []Chain {
	var ch []Chain
	for _, chain := range rawChains {
		ch = append(ch, Chain(chain))
	}
	return ch
}

type ChainWeight []string

func (wc *ChainWeight) String() string {
	return strings.Join([]string(*wc), "")
}

type Chain []string

func (c *Chain) String() string {
	content, _ := json.Marshal([]string(*c))
	return string(content)
}

func (c *Chain) LoadWeight(weights []int) (*ChainWeight, error) {
	var ch []string
	var chain = []string(*c)
	if len(weights) < len(chain)-1 {
		return nil, fmt.Errorf("invalid weight number")
	}
	for idx, node := range chain {
		if len(chain)-1 == idx {
			ch = append(ch, node)
		} else {
			ch = append(ch, node, fmt.Sprintf("{%d}", weights[idx]))
		}
	}
	cw := ChainWeight(ch)
	return &cw, nil
}
