package parp

import (
	"fmt"
	"strconv"
)

type Param struct {
	value string
}

func (p Param) Bool() bool {
	if p.value == "true" {
		return true
	} else if p.value == "false" {
		return false
	}
	panic(fmt.Sprintf("%s cannot be casted to type <bool>\n", p.value))
}

func (p Param) Int() int {
	v, err := strconv.Atoi(p.value)
	if err != nil {
		panic(fmt.Sprintf("%s cannot be casted to type <int>\n", p.value))
	}
	return v
}

func (p Param) String() string {
	return p.value
}

func contains(A []string, e string) bool {
	for _, a := range A {
		if a == e {
			return true
		}
	}
	return false
}

func Process(params []string, mapKeys []string, singleKeys []string) map[string]Param {
	pmap := make(map[string]Param)
	for i := 0; i < len(params); i++ {
		p := params[i]
		if contains(singleKeys, p) {
			pmap[p] = Param{value: "true"}
		} else if contains(mapKeys, p) {
			i++
			pValue := params[i]
			pmap[p] = Param{value: pValue}
		} else {
			panic(fmt.Sprintf("Unknown command parameter : %s", p))
		}
	}
	return pmap
}
