package stg

import (
	"github.com/CuteReimu/stg-engine/config"
	"github.com/CuteReimu/stg-engine/movement"
	"github.com/CuteReimu/stg-engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Bullet struct {
	cfg          *config.BulletDetail
	start        utils.Point
	point        utils.Point
	rad          float64
	frame        int
	duration     int
	move         movement.Movement
	pic          *ebiten.Image
	alreadyGraze bool
}

func NewBullet(id int, point utils.Point, move movement.Movement, duration int) *Bullet {
	cfg := config.Bullet[id]
	return &Bullet{
		cfg:      cfg,
		duration: duration,
		start:    point,
		point:    point,
		move:     move,
		pic:      GetPic(cfg.Pic),
	}
}

func (b *Bullet) IsAlive() bool {
	return b.frame < b.duration
}

func (b *Bullet) GetPosition() utils.Point {
	return b.point
}

func (b *Bullet) SetPosition(point utils.Point) {
	b.point = point
}

func (b *Bullet) CheckCollide(selfPosition utils.Point, selfRadius float64) bool {
	dx, dy := b.point.Diff(selfPosition)
	return dx*dx+dy*dy < selfRadius*selfRadius+b.cfg.Radius*b.cfg.Radius
}

func (b *Bullet) CheckGraze(selfPosition utils.Point, selfRadius float64) bool {
	if !b.alreadyGraze {
		return false
	}
	if b.CheckCollide(selfPosition, selfRadius*8) {
		b.alreadyGraze = true
		return true
	}
	return false
}

func (b *Bullet) Update() error {
	b.point, b.rad = b.move.Move(b.frame, b.start)
	b.frame++
	return nil
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	if b.pic == nil {
		return
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-b.cfg.PicX, -b.cfg.PicY)
	opt.GeoM.Scale(b.cfg.PicScale, b.cfg.PicScale)
	opt.GeoM.Rotate(b.rad - math.Pi/2)
	opt.GeoM.Translate(b.point.X, b.point.Y)
	screen.DrawImage(b.pic, opt)
}

func (b *Bullet) CanClean() bool {
	return b.cfg.CanClean
}
