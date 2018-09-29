package main

import (
	"image/color"
	"reflect"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"

	randomata "github.com/Pallinder/go-randomdata"
)

// Asteroid is your main ENEMY
type Asteroid struct {
	Position pixel.Vec
	Velocity pixel.Vec

	Size   int
	Radius float64
	HP     int

	Color color.Color
	Field *Area
	dead  bool
}

func (a *Asteroid) setField(area *Area) {
	a.Field = area
}

func (a *Asteroid) update(dt float64) {
	a.Position.X += a.Velocity.X * dt
	a.Position.Y += a.Velocity.Y * dt

	for _, object := range a.Field.GameObjects {
		switch reflect.TypeOf(object).String() {
		case "*main.Projectile":
			obj := object.(*Projectile)
			if collide(a.Position.X, obj.Position.X, a.Position.Y, obj.Position.Y, a.Radius, obj.Radius) {
				a.applyDemage(1)
				obj.Field.delObject(object)
			}
		case "*main.Ship":
			obj := object.(*Ship)
			if collide(a.Position.X, obj.Position.X, a.Position.Y, obj.Position.Y, a.Radius, obj.Radius) {
				obj.Die()
				playerDead = true
			}
		}
	}

	screenWidth, screenHeight := win.Bounds().Size().XY()
	if a.Position.X < 0 {
		a.Position.X += screenWidth
	}
	if a.Position.Y < 0 {
		a.Position.Y += screenHeight
	}
	if a.Position.X > screenWidth {
		a.Position.X -= screenWidth
	}
	if a.Position.Y > screenHeight {
		a.Position.Y -= screenHeight
	}
}

func (a *Asteroid) draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.Color = a.Color
	imd.Push(a.Position)
	imd.Circle(a.Radius, 3)
	imd.Draw(a.Field.DrawTarget)
}

func (a *Asteroid) applyDemage(dmg int) {
	a.HP -= dmg
	if a.HP < 0 {
		if a.Size > 1 {
			count := randomata.Number(1, 4)
			for i := 0; i < count; i++ {
				smallerAsteroid := NewAsteroid(a.Position, a.Size-1)
				a.Field.addObject(smallerAsteroid)
			}
		}
		a.Field.Score += a.Size * 100
		a.Field.delObject(a)
	}
}

// Type -
func (a *Asteroid) Type() string {
	return "Asteroid"
}

// NewAsteroid -
func NewAsteroid(position pixel.Vec, size int) *Asteroid {
	a := &Asteroid{
		Position: position,
		Color:    colornames.Crimson,
		Size:     size,
	}
	sizecoef := (4 - a.Size) * randomata.Number(1, 2)
	a.Velocity = pixel.V(randomata.Decimal(-20*sizecoef, 20*sizecoef), randomata.Decimal(-20*sizecoef, 20*sizecoef))
	a.Radius = float64(size * 15)
	a.HP = size
	return a
}
