package main

import(
  "flag"
)

func main(){
  flag.Parse()
  action := flag.Args(0)

  if action == nil{
    // anser action
  }

  if action == "add"{
    // add card
  }

  if action == "list"{
    // get card list
    }
}
