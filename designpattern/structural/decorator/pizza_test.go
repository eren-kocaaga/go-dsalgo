package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &VeggieMania{}

	cheese := &CheeseTopping{pizza: pizza}
	cheeseAndTomato := &TomatoTopping{cheese}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", cheeseAndTomato.getPrice())

}
