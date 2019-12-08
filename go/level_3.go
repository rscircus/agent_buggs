// Objective: Make a successful GPS request

package main

import "strings"

func main() {
	println("Location:", gpsRequest(newAgent()))
}

func newAgent() string {

	// The tough part here is to understand the three functions called for in the printlns below
	// EqualFold checks for equality under Unicode case-folding
	// Count counts the substring a[1] in a[2]
	return "D d Budds"

}

func gpsRequest(agent string) string {
	a := strings.Split(agent, " ")

	println(len(a))
	println(strings.EqualFold(a[0], a[1]))
	println(strings.Count(a[2], a[1]))

	denied := len(a) != 3 ||
		!strings.EqualFold(a[0], a[1]) ||
		!(strings.Count(a[2], a[1]) > 0)

	if denied {
		return "ACCESS DENIED"
	}

	return "\u0031\u0036\u002e\u0037\u0033\u0033\u0033\u002c\u002d\u0031\u0036\u0039\u002e\u0035\u0032\u0037\u0034"
}
