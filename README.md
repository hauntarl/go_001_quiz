# Quiz Game

Create a program to run timed quizzes via the command line.

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

**Run Commands:**

- go build .
- quiz-game.exe (executes application with default flags)
- go build. && quiz-game.exe
- quiz-game.exe -h or --help: gives information about cmd flags
- quiz-game.exe --csv file.csv -limit=10 -shuffle=true
- flag format: -flag=value | --flag value
- for file names with whitespaces: --csv "file name.csv"

**Features:**

- command-line flags
- timer
- file handling
- reading user input
- go routines & channels

**Packages explored:**

- fmt
- flag: to set command line flags
- os: to open file from system and to exit app with a status code
- encoding/csv: to read data from .csv files
- strings: to format and clean user input
- time: to set expiry for quiz
* rand: to shuffle the problem set
