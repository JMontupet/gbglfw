package glfwio

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func (io *GlfwIO) initGlow() {
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	shaderVersion := gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))
	renderer := gl.GoStr(gl.GetString(gl.RENDERER))
	log.Println("OpenGL version  :", version)
	log.Println("GLSL   version  :", shaderVersion)
	log.Println("OpenGL renderer :", renderer)

	// Configure the vertex and fragment shaders
	program, err := newGLProgram(vertexShader, fragmentShaderNearest)
	if err != nil {
		panic(err)
	}
	io.program = program
	gl.UseProgram(program)

	// Configure matrix
	io.projection = mgl32.Ortho2D(-1.0, 1.0, 1.0, -1.0)
	io.projectionUniform = gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(io.projectionUniform, 1, false, &io.projection[0])

	// Prepare texture locations
	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	// Load the texture
	io.texture, err = newGLTexture()
	if err != nil {
		log.Fatalln(err)
	}

	// Prepare palette locations
	texturePaletteUniform := gl.GetUniformLocation(program, gl.Str("pal\x00"))
	gl.Uniform1i(texturePaletteUniform, 1)

	// Load the palette
	io.palette, err = newGLPalette()
	if err != nil {
		log.Fatalln(err)
	}

	// Configure the vertex data
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	io.vao = vao

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(planeVertices)*4, gl.Ptr(planeVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	// gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
}

var planeVertices = []float32{
	//  X, Y, Z, U, V
	-1.0, -1.0, 1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, 1.0, 1.0, 0.0,
	1.0, 1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
}
