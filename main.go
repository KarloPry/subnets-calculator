package main

import (
	"github.com/jmacias1503/subnets-calculator/calculators"
	"github.com/jmacias1503/subnets-calculator/parsers"
	"github.com/jmacias1503/subnets-calculator/validators"
	"fmt"
	"strings"
	"time"
)

type net struct {
	class string
	inputIP string
	ipAddress [4]int
	numberSubnets int
	mask [4]int
	subnets []string
}

func main () {
	fmt.Println("\nsubnet-calculator Copyright (C)", time.Now().Year(), "jmacias1503")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY. This is free software, and you") 
	fmt.Println("are welcome to redistribute it under certain conditions.\n")
	var currentNet net;
	fmt.Print("Enter net class: ")
	fmt.Scan(&currentNet.class)
	fmt.Print("Enter IP Address: ")
	fmt.Scan(&currentNet.inputIP)
	currentNet.class = strings.ToUpper(currentNet.class)
	parsedIp, err := parsers.ParseIpAddress(currentNet.inputIP)
	if err != nil {
		fmt.Println("Error parsing the IP address")
		return
	}
	currentNet.ipAddress = parsedIp
	err = validators.HasValidIpAddress(currentNet.class, [4]int(currentNet.ipAddress))
	if err != nil {
		fmt.Println("IP address not valid: ", err)
		return
	}
	fmt.Print("Enter the number of subnets: ")
	fmt.Scan(&currentNet.numberSubnets)
	currentNet.mask, err = calculators.CalculateMask(currentNet.class, [4]int(currentNet.ipAddress), currentNet.numberSubnets)
	if err != nil {
		fmt.Println("Error calculating the subnet mask: ", err)
		return
	}
	fmt.Println(currentNet.mask)
}
