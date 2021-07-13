package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

type GameState int

const (
	PLAY = iota
	MENU
	LOST
)

var PLAYERDIST int = 0

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 4 {
		log.Fatal("usage: .\\", os.Args[0], " <level_file> <level_width> <level_height> \n")
	}

	var state GameState = MENU

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("SDL Init Error: ")
		fmt.Println(err)
	}

	ttf.Init()

	font, err := ttf.OpenFont("../fonts/arial.ttf", 24)
	if err != nil {
		panic(err)
	}

	// Create window
	window, err := sdl.CreateWindow(
		"Game",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("Error creating window: ", err)
	}
	defer window.Destroy()

	// Create renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		fmt.Println("Error creating renderer: ", err)
	}

	defer renderer.Destroy()

	w, _ := strconv.Atoi(os.Args[2])
	h, _ := strconv.Atoi(os.Args[3])
	level := CreateLevel(os.Args[1], w, h)

	menutext := CreateText(font, sdl.Color{R: 0, G: 0, B: 155, A: 255})
	menutext.SetMessage("Press P To Play!", renderer)
	defer menutext.surface_.Free()
	defer menutext.texture_.Destroy()
	menutext.pos = sdl.Rect{X: 0, Y: 0, W: SCREEN_WIDTH, H: 200}

	player := NewPhysicsSprite()
	player.AddComponent(NewDrawComponent(player.obj_, renderer, "../res/player.bmp"))
	player.AddComponent(NewPhysicsComponenet(player, 0.1))
	player.AddComponent(NewCollisionComponent(player.obj_, level))
	player.AddComponent(NewJumpComponent(player))

	player.obj_.pos_.x_ = 32

	var last, now uint64

	now = sdl.GetPerformanceCounter()
	last = 0
	var delta_time float64 = 0.0

	score := 0

	view := sdl.Rect{X: 0, Y: 0, W: SCREEN_WIDTH, H: SCREEN_HEIGHT}

	// Main game loop
	for {
		last = now
		now = sdl.GetPerformanceCounter()

		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_P] == 1 && state == MENU {
			state = PLAY
		}

		if state == LOST {
			player.obj_.collision_ = false
			player.ResetPhysicsSprite()
			level.Reload(os.Args[1], w, h)
			player.obj_.pos_.x_ = 32
			PLAYERDIST = 0
			score = 0
			state = MENU
		}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		delta_time = ((float64(now) - float64(last)) * 1000) / float64(sdl.GetPerformanceFrequency())

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		switch state {
		case PLAY:
			if PLAYERDIST%128 == 0 {
				score++
			}

			renderer.SetViewport(&view)
			level.DrawLevel(renderer)

			player.DrawObject(renderer)
			player.UpdateObject(delta_time)

			text := CreateText(font, sdl.Color{R: 255, G: 255, B: 0, A: 255})
			defer text.surface_.Free()
			defer text.texture_.Destroy()
			text.pos = sdl.Rect{X: SCREEN_WIDTH - 64, Y: 0, W: 64, H: 64}
			text.SetMessage(strconv.Itoa(score), renderer)
			text.Draw(renderer)

		case MENU:
			menutext.Draw(renderer)
		}

		if player.obj_.collision_ {
			state = LOST
		}

		renderer.Present()
	}
}
