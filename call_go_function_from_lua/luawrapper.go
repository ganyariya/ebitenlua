package main

import (
	"fmt"
	"log"

	lua "github.com/yuin/gopher-lua"
)

type LuaWrapper struct {
	luaState             *lua.LState
	luaMainThread        *lua.LState
	luaMainEntryFunction *lua.LFunction
}

func Double(luaState *lua.LState) int {
	v := luaState.ToInt(1) // 1 つ取得する ?
	luaState.Push(lua.LNumber(v * 2))
	return 1 // 1 つの引数を返す
}

func Calculate(luaState *lua.LState) int {
	v1, v2 := luaState.ToInt(1), luaState.ToInt(2)
	luaState.Push(lua.LNumber(v1 + v2))
	luaState.Push(lua.LNumber(v1 - v2))
	return 2
}

func SetImage(luaState *lua.LState) int {
	// ここで lua から画像パスを受け取る
	// ebiten で描画する
	luaState.ToInt(1) // 1 つ取得する ?
	return 0          // 何も返さない
}

func NewLuaWrapper() *LuaWrapper {
	luaState := lua.NewState()
	if err := luaState.DoFile("call_go_function_from_lua/talk.lua"); err != nil {
		log.Fatal(err)
	}
	luaMainThread, _ := luaState.NewThread()
	luaMainEntryFunction := luaState.GetGlobal("MainEntry").(*lua.LFunction)

	luaState.SetGlobal("call_go_double", luaState.NewFunction(Double))
	luaState.SetGlobal("call_go_calculate", luaState.NewFunction(Calculate))
	luaState.SetGlobal("call_go_set_image", luaState.NewFunction(SetImage))

	return &LuaWrapper{
		luaState:             luaState,
		luaMainThread:        luaMainThread,
		luaMainEntryFunction: luaMainEntryFunction,
	}
}

func (lw *LuaWrapper) Resume() {
	_, _, result := lw.luaState.Resume(
		lw.luaMainThread,
		lw.luaMainEntryFunction,
	)
	fmt.Println("go", result)
}
