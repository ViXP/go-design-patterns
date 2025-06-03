// Package chain_of_responsibility implements the Chain Of Responsibility design pattern.
// Meals are prepared by executing a sequence of cooking steps, each represented as a process in a chain.
// Every process performs a specific action and delegates to the next step by calling its Handle() method.
// After the final step, a printable receipt is generated to represent the completed meal.
package chain_of_responsibility

import (
	"fmt"
)

func Run() {
	cookTheBurger()
	cookTheHotDog(Frankfurter)
}

func cookTheBurger() {
	meal := NewMeal("burger")

	step1 := &AddBun{CookingStep{Meal: meal}, bottom}
	step2 := &AddPatty{CookingStep{Meal: meal}}
	step3 := &AddVeggies{CookingStep{Meal: meal}}
	step4 := &AddBun{CookingStep{Meal: meal}, top}

	// Cooking the burger:
	step1.AddNext(step2).AddNext(step3).AddNext(step4)

	step1.Handle()

	fmt.Print(meal)
}

func cookTheHotDog(sausageType byte) {
	meal := NewMeal("hot-dog")

	step1 := &AddBun{CookingStep{Meal: meal}, bottom}
	step2 := &AddSausage{sausageType, CookingStep{Meal: meal}}
	step3 := &AddSauce{CookingStep{Meal: meal}}
	step3_5 := &AddSauce{CookingStep{Meal: meal}}
	step4 := &AddVeggies{CookingStep{Meal: meal}}
	step5 := &AddBun{CookingStep{Meal: meal}, top}

	// Cooking the Hot-Dog:
	step1.AddNext(step2).AddNext(step3).AddNext(step3_5).AddNext(step4).AddNext(step5)

	step1.Handle()

	fmt.Print(meal)
}
