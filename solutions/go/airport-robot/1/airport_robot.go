package airportrobot

import (
	"fmt"
)

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

func SayHello(name string, greeter Greeter) string{
    return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {}
func (i Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}
func (i Italian) LanguageName() string {
    return "Italian"
}

type Portuguese struct {}
func (p Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}
func (p Portuguese) LanguageName() string {
    return "Portuguese"
}
