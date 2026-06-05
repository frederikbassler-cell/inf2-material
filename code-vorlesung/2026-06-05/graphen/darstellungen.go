package graphen

// Kantenliste
type EdgeList struct {
	Edges [][3]int // Liste von Kanten mit Start, Ziel und Gewicht
	Data  map[int]string
}

// Adjazenzmatrix
type AdjacencyMatrix struct {
	Edges [][]int // Tabelle von Gewichten
	Data  map[int]string
}

// Adjazentliste
type AdjacencyList struct {
	Edges map[int][]int // Abbildung von Knoten-ID auf deren Listen von Nachbarn
	Data  map[int]string
}

// InlineGraph
type InlineGraph struct {
	Vertices map[int]Vertex // Abbildung von Knoten-ID auf Knoten-Objekte
}

type Vertex struct {
	Data       string
	Neighbours []int
}
