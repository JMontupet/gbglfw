package glfwio

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/jmontupet/gbcore/pkg/coreio"
)

type GlfwIO struct {
	coreio.FrameDrawer
	coreio.InputsManager

	mDraw *sync.Mutex

	frameBuffer *coreio.FrameBuffer
	colors      *coreio.FrameColors
	dirtyScreen int32

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

	lastTimeEmu float64
	nbFramesEmu int64
	fpsEmu      float64
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

	currentTime := float64(time.Now().UnixNano())
	delta := (currentTime - io.lastTimeEmu) / float64(time.Second)
	io.nbFramesEmu++
	if delta >= 1.0 {
		io.fpsEmu = float64(io.nbFramesEmu) / delta
		io.nbFramesEmu = 0
		io.lastTimeEmu = currentTime
	}
	io.mDraw.Unlock()
	atomic.StoreInt32(&io.dirtyScreen, 1)

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
