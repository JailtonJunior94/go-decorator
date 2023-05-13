package decorator

type TomatoTopping struct {
	pizza IPizza
}

func NewTomatoTopping(pizza IPizza) *TomatoTopping {
	return &TomatoTopping{pizza: pizza}
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.pizza.GetPrice()
	return pizzaPrice + 7
}
