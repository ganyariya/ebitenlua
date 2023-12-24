package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	lua "github.com/yuin/gopher-lua"
)

const (
	LUA_ENTRY_FUNC = "EntryFunc"
)

type LuaEngine struct {
	luaState             *lua.LState
	luaMainThread        *lua.LState
	luaMainEntryFunction *lua.LFunction
}

func preLoadLuaLibrary(luaState *lua.LState) {
	if err := luaState.DoFile("lua/lib/preload.lua"); err != nil {
		log.Fatal(err)
	}
}

func Double(luaState *lua.LState) int {
	v := luaState.ToInt(1) // 1 つ取得する
	luaState.Push(lua.LNumber(v * 2))
	return 1 // 1 つの引数を返す
}

func ShowMessage(luaState *lua.LState) int {
	msg := luaState.ToString(1)
	GameInstance.message += "\n" + msg
	return 0
}

func DrawImage(luaState *lua.LState) int {
	imagePath := luaState.ToString(1)
	img, _, err := ebitenutil.NewImageFromFile("assets/image/" + imagePath + ".png")
	if err != nil {
		log.Fatal(err)
	}
	GameInstance.images = append(GameInstance.images, img)
	return 0
}

func bindGoFunctions(luaState *lua.LState) {
	luaState.SetGlobal("Call_Go_Double", luaState.NewFunction(Double))
	luaState.SetGlobal("Call_Go_ShowMessage", luaState.NewFunction(ShowMessage))
	luaState.SetGlobal("Call_Go_DrawImage", luaState.NewFunction(DrawImage))
}

func NewLuaEngine() *LuaEngine {
	luaState := lua.NewState()
	preLoadLuaLibrary(luaState)
	bindGoFunctions(luaState)
	return &LuaEngine{
		luaState: luaState,
	}
}

func (le *LuaEngine) LoadScript(scriptPath string) {
	if le.luaMainThread != nil {
		le.luaMainThread.Close()
	}

	luaMainThread, _ := le.luaState.NewThread()
	err := luaMainThread.DoFile(scriptPath)
	if err != nil {
		log.Fatal(err)
	}

	luaMainEntryFunction := luaMainThread.GetGlobal(LUA_ENTRY_FUNC).(*lua.LFunction)
	le.luaMainThread = luaMainThread
	le.luaMainEntryFunction = luaMainEntryFunction
}

/*
実行する内容が残っている場合は true を返す
*/
func (le *LuaEngine) Resume() bool {
	if le.luaMainThread == nil {
		return false
	}

	resumeState, err, _ := le.luaState.Resume(
		le.luaMainThread,
		le.luaMainEntryFunction,
	)
	if err != nil {
		log.Fatal(err)
	}

	if resumeState == lua.ResumeOK {
		le.luaMainThread.Close()
		le.luaMainThread = nil
	}
	return resumeState == lua.ResumeYield
}
