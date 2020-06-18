package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmontupet/gbcore/pkg/emulator"
	"github.com/jmontupet/gbgl/internal/glfwio"
)

func main() {
	// Parse command options
	flag.Parse()
	tails := flag.Args()
	if len(tails) != 1 {
		log.Fatalln("Need valid File name")
	}
	// Read GameData
	gameData, err := ioutil.ReadFile(tails[0])
	if err != nil {
		log.Fatal(err)
	}

	// Init GL render
	glfwRenderer := glfwio.NewGlfwIO()

	// Init OpenAL audio player
	// audioPlayer := openal.NewAudioPlayer(48000, 4354, 3, 1600*2)

	// Init GameBoy. glfw is used as Render and IO manager
	gb, err := emulator.NewEmulator(gameData, glfwRenderer, glfwRenderer, nil)
	if err != nil {
		log.Fatalln(err)
	}
	glfwRenderer.SetWindowTitle(gb.GetGameTitle())

	// Pause
	glfwRenderer.OnPause = func() {
		fmt.Println("SEND PAUSE SIGNAL")
		gb.Pause()
	}
	// Stop
	glfwRenderer.OnStop = func() {
		fmt.Println("SEND STOP SIGNAL")
		gb.Stop()
	}
	// Start
	glfwRenderer.OnStart = func() {
		fmt.Println("SEND START SIGNAL")
		gb.Start()
	}

	// Run GameBoy emulation
	gb.Start()
	// Run Gameboy OpenAL Audio Player
	// go audioPlayer.Play()

	// Run GL render
	glfwRenderer.Render()
}
