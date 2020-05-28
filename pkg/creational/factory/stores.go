package factory

// Пиццерия в Нью-Йоркском стиле.
type nyPizzaStore struct {
	abstractPizzaStore
}

// Создать пиццу в Нью-Йоркском стиле.
func createNYPizza(pizzaType string) (Pizza, error) {
	var (
		pizza Pizza
		err   error
	)

	switch pizzaType {
	case CheesePizza:
		pizza = newCheesePizza(nyPizzaIngredientFactory{}, "diagonal")
		pizza = pizza.setName("New-York Style Cheese Pizza")

	case PepperoniPizza:
		pizza = newPepperoniPizza(nyPizzaIngredientFactory{}, "diagonal")
		pizza = pizza.setName("New-York Style Pepperoni Pizza")

	case ClamPizza:
		pizza = newClamPizza(nyPizzaIngredientFactory{}, "diagonal")
		pizza = pizza.setName("New-York Style Clam Pizza")

	case VeggiePizza:
		pizza = newVeggiePizza(nyPizzaIngredientFactory{}, "diagonal")
		pizza = pizza.setName("New-York Style Cheese Pizza")

	default:
		err = newPizzaTypeError(pizzaType)
	}

	return pizza, err
}

// NewNYPizzaStore Создать пиццерию в Нью-Йоркском стиле.
func NewNYPizzaStore() PizzaStore {
	return nyPizzaStore{
		abstractPizzaStore{
			createPizza: createNYPizza,
		},
	}
}

// Пиццерия в Чикагском стиле.
type chicagoPizzaStore struct {
	abstractPizzaStore
}

// Создать пиццу в Чикагском стиле.
func createChicagoPizza(pizzaType string) (Pizza, error) {
	var (
		pizza Pizza
		err   error
	)

	switch pizzaType {
	case CheesePizza:
		pizza = newCheesePizza(chicagoPizzaIngredientFactory{}, "square")
		pizza = pizza.setName("Chicago Style Cheese Pizza")

	case PepperoniPizza:
		pizza = newPepperoniPizza(chicagoPizzaIngredientFactory{}, "square")
		pizza = pizza.setName("Chicago Style Pepperoni Pizza")

	case ClamPizza:
		pizza = newClamPizza(chicagoPizzaIngredientFactory{}, "square")
		pizza = pizza.setName("Chicago Style Clam Pizza")

	case VeggiePizza:
		pizza = newVeggiePizza(chicagoPizzaIngredientFactory{}, "square")
		pizza = pizza.setName("Chicago Style Cheese Pizza")

	default:
		err = newPizzaTypeError(pizzaType)
	}

	return pizza, err
}

// NewChicagoPizzaStore Создать пиццерию в Чикагском стиле.
func NewChicagoPizzaStore() PizzaStore {
	return chicagoPizzaStore{
		abstractPizzaStore{
			createPizza: createChicagoPizza,
		},
	}
}
