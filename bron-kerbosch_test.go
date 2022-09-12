package main

// func TestFindCliques(t *testing.T) {

// 	allNodes := make(map[string]Node)

// 	str_t := []string{"1", "2", "3", "4"}

// 	for _, v := range str_t {
// 		allNodes[v] = newNode(v)
// 	}

// 	allNodes["1"].neighbors["2"] = allNodes["2"]
// 	allNodes["1"].neighbors["3"] = allNodes["3"]

// 	allNodes["2"].neighbors["1"] = allNodes["1"]
// 	allNodes["2"].neighbors["3"] = allNodes["3"]

// 	allNodes["3"].neighbors["1"] = allNodes["1"]
// 	allNodes["3"].neighbors["2"] = allNodes["2"]
// 	allNodes["3"].neighbors["4"] = allNodes["4"]

// 	allNodes["4"].neighbors["3"] = allNodes["3"]

// 	potentialClique := make(map[string]Node)
// 	skipNodes := make(map[string]Node)

// 	var res int = 0

// 	fmt.Println(findCliques(potentialClique, allNodes, skipNodes, 0, &res))
// 	fmt.Println(res)
// }
