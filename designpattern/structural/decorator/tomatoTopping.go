package decorator

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	return c.pizza.getPrice() + 7
}
