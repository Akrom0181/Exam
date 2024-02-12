package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

type Employees struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Salary     int    `json:"salary"`
	Experience int    `json:"experience"`
	Age        int    `json:"age"`
}

func SortByAge(employees []Employees) ([]string, error) {
	var result []string
	a := readProductsFromJSON("employees.json")
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j].Age >= a[j+1].Age {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for _, employee := range a {
		result = append(result, fmt.Sprintf("%.d", employee.Age))
	}
	return result, nil
}

func ThreeTopSalary(employees []Employees) ([]string, error) {
	var res []string
	var max []int
	readJson := readProductsFromJSON("employees.json")
	for _, employee := range readJson {
		max = append(max, employee.Salary)
	}
	sort.Slice(max, func(i, j int) bool {
		return max[i] > max[j]
	})
	if len(max) >= 3 {
		res = append(res, fmt.Sprintf("%d", max[0]))
		res = append(res, fmt.Sprintf("%s, %s, %v, %d", max[1],employees.FirstName ))
		res = append(res, fmt.Sprintf("%d", max[2]))
	} else {
		return nil, errors.New("not enough employees's 3 salaries")
	}

	return res, nil
}

func main() {
	// 1 ------------------

	employee := readProductsFromJSON("employees.json")
	age, err := SortByAge(employee)
	if err != nil {
		fmt.Println("error while finding average: ", err)
	}
	writeResultsToFile("AgeSorting.txt", age)

	// 2 ------------------
	salary := readProductsFromJSON("employees.json")
	salaries, err := ThreeTopSalary(salary)
	if err != nil {
		fmt.Println("error while finding average: ", err)
	}
	writeResultsToFile("top3salary.txt", salaries)

}

func readProductsFromJSON(filename string) []Employees {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while reading Json:", err)
		return nil
	}

	var employees []Employees
	if err := json.Unmarshal(data, &employees); err != nil {
		fmt.Println("Error while reading Json:", err)
		return nil
	}
	return employees
}

func writeResultsToFile(filename string, results []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()

	for _, result := range results {
		_, err := file.WriteString(result + "\n")
		if err != nil {
			fmt.Println("Error while writing to file:", err)
			return
		}
	}
}
