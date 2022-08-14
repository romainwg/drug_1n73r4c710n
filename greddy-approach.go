package main

// Greedy approach to find a single maximal clique in O(V^2) time complexity
// https://iq.opengenus.org/greedy-approach-to-find-single-maximal-clique/

/* from collections import defaultdict
import random

def find_single_clique(graph):
    clique = []
    vertices = list(graph.keys())
    rand = random.randrange(0, len(vertices), 1)
    clique.append(vertices[rand])
    for v in vertices:
        if v in clique:
            continue
        isNext = True
        for u in clique:
            if u in graph[v]:
                continue
            else:
                isNext = False
                break
        if isNext:
            clique.append(v)

    return sorted(clique)

graph = dict()
graph['A'] = ['B', 'C', 'E']
graph['B'] = ['A', 'C', 'D', 'F']
graph['C'] = ['A', 'B', 'D', 'F']
graph['D'] = ['C', 'E', 'B', 'F']
graph['E'] = ['A', 'D']
graph['F'] = ['B', 'C', 'D']

clique = find_single_clique(graph)
print('A maximal clique in the graph is: ', clique)
*/
