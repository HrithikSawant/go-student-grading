package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, University                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

// Helper Function
func parseStudentRecord(record []string) (student, error) {
	if len(record) < 7 {
		return student{}, fmt.Errorf("record has insufficient fields: %v", record)
	}

	parseScore := func(value string) (int, error) {
		score, err := strconv.Atoi(value)
		if err != nil {
			return 0, fmt.Errorf("invalid score (%v): %w", value, err)
		}
		return score, nil
	}

	test1Score, err := parseScore(record[3])
	if err != nil {
		return student{}, err
	}
	test2Score, err := parseScore(record[4])
	if err != nil {
		return student{}, err
	}
	test3Score, err := parseScore(record[5])
	if err != nil {
		return student{}, err
	}
	test4Score, err := parseScore(record[6])
	if err != nil {
		return student{}, err
	}

	return student{
		firstName:  record[0],
		lastName:   record[1],
		University: record[2],
		test1Score: test1Score,
		test2Score: test2Score,
		test3Score: test3Score,
		test4Score: test4Score,
	}, nil
}

func parseCSV(filePath string) (students []student) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip the header row
	if _, err := reader.Read(); err != nil {
		log.Fatal("error reading header row: %w", err)
		return nil
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error reading record: %w", err)
			return nil
		}

		student, err := parseStudentRecord(record)
		if err != nil {
			log.Fatal("error parsing record: %w", err)
			return nil
		}
		students = append(students, student)
	}
	return students
}

func calculateGrade(students []student) (studentStats []studentStat) {
	/*
		We need to calculate the final score and grade for each of the students.
		The scoring is done as follows:
			Final score is the average of all test scores.
				i.e. if test1=50, test2=60, test3=65 and test4=45,
				then final score = (50 + 60 + 65 + 45)/4, i.e. 55
		Grade is based on the final score.
		If final score is < 35, then student is graded as F (failed)
		If final score is >= 35 and < 50, then student is graded as C
		If final score is >= 50 and < 70, then student is graded as B
		If final score is >= 70, then student is graded as A
		Thus, for a student with final score as 55, the grade will be B
	*/
	var studentGrade Grade

	for _, student := range students {
		finalScore := float32(student.test1Score+student.test2Score+student.test3Score+student.test4Score) / 4

		switch {
		case finalScore < 35:
			studentGrade = F
		case finalScore >= 35 && finalScore < 50:
			studentGrade = C
		case finalScore >= 50 && finalScore < 70:
			studentGrade = B
		case finalScore >= 70 && finalScore <= 100:
			studentGrade = A
		default:
			log.Fatal("Invalid Grade")
			return nil
		}

		studentStat := studentStat{
			student:    student,
			finalScore: float32(finalScore),
			grade:      studentGrade,
		}
		studentStats = append(studentStats, studentStat)
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	if len(gradedStudents) == 0 {
		return studentStat{}
	}

	topper := gradedStudents[0]

	for _, student := range gradedStudents[1:] {
		if student.finalScore > topper.finalScore {
			topper = student
		}
	}

	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	if len(gs) == 0 {
		return map[string]studentStat{}
	}

	topperPerUniversity := make(map[string]studentStat)

	for _, student := range gs {
		// Check if the university is already in the map
		if existingTopper, exists := topperPerUniversity[student.University]; exists {
			// Compare the existing topper with the current student
			if student.finalScore > existingTopper.finalScore {
				topperPerUniversity[student.University] = student
			}
		} else {
			// If no topper exists for this university, assign the current student as the topper
			topperPerUniversity[student.University] = student
		}
	}

	return topperPerUniversity
}
