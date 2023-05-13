package main

import (
	"context"
	"fmt"
	"log"

	decorator "github.com/jailtonjunior94/go-decorator/internal/decorators/address"
	decoratorPizza "github.com/jailtonjunior94/go-decorator/internal/decorators/pizza"
	"github.com/jailtonjunior94/go-decorator/internal/services"
	"github.com/jailtonjunior94/go-decorator/pkg/caching"
)

func main() {
	ctx := context.Background()

	cache := caching.NewCache()
	correiosService := services.NewCorreiosService()
	addressDecorator := decorator.NewAddressDecorator(cache, correiosService)
	addressService := services.NewAddressService(addressDecorator)

	address, err := addressService.FetchAddressByZipcode(ctx, "06503015")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Endereço encontrado com sucesso, CEP %s e %s\n", address.CEP, address.Street)

	// Outro exemplo
	pizza := decoratorPizza.NewVeggeMania()
	pizzaWithCheese := decoratorPizza.NewCheeseTopping(pizza)
	pizzaWithCheeseAndTomato := decoratorPizza.NewTomatoTopping(pizzaWithCheese)

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
