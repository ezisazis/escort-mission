// Copyright 2021 Siôn le Roux.  All rights reserved.
// Use of this source code is subject to an MIT-style
// licence which can be found in the LICENSE file.

package main

import (
	"math"
	"math/rand"
)

// spawnMaxDistance is the distance where the point is activated, if the player is close enough
const spawnMaxDistance = gameWidth/2 + 100

// spawnMinDistance is the distance where the point is deactivated, if the player is too close
const spawnMinDistance = gameWidth/2 + 50

// SpawnPoints is an array of SpawnPoint
type SpawnPoints []*SpawnPoint

// Update updates all the SpawnPoints
func (sps *SpawnPoints) Update(g *GameScreen) {
	for _, s := range *sps {
		s.Update(g)
	}
}

// SpawnPosition describes the spawning position related to the SpawnPoint center
type SpawnPosition struct {
	Distance int // Distance of the position from the center
	Angle    int // Angle of the position
}

// SpawnPoint is a point on the map where zombies are spawn
type SpawnPoint struct {
	Position       Coord
	InitialCount   int
	Continuous     bool
	Zombies        Zombies
	InitialSpawned bool
	PrevPosition   SpawnPosition
	NextSpawn      int
	CanSpawn       bool
	ZombieType     ZombieType
}

// NextPosition gives the offset of the next spawning to the center of the point
func (s *SpawnPoint) NextPosition() Coord {

	// Move further if zombies have been spwaned around the whole circle
	if s.PrevPosition.Angle == 0 {
		s.PrevPosition.Distance = (s.PrevPosition.Distance+1)%2 + 1
	}

	// Spawn positions in every 60 degress (360 / 6)
	s.PrevPosition.Angle = (s.PrevPosition.Angle + 1) % 6

	return Coord{
		X: math.Cos(-2*math.Pi/6*float64(s.PrevPosition.Angle)) * float64(s.PrevPosition.Distance),
		Y: math.Sin(-2*math.Pi/6*float64(s.PrevPosition.Angle)) * float64(s.PrevPosition.Distance),
	}
}

// SpawnZombie spawns one zombie
func (s *SpawnPoint) SpawnZombie(g *GameScreen) {
	var np, nc Coord

	// At least one of the 12 positions should be OK
	for i := 0; i < 12; i++ {
		np = s.NextPosition()
		nc = Coord{s.Position.X + np.X*32, s.Position.Y + np.Y*32}
		if g.LevelMap.isFreeAtCoord(nc) {
			break
		}
	}

	var sprites *SpriteSheet
	switch s.ZombieType {
	case zombieNormal:
		fallthrough
	case zombieCrawler:
		zs := rand.Intn(zombieVariants + 1)
		if zs == zombieVariants {
			// Crawler
			sprites = g.Sprites[spriteZombieCrawler]
			s.ZombieType = zombieCrawler
		} else {
			// Normal
			sprites = g.ZombieSprites[zs]
			s.ZombieType = zombieNormal
		}
	case zombieSprinter:
		sprites = g.Sprites[spriteZombieSprinter]
	case zombieBig:
		sprites = g.Sprites[spriteZombieBig]
	}

	z := NewZombie(s, nc, s.ZombieType, sprites)

	z.Target = g.Player.Object
	g.Space.Add(z.Object)
	g.Zombies = append(g.Zombies, z)
	s.Zombies = append(s.Zombies, z)
	s.NextSpawn = 180 + rand.Intn(180)
}

// Update updates the state of the spawn point
func (s *SpawnPoint) Update(g *GameScreen) {
	if s.InitialSpawned && !s.Continuous {
		return
	}

	// If the player is close to the spawn point then it is activated
	playerDistance := CalcDistance(s.Position.X, s.Position.Y, g.Player.Object.X, g.Player.Object.Y)

	if playerDistance < spawnMaxDistance {
		// Intial spawning
		if !s.InitialSpawned {
			for i := 0; i < s.InitialCount; i++ {
				s.SpawnZombie(g)
			}
			s.InitialSpawned = true
		} else {
			// Continuous spawning one zombie if needed after a while
			if g.Tick%s.NextSpawn == 0 {
				if len(s.Zombies) < s.InitialCount {
					s.CanSpawn = true
				}
			}
		}
	}
	// Zombie is spawned only if the player is not too close
	if playerDistance > spawnMinDistance && s.CanSpawn {
		s.SpawnZombie(g)
		s.CanSpawn = false
	}
}

// RemoveZombie removes a dead zombie from the zombie array of the SpawnPoint
func (s *SpawnPoint) RemoveZombie(z *Zombie) {
	for i, sz := range s.Zombies {
		if sz == z {
			s.Zombies[i] = nil
			s.Zombies = append((s.Zombies)[:i], (s.Zombies)[i+1:]...)
		}
	}
}
