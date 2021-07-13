package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CollisionComponent struct {
	obj_     *GameObject
	collide_ *Level
	active_  bool
}

func NewCollisionComponent(obj *GameObject, objects *Level) *CollisionComponent {
	return &CollisionComponent{obj_: obj, collide_: objects, active_: true}
}

func (c *CollisionComponent) Draw(renderer *sdl.Renderer) error { return nil }

func (c *CollisionComponent) Update(dt float64) error {
	for row := 0; row < c.collide_.levelh_; row++ {
		for col := 0; col < c.collide_.levelw_; col++ {
			if c.collide_.level_[row][col] != nil {
				if c.obj_.pos_.x_+32 < float64(c.collide_.level_[row][col].X) || float64(c.collide_.level_[row][col].X)+32 < c.obj_.pos_.x_ ||
					c.obj_.pos_.y_+32 < float64(c.collide_.level_[row][col].Y) || float64(c.collide_.level_[row][col].Y)+32 < c.obj_.pos_.y_ {
				} else {
					c.obj_.collision_ = true
				}
			}
		}
	}
	return nil
}
