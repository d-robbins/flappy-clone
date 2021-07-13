package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var timet float64 = 0

type PhysicsComponent struct {
	gameobj_ *GamePhysicsObject
	speed_   float32
}

const (
	GRAVITY = 0.098
	MASS    = 50.0
)

func NewPhysicsComponenet(obj *GamePhysicsObject, speed float32) *PhysicsComponent {
	return &PhysicsComponent{gameobj_: obj, speed_: speed}
}

func (csm *PhysicsComponent) Draw(renderer *sdl.Renderer) error { return nil }

func (csm *PhysicsComponent) Update(dt float64) error {
	if csm.gameobj_.obj_.pos_.y_ <= 0 {
		csm.gameobj_.velocity_.y_ = 0
		csm.gameobj_.obj_.pos_.y_ = 0
	}

	if csm.gameobj_.obj_.pos_.y_+32 > SCREEN_HEIGHT {
		csm.gameobj_.obj_.pos_.y_ = SCREEN_HEIGHT - 32
		csm.gameobj_.velocity_.y_ = 0
	}

	csm.gameobj_.velocity_.x_ += (csm.gameobj_.forces_.x_ / MASS) * dt
	csm.gameobj_.velocity_.y_ += (csm.gameobj_.forces_.y_ / MASS) * dt

	csm.gameobj_.obj_.pos_.x_ += csm.gameobj_.velocity_.x_ * dt
	csm.gameobj_.obj_.pos_.y_ += csm.gameobj_.velocity_.y_ * dt

	csm.gameobj_.forces_.y_ = GRAVITY
	return nil
}
