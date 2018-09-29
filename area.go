package main

import (
	"github.com/faiface/pixel"
)

// Area -
type Area struct {
	GameObjects map[gameObject]gameObject
	DrawTarget  pixel.Target
	Player      *Ship
	Score       int
	AstroCount  int
}

// Restart game
func (a *Area) Restart() {
	for _, elem := range a.GameObjects {
		a.delObject(elem)
	}
	player := newDefaultShip(400, 300)
	a.addObject(player)
	a.Score = 0
	slowment = 1
}

func (a *Area) addObject(object gameObject) {
	object.setField(a)
	a.GameObjects[object] = object
}

func (a *Area) delObject(object gameObject) {
	_, ok := a.GameObjects[object]
	if ok {
		delete(a.GameObjects, object)
		object = nil
	}
}

func (a *Area) update(dt float64) {
	a.AstroCount = 0
	for _, v := range a.GameObjects {
		if v.Type() == "Asteroid" {
			if v.(*Asteroid).Size == 3 {
				a.AstroCount++
			}
		}
		v.update(dt)
	}
	if a.AstroCount < a.Score/500 || a.AstroCount < 3 {
		spawner(a, "Asteroid")
	}
}

func (a *Area) draw() {
	for _, v := range a.GameObjects {
		v.draw(a.DrawTarget)
	}
}
