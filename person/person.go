package person

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) SetFirstName(newName string) {
	p.FirstName = newName
}