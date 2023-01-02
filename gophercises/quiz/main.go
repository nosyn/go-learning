/*
		Example how to run this program:
	 - go build ./main.go && ./main -limit=10
	 - Available flags:
	   - limit=<seconds> - Default to 15
	   - csv=<file-name> - Default to problems.csv
*/
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

const (
	DEFAULT_TIME_LIMIT    = 15
	DEFAULT_CSV_FILE_NAME = "problems.csv"
)

func parseFileProblems(data [][]string) (quizzes []Quiz) {
	for _, line := range data {
		quiz := Quiz{}
		for j, field := range line {
			if j == 0 {
				quiz.Question = field
			} else if j == 1 {
				quiz.Answer = field
			}
		}
		quizzes = append(quizzes, quiz)
	}
	return
}

func main() {
	csvFileName := flag.String("csv", DEFAULT_CSV_FILE_NAME, "A CSV file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", DEFAULT_TIME_LIMIT, "the time limit for the quiz in seconds")
	flag.Parse()

	// Open the file
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("Failed to open the CSV file: %s\n", *csvFileName)
		panic(err)
	}
	// Close the file
	defer file.Close()

	// Reading CSV file
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	quizzes := parseFileProblems(data)
	quizzesLength := len(quizzes)
	var score, answeredCount int
	hasTimeExpired := false
	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))

	fmt.Printf("%d second(s). You have %d question(s). Go!\n", *timeLimit, quizzesLength)

ProblemLoop:
	for i, v := range quizzes {
		fmt.Printf("#%d - What is %s: ", i+1, v.Question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			sanitizedAnswer := strings.TrimSpace(answer)
			answerCh <- sanitizedAnswer
			answeredCount++
		}()

		select {
		case <-timer.C:
			fmt.Println("")
			hasTimeExpired = true
			break ProblemLoop
		case answer := <-answerCh:
			if answer == v.Answer {
				score++
				fmt.Println("Corrected")
			} else {
				fmt.Println("Incorrect!. Corrected answer: ", v.Answer)
			}
		}
	}
	fmt.Println("\n---------------------------------------------")
	if hasTimeExpired {
		fmt.Printf("Time up! You have answered %d problems. Your score: %d/%d\n", answeredCount, score, quizzesLength)
	} else {
		fmt.Printf("Congrats! You have answered %d problems. Your score: %d/%d\n", answeredCount, score, quizzesLength)
	}
}
