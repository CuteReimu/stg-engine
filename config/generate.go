package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type GenerateDetail struct {
	Id                  int     `json:"id"`
	StartFrame          int     `json:"start_frame"`
	IntervalFrame       int     `json:"interval_frame"`
	DurationFrame       int     `json:"duration_frame"`
	Bullet              int     `json:"bullet"`
	BulletDurationFrame int     `json:"bullet_duration_frame"`
	Move                string  `json:"move"`
	MoveP1              float64 `json:"move_p1"`
	MoveP2              float64 `json:"move_p2"`
	MoveP3              float64 `json:"move_p3"`
	MoveP4              float64 `json:"move_p4"`
	MoveP5              float64 `json:"move_p5"`
	MoveP6              float64 `json:"move_p6"`
}

type GenerateDict map[int]*GenerateDetail

var Generate = make(GenerateDict)

func (generate GenerateDict) Init(buf []byte) error {
	return errors.WithStack(json.Unmarshal(buf, &generate))
}

func (generate GenerateDict) InitReader(reader io.Reader) error {
	return errors.WithStack(json.NewDecoder(reader).Decode(&generate))
}
