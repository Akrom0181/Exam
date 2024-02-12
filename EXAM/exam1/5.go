package main

import (
	"fmt"
	"os"
)

type Student struct {
	Name   string
	Grades []int
	Course int
}

func (s *Student) calculateAverage() float64 {
	sum := 0
	count := 0
	for i := 0; i < len(s.Grades); i++ {
		sum += s.Grades[i]
		count++
	}
	average := float64(sum) / float64(count)
	if average >= 60 {
		return average
	} else {
		return 0
	}
}

func main() {
	student := Student{
		Name: "Abbos",
		Grades: []int{
			60, 50, 60, 56, 89,
		},
		Course: 2,
	}
	a := student.calculateAverage()
	File, err := os.Create("Fiveth.txt")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}
	defer File.Close()

	var TrueOrFalse string
	if a > 0 {
		TrueOrFalse = "you passed"
	} else {
		TrueOrFalse = "you failed"
	}
	_, err = File.WriteString(TrueOrFalse)
	if err != nil {
		fmt.Println("Error while writing to file: ", err)
	}
}

// Student degan struct name, grades(arrays of int),course. calculateAverage degan method
// yozilsin u studentning o'rtacha grade ini hisoblaydi. Agar averageGrade 60dan kichkina
// bo'lsa u failed bo'lgan hisoblanadi aksi bo'lsa passed. U fail yoki passed bo'lgani haqida
// result.txt filega yo'zing. (You are failed yoki You are passed)
