//Course: CS 524 Principles of programimg language
//Instructor: Mr. Williamson James
//Asssinment: #1
//Author: Iswarya Gadde (A25341583)
//File: 524Program1.go
//Date: 06-09-2024 

// Description: This program reads student data from an input file, processes test and homework grades,
// and calculates the overall weighted average for each student. It then displays a grade
// report sorted by student last name and first name.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Student struct to hold information about each student
type Student struct {
	LastName  string
	FirstName string	
	Tests     []float64
	Homeworks []float64
}

func main() {
	// Prompt the user for input file name
	var inputFileName string
	fmt.Print("Enter the name of your input file: ")
	fmt.Scanln(&inputFileName)

	// Open the input file
	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Process student blocks
	// var students []Student
	// var currentStudent *Student

	// for _, line := range lines {
	// 	if line == "" {
	// 		// If line is empty, it indicates the end of a student's data block			
	// 		if currentStudent != nil {
	// 			students = append(students, *currentStudent)
	// 			currentStudent = nil
	// 		}
	// 	} else {
	// 		if currentStudent == nil {
	// 			// First non-empty line of a block contains the student's name				
	// 			names := strings.Split(line, " ")
	// 			currentStudent = &Student{FirstName: names[1], LastName: names[0]}
	// 		} else if len(currentStudent.Tests) == 0 {
	// 			// Second non-empty line contains the test scores
	// 			currentStudent.Tests = parseFloats(line)
	// 		} else if len(currentStudent.Homeworks) == 0 {
	// 			// Third non-empty line contains the homework scores
	// 			currentStudent.Homeworks = parseFloats(line)
	// 		}
	// 	}
	// }

	var students []Student
	var currentStudent *Student

	// Iterate through the lines in blocks of 3
	for i := 0; i < len(lines); i += 3 { 
		names := strings.Split(lines[i], " ")
		testScores := parseFloats(lines[i+1])
		homeworkScores := parseFloats(lines[i+2])

		currentStudent = &Student{
			FirstName: names[1],
			LastName:  names[0],
			Tests:     testScores,
			Homeworks: homeworkScores,
		}
		students = append(students, *currentStudent)	
	}
	
	// Prompt the user for test weight percentage
	var testWeight float64
	fmt.Print("Enter the % amount to weight test in overall avg: ")
	fmt.Scanln(&testWeight)
	testWeight /= 100.0
	
    // Calculate homework weight
    homeworkWeight := 1.0 - testWeight
	fmt.Printf("homework will be weighted remaining %.1f%%\n",homeworkWeight*100)
	
	// Prompt the user for the number of homework assignments
	var numHomeworks int
	fmt.Print("How many homework assignments are there? ")
	fmt.Scanln(&numHomeworks)

	// Prompt the user for the number of test grades
	var numTests int
	fmt.Print("How many test grades are there? ")
	fmt.Scanln(&numTests)

	// Calculate averages
	var totalTestScore, totalHomeworkScore, totalStudents float64
	for _, student := range students {
		totalStudents++
		totalTestScore += sum(student.Tests)
		totalHomeworkScore += sum(student.Homeworks)
	}
   
    // Calculate the overall average for all students
	overallAverage := (totalTestScore/totalStudents)*(testWeight/float64(numTests)) + (totalHomeworkScore/totalStudents)*(homeworkWeight/float64(numHomeworks))

	// Generate grade report
	fmt.Printf("GRADE REPORT --- %d STUDENTS FOUND IN FILE\n", len(students))
	fmt.Printf("TEST WEIGHT: %.1f%%\n", testWeight*100)
	fmt.Printf("HOMEWORK WEIGHT: %.1f%%\n", homeworkWeight*100)
	fmt.Printf("OVERALL AVERAGE is %.1f\n", overallAverage)
	fmt.Println("STUDENT NAME : TESTS     HOMEWORKS   AVG")
	fmt.Println("---------------------------------------------------------")

	// Sort students by last name and first name
	sort.Slice(students, func(i, j int) bool {
		if students[i].FirstName == students[j].FirstName {
			return students[i].LastName < students[j].LastName
		}
		return students[i].FirstName < students[j].FirstName
	})

	// Debug Function: Print sorted student names to verify order
	//fmt.Println("Sorted students:")
	//for _, student := range students {
		//fmt.Printf("%s, %s\n", student.LastName, student.FirstName)
	//}

	// Print each student's information and average
	for _, student := range students {
		testCount := len(student.Tests)
		homeworkCount := len(student.Homeworks)		
		average := (sum(student.Tests) * testWeight / float64(numTests)) + (sum(student.Homeworks) * homeworkWeight / float64(numHomeworks))

		fmt.Printf("%s, %s :   %s (%d)   %s (%d)    %.1f", student.FirstName, student.LastName, formatScores(student.Tests), testCount, formatScores(student.Homeworks), homeworkCount, average)
		if homeworkCount < numHomeworks {
			fmt.Println(" ** may be missing a homework **")
		} else {
			fmt.Println()
		}
		if testCount < numTests {
			fmt.Println(" ** may be missing a test **")
		} else {
			fmt.Println()
		}
	}
}

// Helper function to calculate the sum of numbers in a slice
func sum(numbers []float64) float64 {
	var total float64
	for _, num := range numbers {
		total += num
	}
	return total
}

// Helper function to parse a string containing float64 values
func parseFloats(str string) []float64 {
	var floats []float64
	str = strings.TrimSpace(str)
	if str == "" {
		return floats
	}
	strFloats := strings.Split(str, " ")
	for _, s := range strFloats {
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			floats = append(floats, f)
		}
	}
	return floats
}

// Helper function to format scores as a string
func formatScores(scores []float64) string {
	var formattedScores []string
	for _, score := range scores {
		formattedScores = append(formattedScores, strconv.FormatFloat(score, 'f', 1, 64))
	}
	return strings.Join(formattedScores, " ")
}
