package main

import "fmt"

func _structs() {
	// Structs
	// Formas de criar suas proprias estruturas de dados
	// Pode ser feita de acordo com as suas necessidades
	// Podemos usar vários tipos diferentes

	// Type, nome da estrutura e struct.
	type Person struct {
		name     string
		age      int
		isEditor bool
		hobbies  []string
	}
	type Profile struct {
		Person
		email    string
		password string
	}

	var profile = Profile{
		Person: Person{
			name:     "Rumo",
			age:      25,
			isEditor: true,
			hobbies:  []string{"Games", "Prog"},
		},
		email:    "Rumo@email",
		password: "Rumo@password",
	}
	fmt.Print(profile)

	persons := []Person{
		{name: "Rômulo", age: 25, isEditor: true, hobbies: []string{"Games", "Prog"}},
		{name: "Gabis", age: 25, isEditor: true, hobbies: []string{"Tiktok", "Instagram"}},
	}
	// fmt.Println(persons)

	// students := map[string][]Person{}
	// students["prog"] = persons
	// fmt.Println(students)

	var students = map[string][]Person{
		"prog": persons,
	}

	var students_of_courses = map[string][]Person{
		"Music": {
			{name: "Ludovic", age: 25, isEditor: true, hobbies: []string{"Flauta", "Violão"}},
			{name: "Lindomar", age: 25, isEditor: true, hobbies: []string{"Piano", "Vilãotchelo"}},
		},
		"Cinema": {
			{name: "Roubert", age: 25, isEditor: true, hobbies: []string{"Filmes", "comédia"}},
		}}
	fmt.Println(students_of_courses)

	fmt.Println(students)

}
