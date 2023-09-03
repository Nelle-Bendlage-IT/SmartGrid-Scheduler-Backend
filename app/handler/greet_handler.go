package handler

import (
	"fmt"

	greetService "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/domain/greet"
)

type Greet struct {
	service greetService.GreetService
}

func NewGreetService(service greetService.GreetService) Greet {
	return Greet{service: service}
}

func (g Greet) HandleGetGreet(name string) string {
	fmt.Println(name)
	return g.service.GetGreeting(name)
}
