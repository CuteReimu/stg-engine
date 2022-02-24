package stg

import (
	"github.com/CuteReimu/stg-engine/config"
	"github.com/CuteReimu/stg-engine/movement"
	"github.com/CuteReimu/stg-engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	cfg   *config.EnemyDetail
	start utils.Point
	point utils.Point
	rad   float64
	frame int
	HP    int
	move  movement.Movement
	pic   *ebiten.Image
}

func NewEnemies(frame int) []*Enemy {
	var result []*Enemy
	for _, cfg := range config.Enemy.FrameMap[frame] {
		result = append(result, &Enemy{
			cfg:   cfg,
			start: utils.Point{X: cfg.StartX, Y: cfg.StartY},
			point: utils.Point{X: cfg.StartX, Y: cfg.StartY},
			HP:    cfg.Hp,
			move:  movement.Get(cfg.Move, cfg.MoveP1, cfg.MoveP2, cfg.MoveP3, cfg.MoveP4, cfg.MoveP5, cfg.MoveP6),
			pic:   GetPic(cfg.Pic),
		})
	}
	return result
}

func (r *Enemy) GetPosition() utils.Point {
	return r.point
}

func (r *Enemy) SetPosition(point utils.Point) {
	r.point = point
}

func (r *Enemy) CheckCollide(selfPosition utils.Point, selfRadius float64) bool {
	dx, dy := r.point.Diff(selfPosition)
	return dx*dx+dy*dy < selfRadius*selfRadius+r.cfg.Radius*r.cfg.Radius
}

func (r *Enemy) IsAlive() bool {
	if r.HP <= 0 {
		return false
	}
	return r.frame < r.cfg.DurationFrame
}

func (r *Enemy) Update() error {
	r.point, r.rad = r.move.Move(r.frame, r.start)
	r.frame++
	return nil
}

func (r *Enemy) Draw(screen *ebiten.Image) {
	if r.pic == nil {
		return
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-r.cfg.PicX, -r.cfg.PicY)
	opt.GeoM.Scale(r.cfg.PicScale, r.cfg.PicScale)
	opt.GeoM.Translate(r.point.X, r.point.Y)
	screen.DrawImage(r.pic, opt)
}
