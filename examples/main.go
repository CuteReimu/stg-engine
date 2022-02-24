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
    "color": 1,
    "duration_frame": 200,
    "generate": [],
    "hp": 100,
    "id": 1,
    "move": "linear",
    "move_p1": 2,
    "move_p2": 2,
    "move_p3": 0,
    "move_p4": 0,
    "move_p5": 0,
    "move_p6": 0,
    "pic": "rice.png",
    "pic_scale": 2,
    "pic_x": 5,
    "pic_x_slice": 1,
    "pic_y": 5,
    "pic_y_slice": 1,
    "start_frame": 0,
    "start_x": 0,
    "start_y": 0
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
	g := &game{}
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("MyGame")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type game struct {
	frame   int
	enemies []*stg.Enemy
}

func (g *game) Update() error {
	runtime.GC()
	g.enemies = append(g.enemies, stg.NewEnemies(g.frame)...)
	var newEnemies []*stg.Enemy
	for _, enemy := range g.enemies {
		if !enemy.IsAlive() {
			continue
		}
		if err := enemy.Update(); err != nil {
			return err
		}
		newEnemies = append(newEnemies, enemy)
	}
	g.enemies = newEnemies
	g.frame++
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}
}

func (g *game) Layout(int, int) (screenWidth int, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
