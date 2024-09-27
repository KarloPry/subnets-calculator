package validators

import (
	"fmt"
)

func HasValidIpAddress(netClass string, formattedIpAddressSection [4]int) (error) {
	if formattedIpAddressSection[0] > 223 {
		err := fmt.Errorf("Invalid IP address (only class A, B or C IP addresses admitted)")
		return err
	}
	for i := 1; i <=3; i++ {
		if formattedIpAddressSection[i] > 255 || formattedIpAddressSection[i] < 0 {
			err := fmt.Errorf("Invalid IP address (range only from 0 to 255)")
			return err
		}
	}
	switch netClass{
	case "A":
		if formattedIpAddressSection[0] > 127 {
			err := fmt.Errorf("Invalid class A IP address")
			return err
		}
		if formattedIpAddressSection[1] != 0 || formattedIpAddressSection[2] != 0 || formattedIpAddressSection[3] != 0 {
			err := fmt.Errorf("Not a net address (did you mean %d.0.0.0?)", formattedIpAddressSection[0])
			return err
		}
	case "B":
		if formattedIpAddressSection[0] < 128 || formattedIpAddressSection[0]  > 191 {
			err := fmt.Errorf("Invalid class B IP address")
			return err
		}
		if formattedIpAddressSection[2] != 0 || formattedIpAddressSection[3] != 0 {
			err := fmt.Errorf("Not a net address (did you mean %d.%d.0.0)", formattedIpAddressSection[0], formattedIpAddressSection[1])
			return err
		}
	case "C":
		if formattedIpAddressSection[0] < 192 || formattedIpAddressSection[0] > 223 {
			err := fmt.Errorf("Invalid class C IP address")
			return err
		}
		if formattedIpAddressSection[3] != 0 {
			err := fmt.Errorf("Not a net address (did you mean %d.%d.%d.0)", formattedIpAddressSection[0], formattedIpAddressSection[1], formattedIpAddressSection[2])
			return err
		}
	default:
		err := fmt.Errorf("Not a valid IP address class")
		return err
	}
	return nil
}
