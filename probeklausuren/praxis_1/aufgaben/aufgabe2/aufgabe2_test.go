package aufgabe2

import "fmt"

// Test für das Entfernen von gültigen Positionen, aber nicht am Anfang der Liste.
func ExampleNode_RemoveAt_valid_notfirst() {
	n := NewNode() // []
	PrintValuesAndLengths(n)

	n.PushBack(42)
	n.PushBack(25)
	n.PushBack(17)
	n.PushBack(53)
	n.PushBack(38)
	PrintValuesAndLengths(n) // [42 25 17 53 38]

	n.RemoveAt(2) // [42 25 53 38]
	PrintValuesAndLengths(n)

	n.RemoveAt(1) // [42 53 38]
	PrintValuesAndLengths(n)

	n.RemoveAt(0) // [53 38]
	PrintValuesAndLengths(n)

	// Output:
	// Values: []
	// Lengths: []
	//
	// Values: [42 25 17 53 38]
	// Lengths: [5 4 3 2 1]
	//
	// Values: [42 25 53 38]
	// Lengths: [4 3 2 1]
	//
	// Values: [42 53 38]
	// Lengths: [3 2 1]
	//
	// Values: [53 38]
	// Lengths: [2 1]

}

// Test für das Entfernen von ungültigen Positionen.
func ExampleNode_RemoveAt_invalid() {
	n := NewNode() // []
	n.PushBack(42)
	n.PushBack(38) // [42 38]

	n.RemoveAt(-1) // [42 38]
	PrintValuesAndLengths(n)

	n.RemoveAt(3) // [42 38]
	PrintValuesAndLengths(n)

	// Output:
	// Values: [42 38]
	// Lengths: [2 1]
	//
	// Values: [42 38]
	// Lengths: [2 1]
}

// Test für das Entfernen der ersten Position.
func ExampleNode_RemoveAt_valid_first() {
	n := NewNode() // []
	n.PushBack(42)
	n.PushBack(38) // [42 38]

	n.RemoveAt(0) // [38]
	PrintValuesAndLengths(n)

	// Output:
	// Values: [38]
	// Lengths: [1]
}

/* Hilfsfunktion für den Test */
func PrintValuesAndLengths(n *Node) {
	values := []int{}
	lengths := []int{}
	for !n.IsEmpty() {
		values = append(values, n.Value)
		lengths = append(lengths, n.Length)
		n = n.Next
	}
	fmt.Printf("Values: %v\nLengths: %v\n\n", values, lengths)
}



func ExampleNode_InsertAt() {
	n := NewNode()
	n.PushBack(10) // Index 0
	n.PushBack(30) // Index 1

	fmt.Println("Länge vor Insert:", n.Length)

	// Füge 20 an Index 1 ein (zwischen 10 und 30)
	n.InsertAt(1, 20)
	
  	fmt.Println("Länge nach Insert:", n.Length)
	
	// Überprüfen der Werte anhand der Struktur
	fmt.Println("Index 0:", n.Value)
	fmt.Println("Index 1:", n.Next.Value)
	fmt.Println("Index 2:", n.Next.Next.Value)

	// Ungültiger Index (ändert nichts)
	n.InsertAt(99, 50)
	fmt.Println("Länge nach falschem Insert:", n.Length)

	// Output:
	// Länge vor Insert: 2
	// Länge nach Insert: 3
	// Index 0: 10
	// Index 1: 20
	// Index 2: 30
	// Länge nach falschem Insert: 3
}


func ExampleNode_RemoveLast() {
	n := NewNode()
	
	// Liste mit drei Elementen füllen
	n.PushBack(10)
	n.PushBack(20)
	n.PushBack(30)

	fmt.Println("Startlänge:", n.Length)

	// 1. RemoveLast (entfernt die 30)
	n.RemoveLast()
	fmt.Println("Länge nach 1. Remove:", n.Length)
	fmt.Println("Neues letztes Element:", n.Next.Value) // An Index 1 (n.Next) sollte jetzt die 20 liegen

	// Restliche Elemente entfernen
	n.RemoveLast()
	n.RemoveLast()
	fmt.Println("Länge nach Leerung:", n.Length)
	fmt.Println("Ist Liste leer?", n.IsEmpty())

	// Auf bereits leerer Liste aufrufen (darf nicht abstürzen)
	n.RemoveLast()
	fmt.Println("Länge wenn schon leer:", n.Length)

	// Output:
	// Startlänge: 3
	// Länge nach 1. Remove: 2
	// Neues letztes Element: 20
	// Länge nach Leerung: 0
	// Ist Liste leer? true
	// Länge wenn schon leer: 0
}