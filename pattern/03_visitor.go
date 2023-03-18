package pattern

type Visitor interface {
	GetDataFromPerson(p *Person) string
	GetDataFromPizzeria(p *Pizzeria) string
}

type CanBeVisited interface {
	Accept(v Visitor)
}

type Person struct {
	data string
}

type Pizzeria struct {
	data string
}

type Nalogovaya struct{}

func (n *Nalogovaya) GetDataFromPerson(p *Person) string {
	return p.data
}

func (n *Nalogovaya) GetDataFromPizzeria(p *Pizzeria) string {
	return p.data
}

func (p *Pizzeria) Accept(v Visitor) string {
	return v.GetDataFromPizzeria(p)
}

func (p *Person) Accept(v Visitor) string {
	return v.GetDataFromPerson(p)
}

//Паттерн Visitor относится к поведенческим паттернам уровня объекта.

//Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
//а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.
