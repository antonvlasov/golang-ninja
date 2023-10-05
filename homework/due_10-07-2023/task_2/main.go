package main

import "fmt"

type AutomobileSpecification struct {
	Name               string
	NDoors             int
	RimCenterBore      float32
	RimBoltPatter      string
	Petrol             string
	EngineVolumeLiters float32
	EngineName         string
	HorsePower         float32
	WeightKilogram     int
	HasSunroof         bool
	HasRoof            bool
	NSportFeatures     int
}

func isSportCar(asd *AutomobileSpecification) bool {
    return (*asd).NDoors < 4 && (*asd).HorsePower > 150 && (*asd).WeightKilogram < 1050
}

func main() {
	car1 := AutomobileSpecification{
		Name:               "ВАЗ2106",
		NDoors:             4,
		RimCenterBore:      58.5,
		RimBoltPatter:      "4x98",
		Petrol:             "A92",
		EngineVolumeLiters: 1.57,
		EngineName:         "2106",
		HorsePower:         71.5,
		WeightKilogram:     1035,
		HasSunroof:         false,
		HasRoof:            true,
	}

	fmt.Printf("is %s a sport cat? %t\n", car1.Name, isSportCar(&car1))
	fmt.Printf("is %s a sport cat? %t\n", car1.Name, isSportCar(&car1))
    fmt.Printf("is %s a sport cat? %t\n", car1.Name, isSportCar(&car1))
}
