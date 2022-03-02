package generate

import (
	"github.com/CuteReimu/stg-engine/movement"
	"github.com/CuteReimu/stg-engine/utils"
)

type rotateLinear struct {
	rad   float64
	alpha float64
	beta  float64
	speed float64
}

// NewRotateLinear (rad, alpha, beta, speed)
func NewRotateLinear(args ...float64) Generator {
	if len(args) < 4 {
		return defaultMovement
	}
	return &rotateLinear{
		rad:   args[0],
		alpha: args[1],
		beta:  args[2],
		speed: args[3],
	}
}

func (s *rotateLinear) Generate(frame int, point utils.Point) (utils.Point, movement.Movement) {
	rad := s.rad + (s.alpha+s.beta/2*float64(frame))*float64(frame)
	return point, movement.NewLinear(0, 0, rad, s.speed)
}

var defaultMovement = &rotateLinear{}
