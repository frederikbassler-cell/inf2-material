package aufgabe1

import "fmt"

func ExampleNode_LengthGreater5() {
	n := NewNode()
	fmt.Println(n.LengthGreater5())

	n.PushBack(1)
	fmt.Println(n.LengthGreater5())

	n.PushBack(2)
	n.PushBack(3)
	n.PushBack(4)
	n.PushBack(5)
	fmt.Println(n.LengthGreater5())

	n.PushBack(6)
	fmt.Println(n.LengthGreater5())

	n.PushBack(7)
	fmt.Println(n.LengthGreater5())

	// Output:
	// false
	// false
	// false
	// true
	// true
}



func ExampleNode_Contains() {
	n := NewNode()
	// Teste eine leere Liste
	fmt.Println(n.Contains(5))

	// Füge Elemente hinzu
	n.PushBack(10)
	n.PushBack(20)
	n.PushBack(30)

	// Suche nach Elementen, die existieren
	fmt.Println(n.Contains(20))
	fmt.Println(n.Contains(10))
	fmt.Println(n.Contains(30))

	// Suche nach einem Element, das nicht existiert
	fmt.Println(n.Contains(99))

	// Output:
	// false
	// true
	// true
	// true
	// false
}


func ExampleNode_CountEven() {
	n := NewNode()
	// Teste eine leere Liste
	fmt.Println(n.CountEven())

	// Füge eine ungerade Zahl hinzu
	n.PushBack(1)
	fmt.Println(n.CountEven())

	// Füge gerade und weitere ungerade Zahlen hinzu
	n.PushBack(2)
	n.PushBack(3)
	n.PushBack(4)
	fmt.Println(n.CountEven())

	// Füge noch eine gerade Zahl hinzu
	n.PushBack(6)
	fmt.Println(n.CountEven())

	// Output:
	// 0
	// 0
	// 2
	// 3
}


func ExampleNode_IsSorted() {
	n := NewNode()
	
	// Test 1: Eine leere Liste gilt als sortiert
	fmt.Println(n.IsSorted())

	// Test 2: Eine Liste mit nur einem Element gilt als sortiert
	n.PushBack(5)
	fmt.Println(n.IsSorted())

	// Test 3: Weitere Elemente hinzufügen, sodass sie sortiert bleibt
	n.PushBack(8)
	n.PushBack(12)
	n.PushBack(20)
	fmt.Println(n.IsSorted())

	// Test 4: Ein Element hinzufügen, das die Sortierung bricht
	n.PushBack(15)
	fmt.Println(n.IsSorted())

	// Output:
	// true
	// true
	// true
	// false
}