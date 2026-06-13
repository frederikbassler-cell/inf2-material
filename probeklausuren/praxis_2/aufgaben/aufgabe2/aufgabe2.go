package aufgabe2

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei linkedlist.go finden Sie eine Implementierung einer einfach
//          verketteten Liste. Die zu implementierende Funktion ist eine Methode dieses Structs.
// HINWEIS: Zusätzlich zu den üblichen Feldern haben diese Listenelemente auch ihre Länge gespeichert.
//          Beim Entfernen eines Elements muss auch die Länge der Liste angepasst werden.
// MAX. PUNKTE: 10

// RemoveAll entfernt alle Elemente mit dem Wert value aus der Liste.
// Wenn value in der Liste nicht vorkommt, soll die Liste unverändert bleiben.
// Die Funktion liefert den neuen Kopf der Liste zurück.
func (n *Node) RemoveAll(value int) *Node {
	if n.IsEmpty() {
		return nil
	}
	n.Remove(value)
	n.lengthadjustment()

	return n
}

func (n *Node) Remove(value int) {
	count := 0
	for !n.IsEmpty() {

		if count == 0 {
			if n.Value == value {
		n.Value = n.Next.Value
		n.Next = n.Next.Next
	}
		}

		if n.Next.Value == value {

			n.Next = n.Next.Next
			break
		}

		n = n.Next
		count++
	}
}
func (n *Node) lengthadjustment() {

	le := n.lenNode()
	for !n.IsEmpty() {
		n.Length = le
		n = n.Next
		le--
	}
}

func (n *Node) lenNode() int {
	Nodelen := 0
	for !n.IsEmpty() {
		n = n.Next
		Nodelen++
	}
	return Nodelen
}

