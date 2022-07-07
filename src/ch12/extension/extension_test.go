package extension

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("pet speaking")
}

func (p *Pet) SpeakTo(host string) {
	fmt.Println("pet SpeakTo" + host)
	p.Speak()
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("dog speaking")
}
func TestDog(t *testing.T) {
	// var pet Pet
	dog := Dog{}
	// := Dog
	dog.SpeakTo("DogTest")
}
