package stg

import (
	"github.com/CuteReimu/stg-engine/config"
	"github.com/CuteReimu/stg-engine/generate"
	"github.com/CuteReimu/stg-engine/movement"
	"github.com/CuteReimu/stg-engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Generator struct {
	gen generate.Generator
	cfg *config.GenerateDetail
}

type Enemy struct {
	cfg   *config.EnemyDetail
	start utils.Point
	point utils.Point
	rad   float64
	frame int
	HP    int
	move  movement.Movement
	pic   *ebiten.Image
	gen   []*Generator
}

func NewEnemies(frame int) []*Enemy {
	var result []*Enemy
	for _, cfg := range config.Enemy.FrameMap[frame] {
		e := &Enemy{
			cfg:   cfg,
			start: utils.Point{X: cfg.StartX, Y: cfg.StartY},
			point: utils.Point{X: cfg.StartX, Y: cfg.StartY},
			HP:    cfg.Hp,
			move:  movement.Get(cfg.Move, cfg.MoveP1, cfg.MoveP2, cfg.MoveP3, cfg.MoveP4, cfg.MoveP5, cfg.MoveP6),
			pic:   GetPic(cfg.Pic),
		}
		for _, genId := range cfg.Generate {
			cfg := config.Generate[genId]
			if cfg == nil {
				log.Printf("cannot find generate, id: %d\n", genId)
				continue
			}
			gen := generate.Get(cfg.Move, cfg.MoveP1, cfg.MoveP2, cfg.MoveP3, cfg.MoveP4, cfg.MoveP5, cfg.MoveP6)
			e.gen = append(e.gen, &Generator{gen: gen, cfg: cfg})
		}
		result = append(result, e)
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

func (r *Enemy) Update() (bullets []*Bullet, err error) {
	r.point, r.rad = r.move.Move(r.frame, r.start)
	for _, gen := range r.gen {
		cfg := gen.cfg
		frame := r.frame - cfg.StartFrame
		if frame >= 0 && frame%cfg.IntervalFrame == 0 && frame <= cfg.DurationFrame {
			point, move := gen.gen.Generate(frame, r.point)
			bullets = append(bullets, NewBullet(cfg.Bullet, point, move, cfg.DurationFrame))
		}
	}
	r.frame++
	return
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
