package aufgabe5

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei bintree.go finden Sie eine Implementierung eines binären Suchbaumes.
//          Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

func (n *Node) Size() int {

	if n.IsEmpty() {
		return 0
	}

	return n.Right.Size() + 1 + n.Left.Size()

}
