package models

type Person struct {
	Name string `json:"name" binding:"required"`
	Pw   string `json:"pw" binding:"required"`
}

// Constructor is make 'Person' object
func Constructor(name string, pw string) *Person {
	person := Person{Name: name, Pw: pw}

	return &person
}

// GetPerson info of person
func GetPerson(p *Person) (string, string) {
	name := p.Name
	pw := p.Pw

	return name, pw
}
