package pkg

type Command interface {
	Name() string
	Data() interface{}
}

// RotateCommand rotate bot to an absolute radian angle.
type RotateCommand struct {
	angle float32
}

func NewRotateCommand(angle float32) RotateCommand {
	return RotateCommand{
		angle: angle,
	}
}

func (r RotateCommand) Name() string {
	return "rotate"
}

func (r RotateCommand) Data() interface{} {
	return r.angle
}

// ThrottleCommand set bot running speed. Range from [0,1]
type ThrottleCommand struct {
	val float32
}

func NewThrottleCommand(speed float32) ThrottleCommand {
	if speed > 1 {
		speed = 1
	}
	if speed < 0 {
		speed = 0
	}
	return ThrottleCommand{val: speed}
}

func (t ThrottleCommand) Name() string {
	return "throttle"
}

func (t ThrottleCommand) Data() interface{} {
	return t.val
}

type FireCommand struct{}

func NewFireCommand() FireCommand {
	return FireCommand{}
}

func (f FireCommand) Name() string {
	return "fire"
}

func (f FireCommand) Data() interface{} {
	return nil
}
