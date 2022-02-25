package movement

import (
	"github.com/CuteReimu/stg-engine/utils"
)

type Movement interface {
	Move(frame int, start utils.Point) (point utils.Point, rad float64)
}

func Get(s string, args ...float64) Movement {
	switch s {
	case "stay":
		return NewStay(args...)
	case "linear":
		return NewLinear(args...)
	}
	return defaultMovement
}

type stay struct {
	rad float64
}

// NewStay (rad)
func NewStay(args ...float64) Movement {
	if len(args) == 0 {
		return defaultMovement
	}
	return &stay{args[0]}
}

func (s *stay) Move(_ int, start utils.Point) (point utils.Point, rad float64) {
	return start, s.rad
}

var defaultMovement = &stay{}
