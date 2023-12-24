package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

var (
	GameInstance *Game
)

type Game struct {
	luaEngine *LuaEngine
	message   string
	images    []*ebiten.Image
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.luaEngine.Resume()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, img := range g.images {
		screen.DrawImage(img, &ebiten.DrawImageOptions{})
	}
	ebitenutil.DebugPrint(screen, g.message)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func init() {
	luaEngine := NewLuaEngine()
	luaEngine.LoadScript("lua/script/day01.lua")
	GameInstance = &Game{
		luaEngine: luaEngine,
		images:    make([]*ebiten.Image, 0),
	}
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_WIDTH)
	ebiten.SetWindowTitle("LuaSample")

	if err := ebiten.RunGame(GameInstance); err != nil {
		log.Fatal(err)
	}
}
