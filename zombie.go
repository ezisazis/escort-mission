// Copyright 2021 Siôn le Roux.  All rights reserved.
// Use of this source code is subject to an MIT-style
// licence which can be found in the LICENSE file.

package main

import (
	"github.com/solarlune/resolv"
)

// Zombie is a monster that's trying to eat the player character
type Zombie struct {
	Object *resolv.Object
	Angle  float64
}

// MoveUp moves the zombie upwards
func (z *Zombie) MoveUp() {
	z.move(0, -1)
}

// MoveDown moves the zombie downwards
func (z *Zombie) MoveDown() {
	z.move(0, 1)
}

// MoveLeft moves the zombie left
func (z *Zombie) MoveLeft() {
	z.move(-1, 0)
}

// MoveRight moves the zombie right
func (z *Zombie) MoveRight() {
	z.move(1, 0)
}

// Move the Zombie by the given vector if it is possible to do so
func (z *Zombie) move(dx, dy float64) {
	if collision := z.Object.Check(dx, dy, "mob", "wall"); collision == nil {
		z.Object.X += dx
		z.Object.Y += dy
	}
}