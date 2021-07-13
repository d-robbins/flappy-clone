package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Level struct {
	level_  [][]*sdl.Rect
	levelw_ int
	levelh_ int
}

func (l *Level) Reload(file string, width, height int) {
	l.level_ = CreateLevel(file, width, height).level_
}

func CreateLevel(file string, width, height int) *Level {
	newlevel := &Level{levelw_: width, levelh_: height}

	newlevel.level_ = make([][]*sdl.Rect, height)

	for row := 0; row < height; row++ {
		newlevel.level_[row] = make([]*sdl.Rect, width)
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return &Level{}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	j := 0
	for scanner.Scan() {
		for i, r := range scanner.Text() {
			if i < width {
				if r == '0' {
					newlevel.level_[j][i] = new(sdl.Rect)
					newlevel.level_[j][i].X = int32(i * 32)
					newlevel.level_[j][i].Y = int32(j * 32)
					newlevel.level_[j][i].W = 32
					newlevel.level_[j][i].H = 32
				} else {
					newlevel.level_[j][i] = nil
				}
			}
		}
		j++
	}

	return newlevel
}

func (level *Level) DrawLevel(renderer *sdl.Renderer) {
	PLAYERDIST++
	for row := 0; row < level.levelh_; row++ {
		for col := 0; col < level.levelw_; col++ {
			if level.level_[row][col] != nil && !OutOfBounds(SCREEN_WIDTH, SCREEN_HEIGHT, level.level_[row][col]) {
				renderer.SetDrawColor(255, 0, 255, 255)
				renderer.FillRect(level.level_[row][col])

				renderer.DrawRect(level.level_[row][col])
				renderer.SetDrawColor(0, 0, 0, 255)
			}
			// horrible but works
			level.level_[row][col].X -= 1
		}
	}
}

func OutOfBounds(screenw, screenh int, r *sdl.Rect) bool {
	if r.X+32 < SCREEN_WIDTH && r.X > 0 {
		return false
	}

	return true
}
