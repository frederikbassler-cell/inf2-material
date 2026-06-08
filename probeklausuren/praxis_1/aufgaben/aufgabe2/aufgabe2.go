package aufgabe2

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei linkedlist.go finden Sie eine Implementierung einer einfach
//          verketteten Liste. Die zu implementierende Funktion ist eine Methode dieses Structs.
// HINWEIS: Zusätzlich zu den üblichen Feldern haben diese Listenelemente auch ihre Länge gespeichert.
//          Beim Entfernen eines Elements muss auch die Länge der Liste angepasst werden.
// MAX. PUNKTE: 10

// RemoveAt entfernt das Element an der Stelle index aus der Liste.
// Wenn index eine ungültige Position ist, soll die Liste unverändert bleiben.
func (n *Node) RemoveAt(index int) {
	if index == 0 {
		n.remove0(index)
	}
	if index > n.lenNode() {
		return
	}
	n.Remove(index)
	n.lengthadjustment()

}

func (n *Node) Remove(index int) {
	count := 0

	for !n.IsEmpty() {

		if count == index-1 {

			n.Next = n.Next.Next
			break
		}
		count++
		n = n.Next
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

func (n *Node) remove0(index int) {
	if index == 0 {
		n.Value = n.Next.Value
	}
	for !n.IsEmpty() {
		
		if index+1 == 1 {

			n.Next = n.Next.Next
			break
		}
	}
}
