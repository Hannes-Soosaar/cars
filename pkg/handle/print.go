package handle

import (
	"fmt"
	"os"

	"gitea.kood.tech/hannessoosaar/cars/pkg/models"
)

func printCompareToFile(cars []models.CarModel) {

	file, err := os.Create("../../comparison.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	for _, car := range cars {
		_, err := fmt.Fprintln(file, car.Name, "-", car.Year, "\n Specs \n Engine: ", car.Specs.Engine, "\n Transmission: ", car.Specs.Transmission, "\n Horsepower: ", car.Specs.Horsepower, "\n Drivetrain: ", car.Specs.Drivetrain, "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

}
