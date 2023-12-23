package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	lua "github.com/yuin/gopher-lua"
)

var (
	game *Game
)

type Game struct {
	luaState          *lua.LState
	coroutineThread   *lua.LState
	coroutineFunction *lua.LFunction
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.luaState.Resume(g.coroutineThread, g.coroutineFunction)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func init() {
	luaState := lua.NewState()
	if err := luaState.DoFile("entry/talk.lua"); err != nil {
		log.Fatal(err)
	}
	coroutine, _ := luaState.NewThread()
	fn := luaState.GetGlobal("coro").(*lua.LFunction)

	game = &Game{luaState: luaState, coroutineThread: coroutine, coroutineFunction: fn}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
