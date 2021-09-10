package pkg

type Mover interface {
	Object
	GetSpeed() float32
	GetAngle() float32
}
