package person

type Person struct {
	name string
	pw   string
}

// Constructor is make 'Person' object
func Constructor(name string, pw string) *Person {
	person := Person{name: name, pw: pw}

	return &person
}

// GetPerson info of person
func GetPerson(p *Person) (string, string) {
	name := p.name
	pw := p.pw

	return name, pw
}
