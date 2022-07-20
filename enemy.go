package stg

import (
	"github.com/CuteReimu/stg-engine/config"
	"github.com/CuteReimu/stg-engine/global"
	"github.com/CuteReimu/stg-engine/utils"
	"github.com/dengsgo/math-engine/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"strconv"
	"strings"
)

type Enemy struct {
	cfg      *config.EnemyDetail
	start    utils.Point
	point    utils.Point
	rad      float64
	frame    int
	HP       int
	pic      *ebiten.Image
	parent   *Enemy
	children []*Enemy
}

func NewEnemies() *Enemy {
	return &Enemy{
		cfg: &config.EnemyDetail{
			DurationFrame: 99999999,
			Generate: func() []int {
				var ids []int
				for _, e := range config.Enemy {
					if e.Type == 1 {
						ids = append(ids, e.Id)
					}
				}
				return ids
			}(),
		},
		start: utils.Point{},
		point: utils.Point{},
		HP:    99999999,
	}
}

func (r *Enemy) Update() {
	var newX, newY float64
	if len(r.cfg.X) != 0 {
		newX = r.calculate(r.cfg.X)
	}
	if len(r.cfg.Y) != 0 {
		newY = r.calculate(r.cfg.Y)
	}
	if len(r.cfg.Rad) != 0 {
		r.rad = r.calculate(r.cfg.Rad)
	} else if newY-r.point.Y != 0 || newX-r.point.X != 0 {
		r.rad = math.Atan2(newY-r.point.Y, newX-r.point.X)
	}
	r.point.X = newX
	r.point.Y = newY
	for _, cfgId := range r.cfg.Generate {
		cfg := config.Enemy[cfgId]
		c1 := r.cfg.IntervalFrame == 0 && r.frame-r.cfg.StartFrame == 0
		c2 := r.cfg.IntervalFrame != 0 && (r.frame-r.cfg.StartFrame)%r.cfg.IntervalFrame == 0
		if c1 || c2 {
			e := &Enemy{
				cfg:    cfg,
				start:  utils.Point{X: cfg.StartX, Y: cfg.StartY},
				point:  utils.Point{X: cfg.StartX, Y: cfg.StartY},
				HP:     cfg.Hp,
				pic:    GetPic(cfg.Pic),
				parent: r,
			}
			r.children = append(r.children, e)
		}
	}
	var result2 []*Enemy
	for _, e := range r.children {
		if e.frame < e.cfg.DurationFrame {
			result2 = append(result2, e)
			e.Update()
		}
	}
	r.children = result2
	r.frame++
	if r.parent == nil {
		global.Frame = r.frame
	}
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

func (r *Enemy) Draw(screen *ebiten.Image) {
	if r.pic != nil {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(-r.cfg.PicX, -r.cfg.PicY)
		opt.GeoM.Scale(r.cfg.PicScale, r.cfg.PicScale)
		opt.GeoM.Rotate(r.cfg.Rotate + r.rad)
		opt.GeoM.Translate(r.point.X, r.point.Y)
		screen.DrawImage(r.pic, opt)
	}
	for _, e := range r.children {
		e.Draw(screen)
	}
}

func (r *Enemy) calculate(s string) float64 {
	s = strings.ReplaceAll(s, "x.player", strconv.FormatFloat(global.SelfPoint.X, 'f', 2, 64))
	s = strings.ReplaceAll(s, "y.player", strconv.FormatFloat(global.SelfPoint.Y, 'f', 2, 64))
	s = strings.ReplaceAll(s, "x.gen", strconv.FormatFloat(r.cfg.StartX, 'f', 2, 64))
	s = strings.ReplaceAll(s, "y.gen", strconv.FormatFloat(r.cfg.StartY, 'f', 2, 64))
	s = strings.ReplaceAll(s, "x.last", strconv.FormatFloat(r.point.X, 'f', 2, 64))
	s = strings.ReplaceAll(s, "y.last", strconv.FormatFloat(r.point.Y, 'f', 2, 64))
	s = strings.ReplaceAll(s, "rad.last", strconv.FormatFloat(r.rad, 'f', 2, 64))
	s = strings.ReplaceAll(s, "t.now", strconv.Itoa(global.Frame))
	s = strings.ReplaceAll(s, "t.gen", strconv.Itoa(global.Frame-r.frame))
	result, err := engine.ParseAndExec(s)
	if err != nil {
		panic(err)
	}
	return result
}
