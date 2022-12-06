package flyweight

type CounterTeroristDress struct {
	color string
}

func (t *CounterTeroristDress) getColor() string {
	return t.color
}

func newCounterTerroristDress() *CounterTeroristDress {
	return &CounterTeroristDress{color: "blue"}
}
