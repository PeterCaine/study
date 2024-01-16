package main

import "fmt"

const spanish = "Spanish"
const spanishGreeting = "Hola, "
const englishGreeting = "Hello, "

func Hello(name, lang string) string {
  if name == "" {
    name = "World"
  }
  if lang == spanish {
    return spanishGreeting + name
  }
  return englishGreeting + name
}

func main(){
  fmt.Println(Hello("Chris", ""))
}
