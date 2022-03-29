package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var rightAnswersCount = 0

type problem struct {
	question string
	answer   string
}

var problems = []problem{}

func main() {

	totaltime := flag.Int("time", 30, "Total time available for the quiz in sec.")
	fileName := flag.String("csv", "problems.csv", "csv file in format of 'question,answer'")
	flag.Parse()

	getQuizData(*fileName)

	/*var start string
	fmt.Scanf("Press ENTER to start the quiz!\n", &start)*/

	timer := time.NewTimer(time.Duration(*totaltime) * time.Second)

problemloop:

	for _, p := range problems {
		fmt.Printf("%s:\n", p.question)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop

		case ans := <-answerCh:
			if ans == p.answer {
				//correct answer
				rightAnswersCount++
			}
		}

	}
	fmt.Printf("Time is up! You scored %d out of %d.", rightAnswersCount, len(problems))

}

func getQuizData(fileName string) {

	file, err := os.Open(fileName)

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
