package decorator

type CheeseTopping struct {
	pizza IPizza
}

func (p *CheeseTopping) getPrice() int {
	return p.pizza.getPrice() + 10
}
