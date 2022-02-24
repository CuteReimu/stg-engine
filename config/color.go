package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type ColorDetail struct {
	Id      int     `json:"id"`
	AScale  float64 `json:"a_scale"`
	BScale  float64 `json:"b_scale"`
	GScale  float64 `json:"g_scale"`
	HRotate float64 `json:"h_rotate"`
	RScale  float64 `json:"r_scale"`
	SScale  float64 `json:"s_scale"`
	VScale  float64 `json:"v_scale"`
}

type ColorDict map[int]*ColorDetail

var Color = make(ColorDict)

func (color ColorDict) Init(buf []byte) error {
	return errors.WithStack(json.Unmarshal(buf, &color))
}

func (color ColorDict) InitReader(reader io.Reader) error {
	return errors.WithStack(json.NewDecoder(reader).Decode(&color))
}
