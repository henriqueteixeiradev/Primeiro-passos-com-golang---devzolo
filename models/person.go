package models

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) SetName(name string) {
	p.Name = name
}
