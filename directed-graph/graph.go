// Copyright 2017. All rights reserved.
// This file is part of ops-backend project
// Created by duguying on 2017/11/10.

package directed_graph

import (
	"fmt"
)

type Graph struct {
	vertexList []*Vertex
	adjMatrix  [][]int

	nVerts   int
	vertsNum int

	i int
	j int
}

func (g *Graph) String() string {
	c := "        "
	for _, vertex := range g.vertexList {
		c = c + fmt.Sprintf("[%5s] ", vertex.Label)
	}
	c = c + "\n"

	for i, line := range g.adjMatrix {
		c = c + fmt.Sprintf("[%5s] ", g.vertexList[i].Label)
		for _, cell := range line {
			c = c + fmt.Sprintf("%7d ", cell)
		}
		c = c + "\n"
	}
	return c
}

func (g *Graph) GetVertexList() []*Vertex {
	return g.vertexList
}

func (g *Graph) GetVertexListLabel() []string {
	var labels []string
	for _, vertex := range g.vertexList {
		labels = append(labels, vertex.Label)
	}
	return labels
}

func (g *Graph) GetAdjMatrix() [][]int {
	return g.adjMatrix
}

func (g *Graph) GetN() int {
	return g.vertsNum
}

func (g *Graph) delEdge(start int, end int) {
	g.adjMatrix[start][end] = 0
}

func (g *Graph) AddEdge(start int, end int) {
	g.adjMatrix[start][end] = 1
}

func (g *Graph) AddVertex(lab string) {
	g.vertexList[g.nVerts] = NewVertex(lab)
	g.nVerts++
}

func (g *Graph) DisplayVertex(i int) string {
	return g.vertexList[i].GetLabel()
}

func (g *Graph) DisplayVertexVisited(i int) bool {
	return g.vertexList[i].GetWasVisited()
}

func (g *Graph) PrintGraph() {
	for i := 0; i < g.vertsNum; i++ {
		fmt.Printf("第 %s 个节点: ", g.DisplayVertex(i))

		for j := 0; j < g.vertsNum; j++ {
			fmt.Printf("[%s-%s]: %3d ", g.DisplayVertex(i), g.DisplayVertex(j), g.adjMatrix[i][j])
		}
		fmt.Println()
	}
}

func NewGraph(vertsNum int) *Graph {
	g := &Graph{
		vertsNum: vertsNum,
	}

	g.initMatrix()
	g.initVertexList()
	g.nVerts = 0

	return g
}

func (g *Graph) initVertexList() {
	g.vertexList = make([]*Vertex, g.vertsNum)
}

func (g *Graph) initMatrix() {
	for i := 0; i < g.vertsNum; i++ {
		sl := make([]int, 0, g.vertsNum)
		for j := 0; j < g.vertsNum; j++ {
			sl = append(sl, 0)
		}
		g.adjMatrix = append(g.adjMatrix, sl)
	}
}
