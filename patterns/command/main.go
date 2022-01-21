package main

import (
	"command/pkg"
	"fmt"
)

func main() {
	// initialize a new resaurant
	r := pkg.NewRestaurant()

	// create the list of tasks to be executed
	tasks := []pkg.Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.CleanDishes(),
		r.MakePizza(4),
	}

	// create the cooks that will execute the tasks
	cooks := []*pkg.Cook{
		&pkg.Cook{Name: "Anton"},
		&pkg.Cook{Name: "Valentina"},
	}

	// Assign tasks to cooks alternating between the existing
	// cooks.
	for i, task := range tasks {
		// Using the modulus of the current task index, we can
		// alternate between different cooks
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	// Now that all the cooks have their commands, we can call
	// the `executeCommands` method that will have each cook
	// execute their respective commands
	for _, c := range cooks {
		fmt.Printf("cook %s:\n", c.Name)
		c.ExecuteCommands()
	}
}
