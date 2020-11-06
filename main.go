package main

import(
  "flag"
  "fmt"
  "os"
  "database/sql"
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
      db, err := getDB()
      defer db.Close()
      if err != nil{
        panic(err.Error())
      }
      query := "INSERT INTO cards(problem_statement, answer_text, question_time, solved_count) VALUES(?, ?, NOW(), 0"
      _, err = db.Exec(query, *addProblem, *addAnswer)
      if err != nil{
        panic(err.Error())
      }
  }
}

func getDB()(db *sql.DB, err error){
  db, err = sql.Open("mysql", "root:@/ankipan_cli")
  if err != nil{
    panic(err.Error())
  }
  return db, err
}
