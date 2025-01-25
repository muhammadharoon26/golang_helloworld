package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95
	studentGrades["David"] = 80
	studentGrades["Eve"] = 92
	studentGrades["Frank"] = 78
	studentGrades["George"] = 88
	studentGrades["Hannah"] = 76
	studentGrades["Ivan"] = 82
	studentGrades["Julia"] = 98
	studentGrades["Kate"] = 86
	studentGrades["Lily"] = 94
	studentGrades["Mike"] = 89
	studentGrades["Nina"] = 96
	studentGrades["Oscar"] = 84
	studentGrades["Pamela"] = 91

	fmt.Println("Marks of Bob:", studentGrades["Bob"])
}
