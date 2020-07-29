package glfwio

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func (io *GlfwIO) initGlfw() {
	// Initialize Glfw
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw :", err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, io.windowTitle, nil, nil)
	if err != nil {
		log.Fatalln("failed to create glfw window :", err)
	}
	window.MakeContextCurrent()
	io.window = window
}
