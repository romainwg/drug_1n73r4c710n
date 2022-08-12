package main

import (
	"log"
	"strconv"
	"testing"
)

func TestTarjanA(t *testing.T) {
	// https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Directed_graph_no_background.svg/330px-Directed_graph_no_background.svg.png
	// Test
	// https://goplay.tools/snippet/MqrEI76bBzH

	var g Graphe = Graphe(make([]Vertex, 0))

	for i := 0; i < 4; i++ {
		g = append(g, Vertex{
			name:  strconv.Itoa(i + 1),
			edges: make([]*Vertex, 0),
		})
	}

	// 1 -> 2
	g[0].edges = append(g[0].edges, &g[1])
	// 1 -> 3
	g[0].edges = append(g[0].edges, &g[2])

	// 3 -> 2
	g[2].edges = append(g[2].edges, &g[1])
	// 3 -> 4
	g[2].edges = append(g[2].edges, &g[3])

	// 4 -> 3
	g[3].edges = append(g[3].edges, &g[2])

	for i := 0; i < len(g); i++ {
		log.Printf("%d ; %p ; %v \n", i+1, &g[i], g[i])
	}

	res := tarjan(g)

	for _, v := range res {
		str_t := ""
		for _, vertex := range v {
			str_t += vertex.name + " ; "
		}
		log.Println(str_t)
	}

}

func TestTarjanC(t *testing.T) {
	// https://upload.wikimedia.org/wikipedia/commons/6/60/Tarjan%27s_Algorithm_Animation.gif
	// Test
	// https://goplay.tools/snippet/MqrEI76bBzH

	var g Graphe = Graphe(make([]Vertex, 0))

	for i := 0; i < 3; i++ {
		g = append(g, Vertex{
			name:  strconv.Itoa(i + 1),
			edges: make([]*Vertex, 0),
		})
	}

	// 1 -> 2
	g.setEdge(0, 1)
	// 2 -> 3
	g.setEdge(1, 2)
	// 3 -> 1
	g.setEdge(2, 0)

	for i := 0; i < len(g); i++ {
		log.Printf("%d ; %p ; %v \n", i+1, &g[i], g[i])
	}

	res := tarjan(g)

	for _, v := range res {
		str_t := ""
		for _, vertex := range v {
			str_t += vertex.name + " ; "
		}
		log.Println(str_t)
	}

}

func TestTarjanD(t *testing.T) {

	var g Graphe = Graphe(make([]Vertex, 0))

	for i := 0; i < 2; i++ {
		g = append(g, Vertex{
			name:  strconv.Itoa(i + 1),
			edges: make([]*Vertex, 0),
		})
	}

	// 1 -> 1
	g.setEdge(0, 0)

	// 2 -> 2
	g.setEdge(1, 1)

	for i := 0; i < len(g); i++ {
		log.Printf("%d ; %p ; %v \n", i+1, &g[i], g[i])
	}

	res := tarjan(g)

	for _, v := range res {
		str_t := ""
		for _, vertex := range v {
			str_t += vertex.name + " ; "
		}
		log.Println(str_t)
	}

}

func TestTarjanE(t *testing.T) {

	var g Graphe = Graphe(make([]Vertex, 0))

	for i := 0; i < 3; i++ {
		g = append(g, Vertex{
			name:  strconv.Itoa(i + 1),
			edges: make([]*Vertex, 0),
		})
	}

	// 1 -> 1
	g.setEdge(0, 0)
	// 1 -> 3
	g.setEdge(0, 2)

	// 2 -> 2
	g.setEdge(1, 1)

	for i := 0; i < len(g); i++ {
		log.Printf("%d ; %p ; %v \n", i+1, &g[i], g[i])
	}

	res := tarjan(g)

	log.Println(res)

	for _, v := range res {
		str_t := ""
		for _, vertex := range v {
			str_t += vertex.name + " ; "
		}
		log.Println(str_t)
	}

}

func TestTarjanF(t *testing.T) {

	var g Graphe = Graphe(make([]Vertex, 0))

	for i := 0; i < 4; i++ {
		g = append(g, Vertex{
			name:  strconv.Itoa(i + 1),
			edges: make([]*Vertex, 0),
		})
	}

	// 1 -> 2
	g.setEdge(0, 1)
	// 2 -> 1
	g.setEdge(1, 0)

	// 2 -> 3
	g.setEdge(1, 2)
	// 3 -> 2
	g.setEdge(2, 1)

	// 1 -> 3
	g.setEdge(0, 2)
	// 3 -> 1
	g.setEdge(2, 0)

	// 3 -> 4
	g.setEdge(2, 3)
	// 4 -> 3
	g.setEdge(3, 2)

	for i := 0; i < len(g); i++ {
		log.Printf("%d ; %p ; %v \n", i+1, &g[i], g[i])
	}

	res := tarjan(g)

	log.Println(res)

	for _, v := range res {
		str_t := ""
		for _, vertex := range v {
			str_t += vertex.name + " ; "
		}
		log.Println(str_t)
	}

}
