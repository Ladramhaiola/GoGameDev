package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Projectile - rockets, lasers etc
type Projectile struct {
	Speed    float64
	Position pixel.Vec
	Velocity pixel.Vec
	Radius   float64
	Lifetime float64
	Angle    float64

	Field *Area
	dead  bool
}

// all logic here
func (p *Projectile) update(dt float64) {
	p.Lifetime -= dt

	if p.Lifetime < 0 {
		p.dead = true
		p.Field.delObject(p)
	}

	p.Position.X += p.Velocity.X * dt
	p.Position.Y += p.Velocity.Y * dt
}

// draw bullet on screen
// TODO different types of projectiles
func (p *Projectile) draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.Push(p.Position)
	imd.Circle(p.Radius, 0)
	imd.Draw(t)
}

func (p *Projectile) setField(area *Area) {
	p.Field = area
}

// Type -
func (p *Projectile) Type() string {
	return "Projectile"
}

// NewProjectile creates new projectile
func NewProjectile(x, y float64, angle float64) *Projectile {
	return &Projectile{
		Speed:    400,
		Position: pixel.V(x, y),
		Velocity: pixel.V(math.Cos(angle)*400, math.Sin(angle)*400),
		Radius:   3,
		Lifetime: 3,
		Angle:    angle,
	}
}
