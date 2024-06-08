package models

import (
	"strconv"
	"strings"
)

type Reflector struct {
	connections []int
}

func GetNewReflector(connections map[string]int) Reflector {
	var r Reflector
	r.connections = make([]int, 26)
	for key, val := range connections {
		intInput, err := strconv.Atoi(string(key))
		if err != nil {
			panic(err)
		}
		r.connections[intInput] = val
		r.connections[val] = intInput
	}
	return r
}

func (r *Reflector) process(input int) int {
	return r.connections[input]
}

type PlugBoard struct {
	plugMap map[string]string
}

func (p *PlugBoard) Process(input string) string {
	return p.plugMap[input]
}

func GetNewPegboard(connections map[string]string) PlugBoard {
	validatedConnections := make(map[string]string)
	for key, val := range connections {
		validatedConnections[strings.ToUpper(key)] = strings.ToUpper(val)
	}
	return PlugBoard{
		plugMap: validatedConnections,
	}
}
