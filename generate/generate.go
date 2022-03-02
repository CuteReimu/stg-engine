package generate

import (
	"github.com/CuteReimu/stg-engine/movement"
	"github.com/CuteReimu/stg-engine/utils"
)

type Generator interface {
	Generate(frame int, point utils.Point) (utils.Point, movement.Movement)
}

func Get(s string, args ...float64) Generator {
	switch s {
	case "rotate_linear":
		return NewRotateLinear(args...)
	case "norm_random_down":
		return NewNormRandomDown(args...)
	}
	return defaultGenerator
}
