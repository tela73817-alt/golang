package main

import (
	"errors"
	"fmt"
	"strings"
)

type Student struct {
	Name  string
	Grade []float64
}

func NewStudent(name string) *Student {
	if strings.TrimSpace(name) == "" {
		return nil
	}

	return &Student{
		Name:  name,
		Grade: []float64{},
	}
}

func (s *Student) AddGrade(grade float64) error {
	if grade < 0 || grade > 100 {
		return errors.New("grade must be between 0 and 100")
	}

	s.Grade = append(s.Grade, grade)
	return nil
}

func (s *Student) GetAverage() float64 {
	if len(s.Grade) == 0 {
		return 0
	}

	total := 0.0
	for _, g := range s.Grade {
		total += g
	}

	return total / float64(len(s.Grade))
}

func (s *Student) GetLetterGrade() string {
	total := 0.0
	for _, g := range s.Grade {
		total += g
	}

	average := total / float64(len(s.Grade))

	if average >= 90 {
		return "A"
	} else if average >= 80 {
		return "B"
	} else if average >= 70 {
		return "C"
	} else if average >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func (s *Student) GetHighestGrade() float64 {
	if len(s.Grade) == 0 {
		return 0
	}

	maxGrade := s.Grade[0]
	for _, g := range s.Grade {
		if g > maxGrade {
			maxGrade = g
		}
	}

	return maxGrade
}

func (s *Student) GetLowestGrade() float64 {
	if len(s.Grade) == 0 {
		return 0
	}

	minGrade := s.Grade[0]
	for _, g := range s.Grade {
		if g < minGrade {
			minGrade = g
		}
	}

	return minGrade
}

func (s *Student) GetGradeCount() int {
	return len(s.Grade)
}

func (s *Student) String() string {
	return fmt.Sprintf("Sinh viên: %s, Số điểm: %d, Điểm TB: %.2f, Xếp loại: %s",
		s.Name, s.GetGradeCount(), s.GetAverage(), s.GetLetterGrade())
}
