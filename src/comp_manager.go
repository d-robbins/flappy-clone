package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Component interface {
	Update(dt float64) error
	Draw(renderer *sdl.Renderer) error
}

type GameObject struct {
	pos_       Vector
	rot_       float32
	active_    bool
	collision_ bool
	comps_     []Component
}

type GamePhysicsObject struct {
	obj_      *GameObject
	velocity_ Vector
	forces_   Vector
}

func (o *GameObject) AddComponent(c Component) {
	for _, comp := range o.comps_ {
		if reflect.TypeOf(c) == reflect.TypeOf(comp) {
			panic(fmt.Sprintf("Added duplicate component to GameObject %v", reflect.TypeOf(c)))
		}
	}

	o.comps_ = append(o.comps_, c)
}

func (o *GameObject) GetComponent(typeT Component) Component {

	typ := reflect.TypeOf(typeT)
	for _, comp := range o.comps_ {
		if typ == reflect.TypeOf(comp) {
			return comp
		}
	}

	panic(fmt.Sprintf("Attempted to get a componet that this object doesnt have: %v", reflect.TypeOf(typeT)))
}

func (g *GameObject) DrawObject(renderer *sdl.Renderer) error {
	for _, comp := range g.comps_ {
		err := comp.Draw(renderer)
		if err != nil {
			return err
		}

	}

	return nil
}

func (g *GameObject) UpdateObject(dt float64) error {
	for _, comp := range g.comps_ {
		err := comp.Update(dt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *GamePhysicsObject) AddComponent(c Component) {
	o.obj_.AddComponent(c)
}

func (o *GamePhysicsObject) GetComponent(typeT Component) Component {
	return o.obj_.GetComponent(typeT)
}

func (g *GamePhysicsObject) DrawObject(renderer *sdl.Renderer) error {
	return g.obj_.DrawObject(renderer)
}

func (g *GamePhysicsObject) UpdateObject(dt float64) error {
	return g.obj_.UpdateObject(dt)
}
