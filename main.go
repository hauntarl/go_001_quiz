package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	fileName  *string
	timeLimit *int
	shuffle   *bool
)

func init() {
	/*
		flag pkg:
			- provides options for command line programs
			- 'prog.exe' (execution with default flag values)
			- 'prog.exe --flag value' (overrides the default flag value)
			  'prog.exe -flag=value' can use quotes to include whitespaces
			  'prog.exe -flag="space separated"'
			- these are useful if you want to generate different behavior
			  based on given flag inputs
	*/
	// set a csv flag
	fileName = flag.String(
		"csv", "problems.csv",
		"csv file format: 'question,answer'",
	)
	timeLimit = flag.Int(
		"limit", 30,
		"the time limit for quiz in seconds",
	)
	shuffle = flag.Bool(
		"shuffle", false,
		"shuffles quiz problems: 'true/false'",
	)
	flag.Parse() // parse all the flags which were set above this statement
}

type problem struct{ question, answer string }

func main() {
	// open file from flag value
	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("failed to open the file: %s", *fileName))
	}
	defer file.Close()

	// read data from file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided .csv file.")
	}
	problems := parseLines(lines) // parse the data to problem struct
	if *shuffle {
		// bonus: shuffle problems if shuffle flag true
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	var (
		// set timer
		timer = time.NewTimer(time.Duration(*timeLimit) * time.Second)
		ch    = make(chan string)
		count int
	)
	// take a quiz
	for i, p := range problems {
		fmt.Printf("Problem #%-3d: %s = ", i+1, p.question)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ch <- strings.ToLower(ans)
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d/%d correct!", count, len(problems))
			return
		case ans := <-ch:
			if ans == p.answer {
				count++
			}
		}
	}
	timer.Stop()
	fmt.Printf("You got %d/%d correct!", count, len(problems))
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{
			question: line[0],
			answer:   strings.ToLower(strings.TrimSpace(line[1])),
		}
	}
	return res
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
