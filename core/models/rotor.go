package models

import (
	"errors"
	"strconv"

	"main/core/utils"
)

type Direction string

const (
	totalAlphaCount      = 26
	DirectionRightToLeft = "RL"
	DirectionLeftToRight = "LR"
)

type Rotor struct {
	position int
	connLR   []int
	connRL   []int
}

// incrementPosition, increments the rotor position with 1 and return true if the position is reset back to 0
func (r *Rotor) increment() bool {
	r.position++
	if r.position == totalAlphaCount {
		// rotor completed full circle
		r.position = 0
		return true
	}
	return false
}

func (r *Rotor) crunch(inputIndex int, direction Direction) int {
	index := (inputIndex + r.position) % totalAlphaCount
	switch direction {
	case DirectionRightToLeft:
		return utils.Modulo(r.connRL[index]-r.position, totalAlphaCount)
	case DirectionLeftToRight:
		return utils.Modulo(r.connLR[index]-r.position, totalAlphaCount)
	default:
		panic("invalid rotor direction")
	}
}

func GetNewRotor(connections map[string]int, pos int) (Rotor, error) {
	connRL := make([]int, totalAlphaCount)
	connLR := make([]int, totalAlphaCount)
	for lString, r := range connections {
		l, err := strconv.Atoi(lString)
		if err != nil {
			return Rotor{}, errors.New("bad configuration for rotor Left Pin is not a Number, val:" + string(lString))
		}
		connLR[r] = l
		connRL[l] = r
	}
	return Rotor{
		position: pos,
		connLR:   connLR,
		connRL:   connRL,
	}, nil
}
