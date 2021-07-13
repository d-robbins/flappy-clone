package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type CompSpriteDraw struct {
	gameobj_ *GameObject
	texture_ *sdl.Texture
}

func (csd *CompSpriteDraw) Update(dt float64) error { return nil }

func NewDrawComponent(obj *GameObject, renderer *sdl.Renderer, file string) *CompSpriteDraw {
	img, err := sdl.LoadBMP(file)
	if err != nil {
		fmt.Println("Error loading bitmap: ", err)
		os.Exit(1)
	}
	defer img.Free()

	ptex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("Error creating texture: ", err)
		os.Exit(1)
	}

	return &CompSpriteDraw{gameobj_: obj, texture_: ptex}
}

func (csd *CompSpriteDraw) Draw(renderer *sdl.Renderer) error {
	_, _, w, h, err := csd.texture_.Query()
	if err != nil {
		fmt.Println("Error querying texture: ", csd.texture_)
		return err
	}

	renderer.CopyEx(
		csd.texture_,
		&sdl.Rect{X: 0, Y: 0, W: int32(w), H: int32(h)},
		&sdl.Rect{X: int32(csd.gameobj_.pos_.x_), Y: int32(csd.gameobj_.pos_.y_), W: int32(w), H: int32(h)},
		float64(csd.gameobj_.rot_),
		&sdl.Point{X: 0, Y: 0}, sdl.FLIP_NONE)

	return nil
}
