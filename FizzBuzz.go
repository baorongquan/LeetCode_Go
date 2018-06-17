// LeetCode 412. Fizz Buzz
/*write a program that outputs the string representation of numbers from 1 to n.

but for multiples of three it should output “fizz” instead of the number and for the multiples of five output “buzz”. for numbers which are multiples of both three and five output “fizzbuzz”.*/

/*Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/
package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(fizzBuzz(30))
}

func fizzBuzz(n int) []string {
	var arr []string = make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			arr[i-1] = "FizzBuzz"
			continue
		}
		if i%3 == 0 {
			arr[i-1] = "Fizz"
			continue
		}
		if i%5 == 0 {
			arr[i-1] = "Buzz"
			continue
		}
		arr[i-1] = strconv.Itoa(i)
	}
	return arr
}
