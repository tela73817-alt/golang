//go:build !test
// +build !test

package main

import "fmt"

func main() {
	student := NewStudent("Alice")
	if student == nil {
		fmt.Println("Failed to create student")
		return
	}

	err := student.AddGrade(95)
	if err != nil {
		fmt.Println("Error adding grade:", err)
		return
	}

	fmt.Print(student.String())
}
