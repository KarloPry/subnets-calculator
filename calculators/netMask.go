package calculators

import (
	"math"
)

func CalculateMask(netClass string, ipAddress [4]int, subnets int) ([4]int, error){
	var fullBits int
	switch netClass {
	case "A":
		fullBits = 1
	case "B":
		fullBits = 2
	case "C":
		fullBits = 3
	}
	var netMask [4] int
	for i := 0; i < fullBits; i++ {
		netMask[i] = 255
	}
	var neededBits int = 1
	for int(math.Pow(2,float64(neededBits))) - 2 < subnets {
		neededBits++
	}
	for i := 0; i < neededBits; i++ {
	}
	return netMask, nil
}
