package main

type Graphe []Vertex

type Vertex struct {
	name string

	// Tarjan variables
	num           int
	numDefined    bool
	numAccessible int
	dansP         bool

	edges []*Vertex
}

func (g Graphe) setEdge(in int, out int) {
	g[in].edges = append(g[in].edges, &g[out])
}

// https://fr.wikipedia.org/wiki/Algorithme_de_Tarjan
func tarjan(g Graphe) [][]*Vertex {

	// num := 0
	var num int = 0
	// P := pile vide
	var p []*Vertex = make([]*Vertex, 0)
	// partition := ensemble vide
	var partition [][]*Vertex = make([][]*Vertex, 0)

	// Recursive anonymous function example
	// https://goplay.tools/snippet/NY027ogEn2S
	var parcours func(v *Vertex)

	parcours = func(v *Vertex) {
		// v.num := num
		v.num = num
		v.numDefined = true
		// v.numAccessible := num
		v.numAccessible = num
		// num := num + 1
		num += 1
		// P.push(v), v.dansP := oui
		p = append(p, v)
		v.dansP = true

		// // Parcours récursif
		for j := 0; j < len(v.edges); j++ { // pour chaque w successeur de v

			// si w.num n'est pas défini
			if !v.edges[j].numDefined {

				//   parcours(w)
				parcours(v.edges[j])

				//   v.numAccessible := min(v.numAccessible, w.numAccessible)
				v.numAccessible = min(v.numAccessible, v.edges[j].numAccessible)

			} else if v.edges[j].dansP { // sinon si w.dansP = oui

				// v.numAccessible := min(v.numAccessible, w.num)
				v.numAccessible = min(v.numAccessible, v.edges[j].num)
			}
		}

		if v.numAccessible == v.num { // si v.numAccessible = v.num
			// // v est une racine, on calcule la composante fortement connexe associée
			// C := ensemble vide
			var c []*Vertex
			var flag bool = true
			// tant que w différent de v
			for flag {
				// w := P.pop(), w.dansP := non
				if len(p) == 0 {
					break
				}
				w := p[len(p)-1]
				w.dansP = false
				if len(p) > 0 {
					p = p[:len(p)-1]
				}
				// ajouter w à C
				c = append(c, w)
				if v.name != w.name {
					flag = false
				}
			}

			// ajouter C à partition
			partition = append(partition, c)
		}

	}

	// Parcours des sommets du graphe
	for i := 0; i < len(g); i++ {
		// Si num n'est pas défini
		if !g[i].numDefined {
			parcours(&g[i])
		}
	}

	return partition

}
