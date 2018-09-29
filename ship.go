package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// Ship - player's ship to dive into space
type Ship struct {
	Position     pixel.Vec
	Rotation     float64
	Velocity     pixel.Vec
	Acceleration float64
	Angle        float64

	ShootDelay float64
	ShootTimer float64

	Radius   float64
	Vertexes []pixel.Vec

	Field *Area
	dead  bool
}

func (s *Ship) update(dt float64) {
	s.ShootTimer -= dt

	if win.Pressed(pixelgl.KeySpace) && s.ShootTimer < 0 {
		s.shoot()
		s.ShootTimer = s.ShootDelay
	}

	if win.Pressed(pixelgl.KeyD) {
		s.Angle -= s.Rotation * dt
	}

	if win.Pressed(pixelgl.KeyA) {
		s.Angle += s.Rotation * dt
	}

	if win.Pressed(pixelgl.KeyW) {
		vxdt := math.Cos(s.Angle) * s.Acceleration * dt
		vydt := math.Sin(s.Angle) * s.Acceleration * dt
		s.Velocity.X += vxdt
		s.Velocity.Y += vydt
	}

	s.Position.X += s.Velocity.X * dt
	s.Position.Y += s.Velocity.Y * dt
	s.Velocity.X -= s.Velocity.X * dt
	s.Velocity.Y -= s.Velocity.Y * dt

	screenWidth, screenHeight := win.Bounds().Size().XY()

	if s.Position.X < 0 {
		s.Position.X += screenWidth
	}

	if s.Position.X > screenWidth {
		s.Position.X -= screenWidth
	}

	if s.Position.Y < 0 {
		s.Position.Y += screenHeight
	}

	if s.Position.Y > screenHeight {
		s.Position.Y -= screenHeight
	}

}

func (s *Ship) draw(t pixel.Target) {
	imd := imdraw.New(nil)
	// draw polygon
	for _, vertex := range s.Vertexes {
		rotated := vertex.Rotated(s.Angle + math.Pi/2)
		imd.Push(rotated.Add(s.Position))
	}
	// set line width
	imd.Polygon(2)
	imd.Draw(t)
}

func (s *Ship) shoot() {
	shotX := s.Position.X + s.Radius*1.2*math.Cos(s.Angle)
	shotY := s.Position.Y + s.Radius*1.2*math.Sin(s.Angle)
	b := NewProjectile(shotX, shotY, s.Angle)
	s.Field.addObject(b)
}

func (s *Ship) setField(area *Area) {
	s.Field = area
}

// Type -
func (s *Ship) Type() string {
	return "Ship"
}

// Die - end game
func (s *Ship) Die() {
	slowment = 0
}

func newDefaultShip(x, y float64) *Ship {
	s := &Ship{
		Position:     pixel.V(x, y),
		Velocity:     pixel.V(0, 0),
		Acceleration: 300,
		Angle:        0,
		Rotation:     math.Pi,
		ShootDelay:   0.3,
		Radius:       30,
		Vertexes: []pixel.Vec{
			pixel.V(0, -30),
			pixel.V(30, 30),
			pixel.V(0, 20),
			pixel.V(-30, 30)},
	}
	return s
}
