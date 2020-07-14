// main.go
package main

import (
	"fmt"
	"math"
)

func toMinSecKm10(km float64, tijd float64) float64 {
	toSec := math.Mod(float64(tijd), 1) / 0.6
	minuten := int(tijd)
	time := float64(minuten) + toSec
	return time / km
}

func toMinSecKm(minseckm float64) float64 {
	toSec := math.Mod(float64(minseckm), 1) * 0.6
	minuten := int(minseckm)
	return float64(minuten) + toSec
}

func kmToTime(minseckm float64, km float64) float64 {
	to10 := toMinSecKm10(1, minseckm)
	time := km * to10
	return toMinSecKm(time)
}

func timeToKmu(minseckm float64) float64 {
	minuten := int(minseckm)
	seconden := math.Mod(float64(minseckm), 1) * 100
	seconden = 60*float64(minuten) + seconden
	return (1000 / seconden) * 3.6
}

func kmuToTime(kmu float64) float64 {
	seconden := 1000 / (kmu / 3.6)
	minuten := int(seconden / 60)
	seconden = (seconden - (float64(minuten) * 60)) / 100
	return float64(minuten) + seconden
}

func main() {
	for {
		var choice int
		fmt.Println("1. Reken min.s/km")
		fmt.Println("2. Reken tijd voor afstand")
		fmt.Println("3. min.s/km naar km/u")
		fmt.Println("4. km/u naar min.s/km")
		fmt.Println("\nKeuze: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("km: ")
			var km float64
			fmt.Scanln(&km)
			fmt.Println("minuten.seconden: ")
			var tijd float64
			fmt.Scanln(&tijd)
			var minseckm10 float64 = toMinSecKm10(km, tijd)
			var minseckm float64 = toMinSecKm(minseckm10)
			fmt.Println("min.s/km: ", minseckm)
		}

		if choice == 2 {
			fmt.Println("minuten.seconden: ")
			var tijd float64
			fmt.Scanln(&tijd)
			fmt.Println("km: ")
			var km float64
			fmt.Scanln(&km)
			time := kmToTime(tijd, km)
			fmt.Println("Tijd(min.s): ", time, " voor ", km, "km")
		}

		if choice == 3 {
			fmt.Println("minuten.seconden: ")
			var tijd float64
			fmt.Scanln(&tijd)
			kmu := timeToKmu(tijd)
			fmt.Println("km/u: ", kmu)
		}

		if choice == 4 {
			fmt.Println("km/u: ")
			var kmu float64
			fmt.Scanln(&kmu)
			tijd := kmuToTime(kmu)
			fmt.Println("Tijd(min.s): ", tijd)
		}
		fmt.Println()
	}
}
