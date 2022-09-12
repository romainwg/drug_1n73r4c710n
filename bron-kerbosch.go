package main

type Node struct {
	name      string
	neighbors map[string]Node
}

func (n Node) ToString() string {
	return n.name
}

func newNode(n string) Node {
	return Node{
		name:      n,
		neighbors: make(map[string]Node, 0),
	}
}

// https://goplay.tools/snippet/6UHTIwo-K7z
func findCliques(potentialClique map[string]Node, remainingNodes map[string]Node, skipNodes map[string]Node, depth int, maxClique *int) int {

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

	// No need to look any further : we search bigger clique only
	if len(remainingNodes)+len(potentialClique) < *maxClique {
		return 0
	}

	// found_cliques = 0
	foundCliques := 0

	// for node in remaining_nodes:
	for knode, vnode := range remainingNodes {

		// new_potential_clique = potential_clique + [node]
		newPotentialClique := make(map[string]Node, 0)
		for k, v := range potentialClique {
			newPotentialClique[k] = v
		}
		newPotentialClique[knode] = vnode

		// new_remaining_nodes = [n for n in remaining_nodes if n in node.neighbors]
		newRemainingNodes := make(map[string]Node, 0)
		for k, v := range remainingNodes {
			if _, b := vnode.neighbors[k]; b {
				newRemainingNodes[k] = v
			}
		}

		// new_skip_list = [n for n in skip_nodes if n in node.neighbors]
		newSkipNodes := make(map[string]Node, 0)
		for k, v := range skipNodes {
			if _, b := vnode.neighbors[k]; b {
				newSkipNodes[k] = v
			}
		}

		// found_cliques += find_cliques(new_potential_clique, new_remaining_nodes, new_skip_list, depth + 1)
		foundCliques += findCliques(newPotentialClique, newRemainingNodes, newSkipNodes, depth+1, maxClique)

		// remaining_nodes.remove(node)
		delete(remainingNodes, knode)

		// skip_nodes.append(node)
		skipNodes[knode] = vnode
	}
	return foundCliques

}
