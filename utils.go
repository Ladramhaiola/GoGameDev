package main

import (
	"math"
	"math/rand"
	"time"

	randomata "github.com/Pallinder/go-randomdata"
	"github.com/faiface/pixel"
)

type gameObject interface {
	update(dt float64)
	draw(t pixel.Target)
	setField(area *Area)
	Type() string
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func collide(x1, x2, y1, y2, r1, r2 float64) bool {
	distance := math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)
	rdist := math.Pow(r2+r1, 2)
	return distance < rdist
}

func spawner(area *Area, objectType string) {
	switch objectType {
	case "Asteroid":
		position := pixel.V(0, randomata.Decimal(int(win.Bounds().Size().Y)))
		asteroid := NewAsteroid(position, 3)
		area.addObject(asteroid)
	}
}
