package main

func NewSprite() *GameObject {
	sp := &GameObject{}

	sp.active_ = true
	sp.pos_ = Vector{x_: 0, y_: 0}
	sp.rot_ = 0

	return sp
}

func NewPhysicsSprite() *GamePhysicsObject {
	sp := &GamePhysicsObject{}

	sp.obj_ = &GameObject{}
	sp.forces_ = Vector{0, 0}
	sp.velocity_ = Vector{0, 0}

	return sp
}

func (p *GamePhysicsObject) ResetPhysicsSprite() {
	p.forces_ = Vector{0, 0}
	p.velocity_ = Vector{0, 0}
	p.obj_.pos_ = Vector{0, 0}
}
