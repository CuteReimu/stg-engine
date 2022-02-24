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
	Radius        float64 `json:"radius"`
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

type EnemyDict struct {
	Data     map[int]*EnemyDetail
	FrameMap map[int][]*EnemyDetail
}

var Enemy = new(EnemyDict)

func (enemy *EnemyDict) Init(buf []byte) error {
	if err := json.Unmarshal(buf, &enemy.Data); err != nil {
		return errors.WithStack(err)
	}
	enemy.init()
	return nil
}

func (enemy *EnemyDict) InitReader(reader io.Reader) error {
	if err := json.NewDecoder(reader).Decode(&enemy.Data); err != nil {
		return errors.WithStack(err)
	}
	enemy.init()
	return nil
}

func (enemy *EnemyDict) init() {
	for _, data := range enemy.Data {
		enemy.FrameMap[data.StartFrame] = append(enemy.FrameMap[data.StartFrame], data)
	}
}
