package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type EnemyDetail struct {
	Id            int     `json:"id"`
	StartFrame    int     `json:"start_frame"`
	DurationFrame int     `json:"duration_frame"`
	Radius        int     `json:"radius"`
	Pic           string  `json:"pic"`
	PicScale      float64 `json:"pic_scale"`
	PicX          int     `json:"pic_x"`
	PicY          int     `json:"pic_y"`
	PicXSlice     int     `json:"pic_x_slice"`
	PicYSlice     int     `json:"pic_y_slice"`
	Color         int     `json:"color"`
	Hp            int     `json:"hp"`
	Generate      []int   `json:"generate"`
	Move          string  `json:"move"`
	MoveP1        int     `json:"move_p1"`
	MoveP2        int     `json:"move_p2"`
	MoveP3        int     `json:"move_p3"`
	MoveP4        int     `json:"move_p4"`
	MoveP5        int     `json:"move_p5"`
	MoveP6        int     `json:"move_p6"`
}

type EnemyDict map[int]*EnemyDetail

var Enemy = make(EnemyDict)

func (enemy EnemyDict) Init(buf []byte) error {
	return errors.WithStack(json.Unmarshal(buf, &enemy))
}

func (enemy EnemyDict) InitReader(reader io.Reader) error {
	return errors.WithStack(json.NewDecoder(reader).Decode(&enemy))
}
