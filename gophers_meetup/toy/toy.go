package toy

import "fmt"

type Toy struct {
	Name string
	Weight float32
	onHand int
	sold int
}

func New(name string, weight float32) *Toy {
	fmt.Println("in Toy:New")
	return &Toy{Name: name, Weight:weight}
}

func (t *Toy) Update(o int, s int)  {
	t.onHand += o
	t.sold += s
}


func (t *Toy) Get() (string, float32, int,int) {
	return t.Name, t.Weight, t.onHand, t.sold
}