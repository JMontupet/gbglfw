package glfwio

import (
	"sync"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/jmontupet/gbcore/pkg/coreio"
)

type GlfwIO struct {
	coreio.FrameDrawer
	coreio.InputsManager

	mDraw *sync.Mutex

	frameBuffer *coreio.FrameBuffer
	colors      *coreio.FrameColors
	dirtyScreen bool

	window  *glfw.Window
	texture uint32
	palette uint32
	program uint32
	vao     uint32

	inputQueue      chan coreio.KeyInputState
	lastSendedInput coreio.KeyInputState

	windowTitle string

	projection        mgl32.Mat4
	projectionUniform int32

	OnPause func()
	OnStop  func()
	OnStart func()
}

func (io *GlfwIO) SwapFrameBuffer(
	frameBuffer *coreio.FrameBuffer,
	colors *coreio.FrameColors,
) (*coreio.FrameBuffer, *coreio.FrameColors) {
	io.mDraw.Lock()
	oldBuff := io.frameBuffer
	io.frameBuffer = frameBuffer
	oldColors := io.colors
	io.colors = colors
	io.dirtyScreen = true
	io.mDraw.Unlock()
	return oldBuff, oldColors
}
func (io *GlfwIO) CurrentInput() coreio.KeyInputState {
	select {
	case inputs := <-io.inputQueue:
		io.lastSendedInput = inputs
		return inputs
	default:
		return io.lastSendedInput
	}
}

func (io *GlfwIO) SetWindowTitle(title string) {
	io.windowTitle = title
}

// NewGlfwIO Create new IO based on glfw
//
// Implements coreio.FrameDrawer & coreio.InputsManager
func NewGlfwIO() *GlfwIO {
	io := new(GlfwIO)
	io.frameBuffer = new(coreio.FrameBuffer)
	io.colors = new(coreio.FrameColors)
	io.inputQueue = make(chan coreio.KeyInputState, 10)
	io.windowTitle = "-"

	io.mDraw = new(sync.Mutex)

	return io
}
