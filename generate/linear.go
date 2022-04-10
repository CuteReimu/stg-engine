package generate

import (
	"github.com/CuteReimu/stg-engine/global"
	"github.com/CuteReimu/stg-engine/movement"
	"github.com/CuteReimu/stg-engine/utils"
	"math"
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
		return defaultGenerator
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

var defaultGenerator = &rotateLinear{}

type normRandomDown struct {
	normRadius float64
	downSpeed  float64
}

// NewNormRandomDown (normRadius, downSpeed)
func NewNormRandomDown(args ...float64) Generator {
	if len(args) < 2 {
		return defaultGenerator
	}
	return &normRandomDown{
		normRadius: args[0],
		downSpeed:  args[1],
	}
}

func (n *normRandomDown) Generate(_ int, point utils.Point) (utils.Point, movement.Movement) {
	x := utils.Rand.NormFloat64() * n.normRadius
	y := utils.Rand.NormFloat64() * n.normRadius
	return utils.Point{X: x + point.X, Y: y + point.Y}, movement.NewLinear(0, n.downSpeed, 0, 0)
}

type selfTarget struct {
	speed float64
}

// NewSelfTarget (speed)
func NewSelfTarget(args ...float64) Generator {
	if len(args) < 1 {
		return defaultGenerator
	}
	return &selfTarget{speed: args[0]}
}

func (s *selfTarget) Generate(_ int, point utils.Point) (utils.Point, movement.Movement) {
	diffX, diffY := global.SelfPoint.Diff(point)
	rad := math.Atan2(diffY, diffX)
	return point, movement.NewLinear(0, 0, rad/math.Pi*180, s.speed)
}
