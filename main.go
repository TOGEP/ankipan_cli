package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main(){
  runAction := flag.NewFlagSet("run", flag.ExitOnError)
  listAction := flag.NewFlagSet("list", flag.ContinueOnError)
  addAction := flag.NewFlagSet("add", flag.ExitOnError)
  addProblem := addAction.String("problem", "", "problem")
  addAnswer := addAction.String("answer", "", "answer")

  if len(os.Args) < 2{
    fmt.Println("expected 'answer' or 'list' or 'add' subcommands")
    os.Exit(1)
  }

  home, err := os.UserHomeDir()
  if err != nil{
    log.Fatal(err)
  }

  switch os.Args[1] {
  case "run":
      runAction.Parse(os.Args[2:])
      file, err := os.Open(filepath.Join(home, ".ankipan"))
      if err != nil{
        log.Fatal(err)
      }
      defer file.Close()

      lineLen, lines := GetLine(file)

      rand.Seed(time.Now().UnixNano())
      fmt.Println(lines[rand.Intn(lineLen)])



    case "list":
      listAction.Parse(os.Args[2:])
      file, err := os.Open(filepath.Join(home, ".ankipan"))
      if err != nil{
        log.Fatal(err)
      }
      defer file.Close()

      b, err := ioutil.ReadAll(file)
      fmt.Println(string(b))

    case "add":
      addAction.Parse(os.Args[2:])
      file, err := os.OpenFile(filepath.Join(home, ".ankipan"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
      if err != nil{
        log.Fatal(err)
      }
      defer file.Close()

      fmt.Fprintf(file, "%s, %s\n", *addProblem, *addAnswer)
  }
}

func GetLine(file *os.File) (int, []string){
  lines := 0
  text := make([]string, 0, 100)
  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    lines++
    text = append(text, scanner.Text())
  }
  return lines, text
}
