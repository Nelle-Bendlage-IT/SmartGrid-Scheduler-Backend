package greetService

import "fmt"

type Service struct {
}

type GreetService interface {
	GetGreeting(name string) string
}

func (svc Service) GetGreeting(name string) string {
	return "Hello " + fmt.Sprint(name)
}
