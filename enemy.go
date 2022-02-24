package stg

import "github.com/CuteReimu/stg-engine/config"

type Enemy struct {
	cfg   *config.EnemyDetail
	point Point
	frame int
	HP    int
}

func NewEnemies(frame int) []*Enemy {
	var result []*Enemy
	for _, cfg := range config.Enemy.FrameMap[frame] {
		result = append(result, &Enemy{
			cfg: cfg,
			// TODO point:
			HP: cfg.Hp,
		})
	}
	return result
}

func (r *Enemy) GetPosition() Point {
	return r.point
}

func (r *Enemy) SetPosition(point Point) {
	r.point = point
}

func (r *Enemy) CheckCollide(selfPosition Point, selfRadius float64) bool {
	dx, dy := r.point.Diff(selfPosition)
	return dx*dx+dy*dy < selfRadius*selfRadius+r.cfg.Radius*r.cfg.Radius
}

func (r *Enemy) IsAlive() bool {
	if r.HP <= 0 {
		return false
	}
	return r.frame < r.cfg.DurationFrame
}
