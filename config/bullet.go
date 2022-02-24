package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type BulletDetail struct {
	Id        int     `json:"id"`
	Radius    float64 `json:"radius"`
	Pic       string  `json:"pic"`
	PicScale  float64 `json:"pic_scale"`
	PicX      float64 `json:"pic_x"`
	PicY      float64 `json:"pic_y"`
	PicXSlice int     `json:"pic_x_slice"`
	PicYSlice int     `json:"pic_y_slice"`
	Color     int     `json:"color"`
}

type BulletDict map[int]*BulletDetail

var Bullet = make(BulletDict)

func (bullet BulletDict) Init(buf []byte) error {
	return errors.WithStack(json.Unmarshal(buf, &bullet))
}

func (bullet BulletDict) InitReader(reader io.Reader) error {
	return errors.WithStack(json.NewDecoder(reader).Decode(&bullet))
}
