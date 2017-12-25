package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type QuestionAnswer struct {
	Question string
	Answer   string
}

func main() {
	csvFilename := os.Args[1]
	csvFile, err := os.Open(csvFilename)
	var yourScore int
	var totalNumberOfQuestions int
	var yourAnswer string
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}
	totalNumberOfQuestions = len(lines)
	timer1 := time.NewTimer(time.Second * 10)

problemloop:
	for i, line := range lines {
		data := QuestionAnswer{
			Question: line[0],
			Answer:   line[1],
		}

		fmt.Println("Question number ", i+1)
		fmt.Println(data.Question)
		answerCh := make(chan string)
		go func() {
			fmt.Scan(&yourAnswer)
			answerCh <- yourAnswer
		}()
		select {
		case <-timer1.C:
			fmt.Println("Time up!")
			break problemloop
		case <-answerCh:
			if yourAnswer == data.Answer {
				fmt.Println("Right Answer")
				yourScore++
			} else {
				fmt.Println("Wrong Answer!")
			}
		}

	}

	fmt.Println("Your score is", yourScore, "Out of", totalNumberOfQuestions)
}
