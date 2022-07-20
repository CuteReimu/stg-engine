package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type EnemyDetail struct {
	Id            int     `json:"id"`
	Type          int     `json:"type"`
	StartFrame    int     `json:"start_frame"`
	LoopEndFrame  int     `json:"loop_end_frame"`
	DurationFrame int     `json:"duration_frame"`
	IntervalFrame int     `json:"interval_frame"`
	StartX        float64 `json:"start_x"`
	StartY        float64 `json:"start_y"`
	Radius        float64 `json:"radius"`
	Pic           string  `json:"pic"`
	PicScale      float64 `json:"pic_scale"`
	PicX          float64 `json:"pic_x"`
	PicY          float64 `json:"pic_y"`
	Rotate        float64 `json:"rotate"`
	PicXSlice     int     `json:"pic_x_slice"`
	PicYSlice     int     `json:"pic_y_slice"`
	Color         int     `json:"color"`
	Hp            int     `json:"hp"`
	Generate      []int   `json:"generate"`
	X             string  `json:"x"`
	Y             string  `json:"y"`
	Rad           string  `json:"rad"`
	Clean         bool    `json:"clean"`
}

type EnemyDict map[int]*EnemyDetail

var Enemy = make(EnemyDict)

func (enemy *EnemyDict) Init(buf []byte) error {
	if err := json.Unmarshal(buf, &enemy); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (enemy *EnemyDict) InitReader(reader io.Reader) error {
	if err := json.NewDecoder(reader).Decode(&enemy); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
