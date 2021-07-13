package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type GText struct {
	font_    *ttf.Font
	col_     sdl.Color
	texture_ *sdl.Texture
	surface_ *sdl.Surface
	text_    string
	pos      sdl.Rect
}

func CreateText(font *ttf.Font, color sdl.Color) *GText {
	return &GText{font_: font, col_: color}
}

func (t *GText) SetMessage(msg string, renderer *sdl.Renderer) {
	var err error
	t.surface_, err = t.font_.RenderUTF8Solid(msg, t.col_)
	if err != nil {
		log.Fatal(err)
	}

	t.texture_, err = renderer.CreateTextureFromSurface(t.surface_)
	if err != nil {
		log.Fatal(err)
	}
}

func (t *GText) Draw(renderer *sdl.Renderer) {
	renderer.Copy(t.texture_, &sdl.Rect{X: 0, Y: 0, W: SCREEN_WIDTH, H: 100}, &t.pos)
}
