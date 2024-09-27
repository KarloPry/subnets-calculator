package parsers

import (
	"strings"
	"strconv"
	"fmt"
)

func ParseIpAddress(ipAddress string) ([4]int, error) {
	ipAddressSection := strings.Split(ipAddress, ".")
	if len(ipAddressSection) != 4 {
		err := fmt.Errorf("IP address should have 4 sections")
		return [4]int{0,0,0,0}, err
	}
	var formatedIpAddressSection [4]int
	for i := 0; i < 4; i++ {
		formatedSlice, err := strconv.Atoi(ipAddressSection[i])
		if err != nil {
			fmt.Println("error parsing the string to integer")
			return [4]int{0,0,0,0}, err
		}
		formatedIpAddressSection[i] = formatedSlice
	}
	return formatedIpAddressSection, nil
}
