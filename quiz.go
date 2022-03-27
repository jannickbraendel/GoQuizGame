package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var questions = []string{}
var answers = []string{}
var rightAnswersCount = 0

func main() {
	getQuizData()

	for i, val := range questions {
		fmt.Printf(val + ":")
		var ans string
		fmt.Scanf("%s", &ans)
		if ans == answers[i] {
			//correct answer
			rightAnswersCount++
			fmt.Printf("Correct Answer! Let's go next..\n")
		} else {
			fmt.Printf("Wrong Answer.. Next one will be better.\n")
		}
	}

	fmt.Printf("You have answered all the questions. You had ", rightAnswersCount, " out of ", len(questions), " correct. Congratulations!")
}

func getQuizData() {

	fileName := flag.String("csv", "problems.csv", "csv file in format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*fileName)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		questions = append(questions, line[0])
		answers = append(answers, line[1])
	}

}
