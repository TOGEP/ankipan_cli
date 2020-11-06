package main

import(
  "flag"
  "fmt"
  "os"
)

func main(){
  anserAction := flag.NewFlagSet("anser", flag.ExitOnError)
  listAction := flag.NewFlagSet("list", flag.ContinueOnError)
  addAction := flag.NewFlagSet("add", flag.ExitOnError)
  addProblem := addAction.String("problem", "", "problem")
  addAnser := addAction.String("anser", "", "anser")

  if len(os.Args) < 2{
    fmt.Println("expected 'anser' or 'list' or 'add' subcommands")
    os.Exit(1)
  }

  switch os.Args[1] {
    case "anser":
      anserAction.Parse(os.Args[2:])

    case "list":
      listAction.Parse(os.Args[2:])

    case "add":
      addAction.Parse(os.Args[2:])
      fmt.Println(*addProblem)
      fmt.Println(*addAnser)
  }

}
