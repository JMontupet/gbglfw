package glfwio

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/jmontupet/gbcore/pkg/coreio"
)

func (io *GlfwIO) listenInputs() {
	var currentKeys coreio.KeyInputState
	io.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		old := currentKeys
		if action == glfw.Press {
			switch key {
			case glfw.KeyW:
				currentKeys |= coreio.GBKeyUP
			case glfw.KeyS:
				currentKeys |= coreio.GBKeyDOWN
			case glfw.KeyD:
				currentKeys |= coreio.GBKeyRIGHT
			case glfw.KeyA:
				currentKeys |= coreio.GBKeyLEFT
			case glfw.KeyUp:
				currentKeys |= coreio.GBKeyA
			case glfw.KeyLeft:
				currentKeys |= coreio.GBKeyB
			case glfw.KeyO:
				currentKeys |= coreio.GBKeySELECT
			case glfw.KeyP:
				currentKeys |= coreio.GBKeySTART
			}
		} else if action == glfw.Release {
			switch key {
			case glfw.KeyW:
				currentKeys &= ^coreio.GBKeyUP
			case glfw.KeyS:
				currentKeys &= ^coreio.GBKeyDOWN
			case glfw.KeyD:
				currentKeys &= ^coreio.GBKeyRIGHT
			case glfw.KeyA:
				currentKeys &= ^coreio.GBKeyLEFT
			case glfw.KeyUp:
				currentKeys &= ^coreio.GBKeyA
			case glfw.KeyLeft:
				currentKeys &= ^coreio.GBKeyB
			case glfw.KeyO:
				currentKeys &= ^coreio.GBKeySELECT
			case glfw.KeyP:
				currentKeys &= ^coreio.GBKeySTART
			}
		}
		if old != currentKeys {
			io.inputQueue <- currentKeys
		}
	})
}
