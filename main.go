package main

import(
  "flag"
  "fmt"
  "os"
)

func main(){
  answerAction := flag.NewFlagSet("answer", flag.ExitOnError)
  listAction := flag.NewFlagSet("list", flag.ContinueOnError)
  addAction := flag.NewFlagSet("add", flag.ExitOnError)
  addProblem := addAction.String("problem", "", "problem")
  addAnswer := addAction.String("answer", "", "answer")

  if len(os.Args) < 2{
    fmt.Println("expected 'answer' or 'list' or 'add' subcommands")
    os.Exit(1)
  }

  switch os.Args[1] {
    case "answer":
      answerAction.Parse(os.Args[2:])

    case "list":
      listAction.Parse(os.Args[2:])

    case "add":
      addAction.Parse(os.Args[2:])
      fmt.Println(*addProblem)
      fmt.Println(*addAnser)
  }

}
