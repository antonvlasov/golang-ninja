package main

import "fmt"

type AutomobileSpecification struct {
	Name           string
	NDoors         int
	HorsePower     float32
	WeightKilogram int
	NSportFeatures int
}

func isSportCar(spec AutomobileSpecification) bool {
	if spec.NDoors < 4 {
		spec.NSportFeatures += 1
	}

	if spec.HorsePower > 150 {
		spec.NSportFeatures += 1
	}

	if spec.WeightKilogram < 1050 {
		spec.NSportFeatures += 1
	}

	return spec.NSportFeatures > 2
}

func main() {
	car1 := AutomobileSpecification{
		Name:           "ВАЗ2106",
		NDoors:         4,
		HorsePower:     71.5,
		WeightKilogram: 1035,
	}

	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(car1))
	// did something change?
	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(car1))
	// did something change?
	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(car1))
}
