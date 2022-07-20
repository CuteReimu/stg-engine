package main

import (
	"bytes"
	_ "embed"
	"github.com/CuteReimu/stg-engine"
	"github.com/CuteReimu/stg-engine/config"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	_ "image/png"
	"log"
	"runtime"
)

const enemyConfig = `{
  "1": {
    "id": 1,
    "type": 1,
    "start_frame": 0,
    "duration_frame": 1000,
    "interval_frame": 20,
    "color": 0,
    "generate": [],
    "hp": 100,
    "pic": "rice.png",
    "pic_scale": 2,
    "pic_x": 5,
    "pic_x_slice": 1,
    "pic_y": 5,
    "pic_y_slice": 1,
    "start_frame": 0,
    "start_x": 200,
    "start_y": 200,
    "rotate": 1.732,
	"x": "200 + 100 * sin(0.01 * t.now)",
	"y": "200 + 100 * cos(0.01 * t.now)"
  }
}`

//go:embed rice.png
var pic []byte

func init() {
	err := stg.InitPicReader("rice.png", bytes.NewReader(pic))
	if err != nil {
		panic(err)
	}
}

func main() {
	err := config.Enemy.Init([]byte(enemyConfig))
	if err != nil {
		panic(err)
	}
	g := &game{
		enemy: stg.NewEnemies(),
	}
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("MyGame")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type game struct {
	frame int
	enemy *stg.Enemy
}

func (g *game) Update() error {
	runtime.GC()
	g.enemy.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.enemy.Draw(screen)
}

func (g *game) Layout(int, int) (screenWidth int, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
