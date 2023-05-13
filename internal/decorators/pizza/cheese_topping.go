package decorator

type CheeseTopping struct {
	pizza IPizza
}

func NewCheeseTopping(pizza IPizza) *CheeseTopping {
	return &CheeseTopping{pizza: pizza}
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.pizza.GetPrice()
	return pizzaPrice + 10
}
