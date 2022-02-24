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
	StartCoordinate     int     `json:"start_coordinate"`
	StartP1             float64 `json:"start_p1"`
	StartP2             float64 `json:"start_p2"`
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

type GenerateDict struct {
	Data     map[int]*GenerateDetail
	FrameMap map[int][]*GenerateDetail
}

var Generate = &GenerateDict{
	Data:     make(map[int]*GenerateDetail),
	FrameMap: make(map[int][]*GenerateDetail),
}

func (generate *GenerateDict) Init(buf []byte) error {
	if err := errors.WithStack(json.Unmarshal(buf, &generate)); err != nil {
		return err
	}
	generate.init()
	return nil
}

func (generate *GenerateDict) InitReader(reader io.Reader) error {
	if err := errors.WithStack(json.NewDecoder(reader).Decode(&generate)); err != nil {
		return err
	}
	generate.init()
	return nil
}

func (generate *GenerateDict) init() {
	for _, data := range generate.Data {
		generate.FrameMap[data.StartFrame] = append(generate.FrameMap[data.StartFrame], data)
	}
}
