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

type problem struct {
	ques, ans string
}

func main() {
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
	csvFileName := flag.String("csv", "problems.csv", "csv file format: 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	isShuffle := flag.Bool("shuffle", false, "shuffles quiz problems: 'true/false'")
	// parse all the flags which were set above this statement
	flag.Parse()

	// open file from flag value
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file: %s", *csvFileName))
	}
	defer file.Close()

	// read data from file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	// parse the data to problem struct
	problems := parseLines(lines)
	// bonus: shuffle problems if shuffle flag true
	if *isShuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	}

	// set timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var (
		ansCh   chan string = make(chan string)
		correct int
	)
	// take a quiz
	for i, p := range problems {
		fmt.Printf("Problem #%-3d: %s = ", i+1, p.ques)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- strings.ToLower(ans)
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d/%d correct!", correct, len(problems))
			return
		case ans := <-ansCh:
			if ans == p.ans {
				correct++
			}
		}
	}
	timer.Stop()
	fmt.Printf("You got %d/%d correct!", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			ques: line[0],
			ans:  strings.ToLower(strings.TrimSpace(line[1])),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
