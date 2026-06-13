package aufgabe1

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei linkedlist.go finden Sie eine Implementierung einer einfach
//          verketteten Liste. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// LengthGreater5 gibt true zurück, wenn die Länge der Liste größer als 5 ist, sonst false.
func (n *Node) LengthGreater5() bool {
	count := n.count()

	if count > 5 {
		return true
	} else {
		return false
	}

}

func (n *Node) count() int {
	if n.IsEmpty() {
		return 0
	}
	return n.Next.count() + 1

}

func (n *Node) Contains(target int) bool {

	for !n.IsEmpty() {

		if n.Value == target {
			return true
		}
		n = n.Next
	}
	return false
}

func (n *Node) CountEven() int {
	count := 0
	for !n.IsEmpty() {

		if n.Value%2 == 0 {
			count++
		}
		n = n.Next
	}
	return count
}

// IsSorted prüft, ob die Elemente der Liste in aufsteigender Reihenfolge sortiert sind.
// Eine leere Liste oder eine Liste mit nur einem Element gilt ebenfalls als sortiert (true).

func (n *Node) IsSorted() bool {
	count := 0
	ctrue := 0
	if n.Next== nil {
		return true
	}
	if n.IsEmpty(){
		return true
	}
	for !n.Next.IsEmpty() {
	
		if n.Value < n.Next.Value {

			ctrue++
		}
		n = n.Next
		count++
	}

	if count == ctrue {
		return true
	}
	return false
}
