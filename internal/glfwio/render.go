package glfwio

import (
	"fmt"
	"log"
	"runtime"
	"sync/atomic"
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const screenWidth = 160
const screenHeight = 144
const windowWidth = screenWidth * 3
const windowHeight = screenHeight * 3
const screenRation = float64(screenWidth) / float64(screenHeight)

func (io *GlfwIO) Render() {
	io.initGlfw()
	io.initGlow()
	io.listenInputs()

	io.window.SetFramebufferSizeCallback(func(win *glfw.Window, w int, h int) {
		if win == io.window {
			ratio := float64(w) / float64(h)
			if ratio > screenRation {
				nw := int32(h * screenWidth / screenHeight)
				gl.Viewport((int32(w)-nw)/2, 0, int32(h*screenWidth/screenHeight), int32(h))
			} else {
				nh := int32(w * screenHeight / screenWidth)
				gl.Viewport(0, (int32(h)-nh)/2, int32(w), int32(w*screenHeight/screenWidth))
			}
		}
	})

	// Initialize Glow
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize glow :", err)
	}
	// Disable VSYNC
	glfw.SwapInterval(0)

	lastTime := glfw.GetTime()
	currentTime := lastTime
	delta := currentTime - lastTime
	nbFrames := 0

	// gl.Enable(gl.FRAMEBUFFER_SRGB)

	// No need to change on the loop.
	gl.UseProgram(io.program)
	gl.BindVertexArray(io.vao)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.ActiveTexture(gl.TEXTURE1)

	for !io.window.ShouldClose() {
		io.mDraw.Lock()
		io.cDraw.Wait()

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		currentTime := glfw.GetTime()
		delta = currentTime - lastTime
		nbFrames++
		if delta >= 1.0 {
			fps := float64(nbFrames) / delta
			nbFrames = 0
			lastTime = currentTime
			// io.window.SetTitle(io.windowTitle)
			io.window.SetTitle(fmt.Sprintf("FPS : %.2f - %s", fps, io.windowTitle))
		}

		gl.BindTexture(gl.TEXTURE_2D, io.texture)
		// C'est ici que tout ce passe !!!
		gl.TexSubImage2D(
			gl.TEXTURE_2D,
			0,
			0,
			0,
			screenWidth,
			screenHeight,
			gl.RED,
			gl.UNSIGNED_BYTE,
			atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&io.frameBuffer))),
		)
		gl.BindTexture(gl.TEXTURE_1D, io.palette)
		gl.TexSubImage1D(
			gl.TEXTURE_1D,
			0,
			0,
			64,
			gl.RGB,
			gl.UNSIGNED_BYTE,
			atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&io.colors))),
		)
		gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
		io.window.SwapBuffers()

		io.mDraw.Unlock()

		glfw.PollEvents()
	}
	glfw.Terminate()
}
func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}
