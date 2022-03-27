package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var questions = []string{}
var answers = []string{}
var rightAnswersCount = 0

func main() {
	fmt.Printf("Welcome to this Quiz!")
	getQuizData()

	for i, val := range questions {
		fmt.Printf(val + ":")
		var ans string
		fmt.Scanf("%s", &ans)
		if ans == answers[i] {
			//correct answer
			rightAnswersCount++
			fmt.Printf("Correct Answer! Let's go next..")
		} else {
			fmt.Printf("Wrong Answer.. Next one will be better.")
		}
	}

	fmt.Printf("You have answered all the questions. You had ", rightAnswersCount, " out of ", len(questions), " correct. Congratulations!")
}

func getQuizData() {

	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}

		questions = append(questions, line[0])
		answers = append(answers, line[1])
	}

}
