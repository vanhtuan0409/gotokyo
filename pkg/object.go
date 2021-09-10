package pkg

type ObjectType uint

const (
	ObjectPlayer ObjectType = iota
	ObjectBullet
	ObjectPosition
	ObjectBot
)

type Object interface {
	GetPosition() Position
	GetType() ObjectType
}
