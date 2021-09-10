package pkg

type Command interface {
	Name() string
	Data() interface{}
}

type RotateCommand struct {
	Angle float64
}

func (r RotateCommand) Name() string {
	return "rotate"
}

func (r RotateCommand) Data() interface{} {
	return r.Angle
}

type ThrottleCommand struct {
	Val float64
}

func (t ThrottleCommand) Name() string {
	return "throttle"
}

func (t ThrottleCommand) Data() interface{} {
	return t.Val
}

type FireCommand struct{}

func (f FireCommand) Name() string {
	return "fire"
}

func (f FireCommand) Data() interface{} {
	return nil
}
