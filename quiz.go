package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var rightAnswersCount = 0

type problem struct {
	question string
	answer   string
}

var problems = []problem{}

func main() {
	getQuizData()

	for _, p := range problems {
		fmt.Printf("%s:\n", p.question)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.answer {
			//correct answer
			rightAnswersCount++
		}
	}

	fmt.Printf("You have answered all the questions. You had %d out of %d correct. Congratulations!", rightAnswersCount, len(problems))
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

		problems = append(problems, problem{question: line[0], answer: line[1]})
	}

}
