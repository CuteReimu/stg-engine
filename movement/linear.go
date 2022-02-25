package movement

import (
	"github.com/CuteReimu/stg-engine/utils"
	"math"
)

type linear struct {
	speedX, speedY float64
	rad            float64
}

// NewLinear (speedX, speedY, rad, speed)
func NewLinear(args ...float64) Movement {
	if len(args) < 4 {
		return defaultMovement
	}
	y, x := math.Sincos(args[2] / 180 * math.Pi)
	return &linear{
		speedX: args[0] + x*args[3],
		speedY: args[1] + y*args[3],
		rad:    math.Atan2(y+args[1], x+args[2]),
	}
}

func (l *linear) Move(frame int, start utils.Point) (point utils.Point, rad float64) {
	return utils.Point{X: start.X + l.speedX*float64(frame), Y: start.Y + l.speedY*float64(frame)}, l.rad
}
