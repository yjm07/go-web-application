package main

import "fmt"

// func lenAndUpper(age int, w string, name ...string) (int, string, []string, int) {
// 	defer fmt.Println("done")
// 	return len(name), w, name, age
// }

// func add(numbers ...int) (total int) {
// 	total = 0
// 	for _, num := range numbers {
// 		total += num
// 	}
// 	return
// }

// func canIDrink(age int) bool {
// 	if koAge := age + 2; koAge < 19 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

// func arrayAndSlice() {
// 	names := [3]string{"ab", "cd", "ef"}
// 	ages := []int{1, 2, 3}
// 	age := []int{4, 5, 6}
// 	ages = append(ages, age[0])
// 	fmt.Println(names, ages)
// }

type person struct {
	name    string
	age     int
	favFood []string
}

func mapAndStruct() {
	people := map[string]int{"jimmy": 27, "jinsung": 28}
	fmt.Println(people)

	favFood := []string{"jinro", "iseul"}
	jimmy := person{name: "jimmy", age: 27, favFood: favFood}
	fmt.Println(jimmy)
}

func main() {
	mapAndStruct()
}
