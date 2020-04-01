package main

import (
	"fmt"
)

func gradingStudents(grades []int) (roundedGrades []int) {
	fmt.Println("Original grades:", grades)
	for _, grade := range grades {
		if grade >= 38 {
			distanceFromNextMultiple := 5 - (grade+5)%5
			if distanceFromNextMultiple < 3 {
				grade += distanceFromNextMultiple
			}
		}
		roundedGrades = append(roundedGrades, grade)
	}
	fmt.Println("Rounded grades:", roundedGrades)
	return
}

func main() {
	grades := []int{20, 36, 39, 41, 43, 74, 75, 76, 98, 100}
	gradingStudents(grades)
}
