package raindrops

import "fmt"

const testVersion = 2

func Convert(num int) string {
	str := ""
	if num%3 == 0 {
		str += "Pling"
	}
	if num%5 == 0 {
		str += "Plang"
	}
	if num%7 == 0 {
		str += "Plong"
	}
	if str == "" {
		str = fmt.Sprintf("%v", num)
	}
	return str
}
