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

const bulletConfig = `{
  "1": {
    "color": 0,
    "id": 1,
    "pic": "rice.png",
    "pic_scale": 1,
    "pic_x": 3,
    "pic_x_slice": 1,
    "pic_y": 5,
    "pic_y_slice": 1,
    "radius": 3
  }
}`

const generateConfig = `{
  "1": {
    "bullet": 1,
    "bullet_duration_frame": 0,
    "duration_frame": 10000,
    "id": 1,
    "interval_frame": 10,
    "move": "rotate_linear",
    "move_p1": 0,
    "move_p2": 0,
    "move_p3": 0.01,
    "move_p4": 3,
    "move_p5": 0,
    "move_p6": 0,
    "start_frame": 0
  },
  "2": {
	"bullet": 1,
	"bullet_duration_frame": 0,
	"duration_frame": 10000,
	"id": 2,
	"interval_frame": 10,
	"move": "rotate_linear",
	"move_p1": 90,
	"move_p2": 0,
	"move_p3": 0.01,
	"move_p4": 3,
	"move_p5": 0,
	"move_p6": 0,
	"start_frame": 0
  },
  "3": {
	"bullet": 1,
	"bullet_duration_frame": 0,
	"duration_frame": 10000,
	"id": 2,
	"interval_frame": 10,
	"move": "rotate_linear",
	"move_p1": 180,
	"move_p2": 0,
	"move_p3": 0.01,
	"move_p4": 3,
	"move_p5": 0,
	"move_p6": 0,
	"start_frame": 0
  },
  "4": {
	"bullet": 1,
	"bullet_duration_frame": 0,
	"duration_frame": 10000,
	"id": 2,
	"interval_frame": 10,
	"move": "rotate_linear",
	"move_p1": 270,
	"move_p2": 0,
	"move_p3": 0.01,
	"move_p4": 3,
	"move_p5": 0,
	"move_p6": 0,
	"start_frame": 0
  }
}`

const enemyConfig = `{
  "1": {
    "color": 1,
    "duration_frame": 10000,
    "generate": [1,2,3,4],
    "hp": 100,
    "id": 1,
    "move": "stay",
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
    "start_x": 200,
    "start_y": 200
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
	err = config.Bullet.Init([]byte(bulletConfig))
	if err != nil {
		panic(err)
	}
	err = config.Generate.Init([]byte(generateConfig))
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
	bullets []*stg.Bullet
}

func (g *game) Update() error {
	runtime.GC()
	g.enemies = append(g.enemies, stg.NewEnemies(g.frame)...)
	var newEnemies []*stg.Enemy
	for _, enemy := range g.enemies {
		if !enemy.IsAlive() {
			continue
		}
		if bullets, err := enemy.Update(); err != nil {
			return err
		} else {
			g.bullets = append(g.bullets, bullets...)
		}
		newEnemies = append(newEnemies, enemy)
	}
	g.enemies = newEnemies
	var newBullets []*stg.Bullet
	for _, bullet := range g.bullets {
		if !bullet.IsAlive() {
			continue
		}
		if err := bullet.Update(); err != nil {
			return err
		}
		newBullets = append(newBullets, bullet)
	}
	g.bullets = newBullets
	g.frame++
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}
	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}
}

func (g *game) Layout(int, int) (screenWidth int, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
