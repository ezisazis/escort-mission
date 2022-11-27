// Use of this source code is subject to an MIT-style
// licence which can be found in the LICENSE file.

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// LoadingScreen is shown while all the assets are loading.
// When loading is ready it switches to Intro screen
type LoadingScreen struct{}

// Update handles player input to update the start screen
func (s *LoadingScreen) Update() (GameState, error) {
	return gameLoading, nil
}

// Draw renders the start screen to the screen
func (s *LoadingScreen) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Loading ...")
}