// Copyright 2019. All rights reserved.
// This file is part of gogather project
// I am coding in Tencent
// Created by rainesli on 2020/1/22.

package smooth_weighted_round_robin

type SmoothWeightedRoundRobin struct {
	nodeList []*Node
}

func NewSmoothWeightedRoundRobin(nodes []*Node) *SmoothWeightedRoundRobin {
	swrr := &SmoothWeightedRoundRobin{}
	for _, node := range nodes {
		swrr.nodeList = append(swrr.nodeList, node)
	}

	return swrr
}

func (swrr *SmoothWeightedRoundRobin) SelectNode() (best *Node) {
	total := 0
	for i := 0; i < len(swrr.nodeList); i++ {
		w := swrr.nodeList[i]
		if w == nil {
			continue
		}
		w.currentWeight += w.effectiveWeight
		total += w.effectiveWeight
		if w.effectiveWeight < w.Weight {
			w.effectiveWeight++
		}
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return nil
	}
	best.currentWeight -= total
	return best
}
