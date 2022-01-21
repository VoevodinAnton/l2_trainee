package pkg

type transport struct {
	name  string
	speed uint
}

// implement interface
func (t *transport) setName(n string) {
	t.name = n
}

func (t *transport) getName() string {
	return t.name
}

func (t *transport) setSpeed(s uint) {
	t.speed = s
}

func (t *transport) getSpeed() uint {
	return t.speed
}
