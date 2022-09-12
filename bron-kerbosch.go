package main

// type Node struct {
// 	name      string
// 	neighbors []Node
// }

// func (n Node) ToString() string {
// 	return n.name
// }

// func newNode(n string) Node {
// 	return Node{
// 		name:      n,
// 		neighbors: make([]Node, 0),
// 	}
// }

// https://goplay.tools/snippet/6UHTIwo-K7z
func findCliques(graph map[int][]int, potentialClique []int, remainingNodes []int, skipNodes []int, depth int, maxClique *int) int {

	if len(remainingNodes) == 0 && len(skipNodes) == 0 {
		// log.Println("Find clique :", potentialClique)
		// log.Println("Find clique")
		// for k := range potentialClique {
		// 	log.Println(k)
		// }
		if len(potentialClique) > *maxClique {
			*maxClique = len(potentialClique)
		}
		return 1
	}

	if len(remainingNodes)+len(potentialClique) < *maxClique {
		return 0
	}

	// found_cliques = 0
	foundCliques := 0

	// for node in remaining_nodes:
	for _, vnode := range remainingNodes {

		// new_potential_clique = potential_clique + [node]
		newPotentialClique := make([]int, 0)
		newPotentialClique = append(potentialClique, vnode)

		// log.Println("newPotentialClique :", vnode, potentialClique, newPotentialClique)

		// new_remaining_nodes = [n for n in remaining_nodes if n in node.neighbors]
		newRemainingNodes := make([]int, 0)
		for _, v := range remainingNodes {
			for _, k0 := range graph[vnode] {
				if k0 == v {
					newRemainingNodes = append(newRemainingNodes, k0)
					break
				}
			}
		}

		// new_skip_list = [n for n in skip_nodes if n in node.neighbors]
		newSkipNodes := make([]int, 0)
		for _, v := range skipNodes {
			for _, k0 := range graph[vnode] {
				if k0 == v {
					newSkipNodes = append(newSkipNodes, k0)
					break
				}
			}
		}

		// found_cliques += find_cliques(new_potential_clique, new_remaining_nodes, new_skip_list, depth + 1)
		foundCliques += findCliques(graph, newPotentialClique, newRemainingNodes, newSkipNodes, depth+1, maxClique)
	}
	return foundCliques

}
