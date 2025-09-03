package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStudent_ValidName_ReturnStudent(t *testing.T) {
	name := "John Doe"
	student := NewStudent(name)

	assert.Equal(t, name, student.Name)
	assert.Equal(t, 0, len(student.Grade))
	assert.NotNil(t, student)
}

func TestCreateStudent_EmptyName_ReturnNil(t *testing.T) {
	name := ""
	student := NewStudent(name)

	assert.Nil(t, student)
}

func TestAddGrade_ValidGrade_ReturnNil(t *testing.T) {
	student := NewStudent("Jane Doe")
	err := student.AddGrade(85)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(student.Grade))
	assert.NotEmpty(t, student.Grade)
	assert.Equal(t, 85.0, student.Grade[0])
}

func TestAddGrade_GradeBelowZero_ReturnError(t *testing.T) {
	student := NewStudent("Jane Doe")
	err := student.AddGrade(-5)

	assert.NotNil(t, err)
	assert.Equal(t, "grade must be between 0 and 100", err.Error())
	assert.Empty(t, student.Grade)
}

func TestAddGGrade_GradeAboveHundred_ReturnError(t *testing.T) {
	student := NewStudent("Jane Doe")
	err := student.AddGrade(105)

	assert.NotNil(t, err)
	assert.Equal(t, "grade must be between 0 and 100", err.Error())
	assert.Empty(t, student.Grade)
}

func TestAddGrade_GradeZero_ReturnNil(t *testing.T) {
	student := NewStudent("Jane Doe")
	err := student.AddGrade(0)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(student.Grade))
	assert.NotEmpty(t, student.Grade)
	assert.Equal(t, 0.0, student.Grade[0])
}

func TestAddGrade_GradeHundred_ReturnNil(t *testing.T) {
	student := NewStudent("Jane Doe")
	err := student.AddGrade(100)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(student.Grade))
	assert.NotEmpty(t, student.Grade)
	assert.Equal(t, 100.0, student.Grade[0])
}

func TestAddGrade_AddMultipleGrades_ReturnNil(t *testing.T) {
	student := NewStudent("Jane Doe")
	err1 := student.AddGrade(85)
	err2 := student.AddGrade(90)
	err3 := student.AddGrade(78)

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Equal(t, 3, len(student.Grade))
	assert.NotEmpty(t, student.Grade)
	assert.Equal(t, 85.0, student.Grade[0])
	assert.Equal(t, 90.0, student.Grade[1])
	assert.Equal(t, 78.0, student.Grade[2])
}

func TestGetAverage_NoHasGrade_ReturnZero(t *testing.T) {
	student := NewStudent("Jane Doe")
	average := student.GetAverage()
	assert.Equal(t, 0.0, average)
	assert.NotNil(t, average)
}

func TestGetAverage_WithGrades_ReturnCorrectAverage(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(80)
	student.AddGrade(90)
	student.AddGrade(70)

	average := student.GetAverage()
	assert.Equal(t, 80.0, average)
	assert.NotNil(t, average)
}

func TestGetLetterGrade_Average90OrAbove_ReturnA(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(95)
	student.AddGrade(90)
	student.AddGrade(92)

	letterGrade := student.GetLetterGrade()
	assert.Equal(t, "A", letterGrade)
	assert.NotNil(t, letterGrade)
}

func TestGetLetterGrade_Average80To89_ReturnB(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(85)
	student.AddGrade(80)
	student.AddGrade(82)
	letterGrade := student.GetLetterGrade()
	assert.Equal(t, "B", letterGrade)
	assert.NotNil(t, letterGrade)
}

func TestGetLetterGrade_Average70To79_ReturnC(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(75)
	student.AddGrade(70)
	student.AddGrade(72)
	letterGrade := student.GetLetterGrade()
	assert.Equal(t, "C", letterGrade)
	assert.NotNil(t, letterGrade)
}

func TestGetLetterGrade_Average60To69_ReturnD(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(65)
	student.AddGrade(60)
	student.AddGrade(62)
	letterGrade := student.GetLetterGrade()
	assert.Equal(t, "D", letterGrade)
	assert.NotNil(t, letterGrade)
}

func TestGetLetterGrade_AverageBelow60_ReturnF(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(55)
	student.AddGrade(50)
	student.AddGrade(52)
	letterGrade := student.GetLetterGrade()
	assert.Equal(t, "F", letterGrade)
	assert.NotNil(t, letterGrade)
}

func TestGetHighestGrade_NoGrades_ReturnZero(t *testing.T) {
	student := NewStudent("Jane Doe")
	highest := student.GetHighestGrade()
	assert.Equal(t, 0.0, highest)
	assert.NotNil(t, highest)
}

func TestGetHighestGrade_WithGrades_ReturnCorrectHighest(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(85)
	student.AddGrade(92)
	student.AddGrade(78)
	highest := student.GetHighestGrade()
	assert.Equal(t, 92.0, highest)
	assert.NotNil(t, highest)
}

func TestGetLowestGrade_NoGrades_ReturnZero(t *testing.T) {
	student := NewStudent("Jane Doe")
	lowest := student.GetLowestGrade()
	assert.Equal(t, 0.0, lowest)
	assert.NotNil(t, lowest)
}

func TestGetLowestGrade_WithGrades_ReturnCorrectLowest(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(85)
	student.AddGrade(92)
	student.AddGrade(78)
	lowest := student.GetLowestGrade()
	assert.Equal(t, 78.0, lowest)
	assert.NotNil(t, lowest)
}

func TestGetGradeCount_NoGrades_ReturnZero(t *testing.T) {
	student := NewStudent("Jane Doe")
	count := student.GetGradeCount()
	assert.Equal(t, 0, count)
	assert.NotNil(t, count)
}

func TestGetGradeCount_WithGrades_ReturnCorrectCount(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(85)
	student.AddGrade(92)
	student.AddGrade(78)
	count := student.GetGradeCount()
	assert.Equal(t, 3, count)
	assert.NotNil(t, count)
}

func TestString_ValidStudent_ReturnCorrectString(t *testing.T) {
	student := NewStudent("Jane Doe")
	student.AddGrade(88)
	student.AddGrade(92)

	expected := "Sinh viên: Jane Doe, Số điểm: 2, Điểm TB: 90.00, Xếp loại: A"

	// Act
	result := student.String()

	// Assert
	assert.Equal(t, expected, result)
}
