package main

import "fmt"

type animal interface {
	mover() string
}

type dog struct {
}

type cat struct {
}

type bird struct {
}

func (dog) mover() string {
	return fmt.Sprintln("I'm a dog and I'm running")
}
func (cat) mover() string {
	return fmt.Sprintln("I'm a cat and I'm running fast")
}
func (bird) mover() string {
	return fmt.Sprintln("I'm a bird and I'm flying")
}
func moverAnimal(a animal) {
	fmt.Println(a.mover())
}
func main() {

	d := dog{}
	c := cat{}
	b := bird{}

	moverAnimal(d)
	moverAnimal(c)
	moverAnimal(b)

}
