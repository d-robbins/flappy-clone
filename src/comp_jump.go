package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type JumpComponent struct {
	obj_   *GamePhysicsObject
	click_ bool
}

const (
	JUMP_COOLDOWN = 400
	JUMP_FORCE    = 0.098
)

func NewJumpComponent(obj *GamePhysicsObject) *JumpComponent {

	for _, r := range obj.obj_.comps_ {
		if r == obj.GetComponent(&PhysicsComponent{}) {
			return &JumpComponent{obj_: obj}
		}
	}

	return nil
}

func (csm *JumpComponent) Draw(renderer *sdl.Renderer) error { return nil }

func (csm *JumpComponent) Update(dt float64) error {
	keys := sdl.GetKeyboardState()

	if timet > JUMP_COOLDOWN {
		timet = 0
		csm.click_ = false
	}

	if !csm.click_ {
		if keys[sdl.SCANCODE_SPACE] == 1 {
			csm.obj_.velocity_.y_ = 0
			csm.obj_.forces_.y_ -= JUMP_FORCE * dt
			csm.click_ = true
		}
	} else {
		timet += dt
	}

	return nil
}
