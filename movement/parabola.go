package movement

import "github.com/CuteReimu/stg-engine/utils"

type parabola struct {
	speedX float64
	speedY float64
	g      float64
}

// NewParabola (speedX, speedY, g)
func NewParabola(args ...float64) Movement {
	if len(args) < 3 {
		return defaultMovement
	}
	return &parabola{
		speedX: args[0],
		speedY: args[1],
		g:      args[2],
	}
}

func (p *parabola) Move(frame int, start utils.Point) (point utils.Point, rad float64) {
	x := start.X + p.speedX*float64(frame)
	y := start.Y + p.speedY*float64(frame) + p.g*float64(frame)*float64(frame)/2
	return utils.Point{X: x, Y: y}, 0
}
