package main

import "fmt"

const (
  spanish = "Spanish"
  french = "French"

  spanishGreeting = "Hola, "
  frenchGreeting = "Bonjour, "
  englishGreeting = "Hello, "
)

func Hello(name, lang string) string {
  if name == "" {
    name = "World"
  }
  return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (greeting string) {
  switch lang {
    case french:
      greeting = frenchGreeting
    case spanish:
      greeting = spanishGreeting
    default:
      greeting = englishGreeting
  }
  return 
  
}

func main(){
  fmt.Println(Hello("Chris", ""))
}
