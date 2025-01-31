/*

=======================
Scilla - Information Gathering Tool
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/scilla

	@Author:      edoardottt, https://www.edoardoottavianelli.it

*/

package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//CheckPortsArray checks the basic rules to
//be valid and then returns the ports array to scan.
func CheckPortsArray(input string) []int {
	delimiter := byte(',')
	sliceOfPorts := strings.Split(input, string(delimiter))
	sliceOfPorts = RemoveDuplicateValues(sliceOfPorts)
	result := []int{}
	for _, elem := range sliceOfPorts {
		try, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("The inputted ports array is not valid.")
			os.Exit(1)
		}
		if try > 0 && try < 65536 {
			result = append(result, try)
		}
	}
	return result
}

//CheckPortsRange checks the basic rules to
//be valid and then returns the starting port and the ending port.
func CheckPortsRange(portsRange string, StartPort int, EndPort int) (int, int) {
	// If there's ports range, define it as inputs for the struct
	delimiter := byte('-')
	//If there is only one number

	// If starting port isn't specified
	if portsRange[0] == delimiter {
		maybeEnd, err := strconv.Atoi(portsRange[1:])
		if err != nil {
			fmt.Println("The inputted port range is not valid.")
			os.Exit(1)
		}
		if maybeEnd >= 1 && maybeEnd <= EndPort {
			EndPort = maybeEnd
		}
	} else if portsRange[len(portsRange)-1] == delimiter {
		// If ending port isn't specified
		maybeStart, err := strconv.Atoi(portsRange[:len(portsRange)-1])
		if err != nil {
			fmt.Println("The inputted port range is not valid.")
			os.Exit(1)
		}
		if maybeStart > 0 && maybeStart < EndPort {
			StartPort = maybeStart
		}
	} else if !strings.Contains(portsRange, string(delimiter)) {
		// If a single port is specified
		maybePort, err := strconv.Atoi(portsRange)
		if err != nil {
			fmt.Println("The inputted port range is not valid.")
			os.Exit(1)
		}
		if maybePort > 0 && maybePort < EndPort {
			StartPort = maybePort
			EndPort = maybePort
		}
	} else {
		// If a range is specified
		sliceOfPorts := strings.Split(portsRange, string(delimiter))
		if len(sliceOfPorts) != 2 {
			fmt.Println("The inputted port range is not valid.")
			os.Exit(1)
		}
		maybeStart, err := strconv.Atoi(sliceOfPorts[0])
		if err != nil {
			fmt.Println("The inputted port range is not valid.")
			os.Exit(1)
		}
		maybeEnd, err := strconv.Atoi(sliceOfPorts[1])
		if err != nil {
			fmt.Println("The inputted port range is not valid.")
			os.Exit(1)
		}
		if maybeStart > maybeEnd || maybeStart < 1 || maybeEnd > EndPort {
			fmt.Println("The inputted port range is not valid.")
			os.Exit(1)
		}
		StartPort = maybeStart
		EndPort = maybeEnd
	}
	return StartPort, EndPort
}
