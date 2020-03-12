package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/observer"
)

func main() {
	wd := observer.NewWeatherData()

	fmt.Println("Update 1:")
	currentConditionalsDisplay := observer.NewCurrentConditionsDisplay(wd)
	statisticsDisplay := observer.NewStatisticsDisplay(wd)
	forecastDisplay := observer.NewForecastDisplay(wd)
	fmt.Println(wd.SetMeasurements(20, 60, 560))

	fmt.Println("Update 2:")
	wd.RemoveObserver(forecastDisplay)
	fmt.Println(wd.SetMeasurements(10, 50, 620))

	fmt.Println("Update 3:")
	wd.RemoveObserver(statisticsDisplay)
	fmt.Println(wd.SetMeasurements(30, 80, 700))

	fmt.Println("Update 4:")
	wd.RemoveObserver(currentConditionalsDisplay)
	fmt.Println(wd.SetMeasurements(0, 40, 680))
}
