package main

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel: string
}

type Chevy struct {
	ChevyModel : string
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main() {

}