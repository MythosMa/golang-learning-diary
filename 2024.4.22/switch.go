package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch {
	case i == 1:
		fmt.Println("one")

	case i == 2:
		fmt.Println("two")

	case i == 3:
		fmt.Println("three")
	}

	switch result := sum(1, 2, 3); {
	case result < 0:
		fmt.Println("negative")
	case result > 0 && result < 10:
		fmt.Println("positive and less than 10")
	case result >= 10:
		fmt.Println("positive and greater than or equal to 10")
	default:
		fmt.Println("neutral")
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	month := 10
	switch month {
	case 1, 2, 12:
		fmt.Println("It's Winter!")
	case 3, 4, 5:
		fmt.Println("It's Spring!")
	case 6, 7, 8:
		fmt.Println("It's Summer!")
	case 9, 10, 11:
		fmt.Println("It's Autumn!")
	}

	s := ""
	for s != "aaaaa" {
		fmt.Println("value of s is: ", s)
		s = s + "a"
	}

}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
