package main

import (
	"fmt"
)

type People struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Job  string `json:"job"`
}

func main() {
	var num int
	fmt.Println("Nechta Inson kiritmoqchisiz: ")
	fmt.Scan(&num)

	people := make([]People, num)

	for i := 1; i <= num; i++ {
		var a People
		fmt.Println("input name: ")
		fmt.Scan(&a.Name)

		fmt.Println("input sex: ")
		fmt.Scan(&a.Sex)

		fmt.Println("input job: ")
		fmt.Scan(&a.Job)

		people = append(people, People{Name: a.Name, Sex: a.Sex, Job: a.Job})

	}
	var sex int

	fmt.Println("1 male, 2 female, 3 all")
	fmt.Scan(&sex)
	switch sex {
	case 1:
		for _, person := range people {
			if person.Sex == "male" {
				fmt.Println(person)
			}
		}
	case 2:
		for _, person := range people {
			if person.Sex == "female" {
				fmt.Println(person)
			}
		}
	case 3:
		for _, person := range people {
			fmt.Println(person)
		}
	default:
		fmt.Println("no such command")
	}

}

// Terminaldan people count kiritaman. Va undan keyin name, sex, job, age so'ridi.
// Hammasini kiritib bo'lganimdan keyin yana so'ridi faqat males | females | all shunga
// qarab manga male yoki female yoki hammasini chiqarib berishi kerak
