package algorithms

// As people come to the dance floor, they should be paired off as quickly as possible: man with woman, man with woman, all the way down the line. If several men arrive in a row, they should be paired in the order they came, and likewise if several women do. Maintain a queue of "spares" (men for whom you have no women yet, or vice versa), and pair them as appropriate.
// Input:
// []Person{
// 		John,
// 		Jane,
// 		Serena,
// 		Luna,
// 		Darien,
// 		Artemis,
// 		Rei,
// 		Nicolas,
// 	}
// Output:
// [][]Person{
// 		{John, Jane},
// 		{Serena, Darien},
// 		{Luna, Artemis},
// 		{Rei, Nicolas},
// 	}



type Person struct{
	name 		string
	gender 	string
}

const (
	MAN = "man"
	WOMAN = "woman"
)

var John Person = Person{name: "John", gender: MAN}
var Jane Person = Person{name: "Jane", gender: WOMAN}
var Serena Person = Person{name: "Serena", gender: WOMAN}
var Luna Person = Person{name: "Luna", gender: WOMAN}
var Darien Person = Person{name: "Darien", gender: MAN}
var Artemis Person = Person{name: "Artemis", gender: MAN}
var Rei Person = Person{name: "Rei", gender: WOMAN}
var Nicolas Person = Person{name: "Nicolás", gender: MAN}

func SquareDancePairing([]Person) [][]Person{
}