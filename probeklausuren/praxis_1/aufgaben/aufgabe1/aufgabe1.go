package aufgabe1

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei linkedlist.go finden Sie eine Implementierung einer einfach
//          verketteten Liste. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// LengthGreater5 gibt true zurück, wenn die Länge der Liste größer als 5 ist, sonst false.
func (n *Node) LengthGreater5() bool {
	count :=n.count() 

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
