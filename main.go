package main

import(
    "fmt"
    "time"
)


//:= declares new variables, not assigns.
func main(){
    name:="saurabh"
    today:=time.Now().Format("2006-08-26")
    message:="Day 1 done and 29 days more to go"
     fmt.Println("Name:", name)
    fmt.Println("Date:", today)
    fmt.Println("Message:", message)

     fmt.Printf("Name: %s  Date: %s Message: %s\n", name, today, message)
    
}
