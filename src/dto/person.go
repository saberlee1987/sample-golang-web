package dto

var persons []Person

type Person struct {
	FirstName string
	LastName  string
	Books     []string
}

func AddPerson(person Person) {
	persons = append(persons, person)
}
func GetPersons() []Person {
	return persons
}
