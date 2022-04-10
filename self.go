package stg

import (
	"github.com/CuteReimu/stg-engine/global"
	"github.com/CuteReimu/stg-engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type BaseSelf struct {
}

func (s *BaseSelf) SetPosition(point utils.Point) {
	global.SelfPoint = point
}

func (s *BaseSelf) GetPosition() utils.Point {
	return global.SelfPoint
}

type Self interface {
	Radius() float64
	Pic() *ebiten.Image
	Opt() *ebiten.DrawImageOptions
	PicImg() *ebiten.Image
	OptImg() *ebiten.DrawImageOptions
	IsInvincible() bool
	SetInvincibleFrame(frame int)
	ResetPosition()
	SetPosition(point utils.Point)
	GetPosition() utils.Point
	NextTick() (selfBullets []SelfBullet)
	GetLife() int8
	GetBomb() int8
	DecreaseLife() bool
	DecreaseBomb() bool
}

type SelfBullet interface {
	// IsAlive 例如超出屏幕应该消失时，为了回收内存，应该返回false。如果在屏幕内不该消失，则返回true
	IsAlive() bool
	// Pic 返回图片
	Pic() *ebiten.Image
	// Opt 返回*ebiten.DrawImageOptions
	Opt() *ebiten.DrawImageOptions
	// NextTick 每一帧时应该做什么
	NextTick(bullets []*Bullet, enemy []*Enemy) ([]*Bullet, []*Enemy)
}
