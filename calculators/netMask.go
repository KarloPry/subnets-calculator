package calculators

import (
	"fmt"
	"math"
)

func CalculateMask(parsedIpAddress [4]int, subnets int) ([4]int, error){
	var netMask [4] int
	var neededBits int = 1
	for int(math.Pow(2,float64(neededBits))) - 2 < subnets {
		neededBits++
	}
	maskBits := make([] int, neededBits)
	fmt.Println(maskBits)
	return netMask, nil
}
