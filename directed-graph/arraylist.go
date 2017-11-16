// Copyright 2017. All rights reserved.
// This file is part of ops-backend project
// Created by duguying on 2017/11/10.

package directed_graph

import (
	"errors"
	"reflect"
)

type ArrayList struct {
	l []interface{}
}

func (al *ArrayList) Add(e interface{}) {
	al.l = append(al.l, e)
}

func (al *ArrayList) Get(i int) interface{} {
	return al.l[i]
}

func (al *ArrayList) Set(i int, e interface{}) {
	al.l[i] = e
}

func (al *ArrayList) IsEmpty() bool {
	return 0 == len(al.l)
}

func (al *ArrayList) Clear() {
	al.l = []interface{}{}
}

func (al *ArrayList) Remove(i int) {
	l := append(al.l[:i], al.l[i+1:]...)
	al.l = l
}

func (al *ArrayList) Contains(e interface{}) bool {
	c, _ := contain(e, al.l)
	return c
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

func contain(e interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == e {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(e)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
