package main

import (
	"fmt"
	"os"
	"strings"
)

type drugMap map[string]int

// min returns minimum of the two input (int)
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// createMapFromDrug returns a drugMap (map[string]int)
// num of occurence (value) of each char (key)
func createMapFromDrug(drugName string) drugMap {
	var drugMap drugMap = make(drugMap)

	// Set drugName to lowercase
	drugName = strings.ToLower(drugName)

	for _, v := range []byte(drugName) {
		if _, found := drugMap[string(v)]; found {
			drugMap[string(v)]++
		} else {
			drugMap[string(v)] = 1
		}
	}

	return drugMap
}

// compareDrugMap checks the interaction between two drugs
// return true if the drugs have a good interaction
// otherwise false
func compareDrugMap(d1 drugMap, d2 drugMap) bool {

	var counter int = 0

	// Browse key from drugMap(1)
	for k1, v1 := range d1 {

		// Check if the key exist in drugMap(2)
		if v2, found := d2[k1]; found {
			// Inc. counter with the minimum occur of common char
			counter += min(v1, v2)
		}

		// If the num. of occurs >= 3 : bad interaction
		if counter >= 3 {
			return false
		}
	}

	return true
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	// scanner.Scan()
	// fmt.Sscan(scanner.Text(), &N)

	// For VS Code only

	/* JS code to get tests input from codingame
	const N = parseInt(readline());
	var in_t = []
	for (let i = 0; i < N; i++) {
		const s = "\""+readline()+"\"";
		in_t.push(s);
	}
	console.error(in_t.join(", "));
	*/

	/*
		// Test 1 / Result expected : 2
		var inputGenerated []string = []string{"Xanax", "Ativan", "Viagra"}

		// Test 2 / Result expected : 5
		var inputGenerated []string = []string{"Zubsolv", "Juluca", "Invega", "Herceptin", "Cinryze", "Sustol", "Zonegran", "Tecartus", "Alfamino", "Depakote"}

		// Test 3 / Result expected : 11
		var inputGenerated []string = []string{"Xeljanz", "Trisenox", "Afinitor", "Nityr", "Mylotarg", "Phesgo", "Ampyra", "Odomzo", "Delstrigo", "Enhertu", "Thiola", "Gelnique", "Nardil", "Cardura", "Cortef", "Gleevec", "Daypro", "Evista", "Myobloc", "Treanda", "Lumoxiti", "Bosulif", "Levoxyl", "Piqray", "Ciprodex", "Accupril", "Gilotrif", "Cipro", "Anafranil", "Gardasil", "Caplyta", "Imovax", "BeneFIX", "Fycompa", "Iclusig", "Thymoglobulin", "Pristiq", "Korlym", "Jardiance", "Botox"}
	*/
	var inputGenerated []string = []string{"Xeljanz", "Trisenox", "Afinitor", "Nityr", "Mylotarg", "Phesgo", "Ampyra", "Odomzo", "Delstrigo", "Enhertu", "Thiola", "Gelnique", "Nardil", "Cardura", "Cortef", "Gleevec", "Daypro", "Evista", "Myobloc", "Treanda", "Lumoxiti", "Bosulif", "Levoxyl", "Piqray", "Ciprodex", "Accupril", "Gilotrif", "Cipro", "Anafranil", "Gardasil", "Caplyta", "Imovax", "BeneFIX", "Fycompa", "Iclusig", "Thymoglobulin", "Pristiq", "Korlym", "Jardiance", "Botox"}

	N = len(inputGenerated)

	// Creation of a list to contain drugMap of each drugName
	var listDrugMap []drugMap = make([]drugMap, 0, N)

	listName := make([]string, 0)
	// https://goplay.tools/snippet/sCHnZ61L2-D / allNodes := make(map[int][]int)
	allNodes := make(map[string]Node)

	for i := 0; i < N; i++ {
		// scanner.Scan()
		// s := scanner.Text()
		s := inputGenerated[i]

		// Create actual drugMap
		drugMap_t := createMapFromDrug(strings.ToLower(s))

		// Append drugMap to list
		listDrugMap = append(listDrugMap, drugMap_t)

		// Creation of the node
		allNodes[s] = newNode(s)
		listName = append(listName, s)
	}

	// fmt.Fprintln(os.Stderr, "N", N)
	// fmt.Fprintln(os.Stderr, "listDrugMap", listDrugMap)

	// Liste des compatibles
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j == i {
				continue
			}

			if compareDrugMap(listDrugMap[i], listDrugMap[j]) {
				allNodes[listName[i]].neighbors[listName[j]] = allNodes[listName[j]]
			}
		}
	}

	// fmt.Fprintln(os.Stderr, "end prep", (time.Now().UnixNano()-now)/1000000)

	var counter int = 0

	potentialClique := make(map[string]Node)
	skipNodes := make(map[string]Node)

	fmt.Fprintln(os.Stderr, findCliques(potentialClique, allNodes, skipNodes, 0, &counter))

	// fmt.Fprintln(os.Stderr, "end calc", (time.Now().UnixNano()-now)/1000000)

	fmt.Println(counter) // Write answer to stdout
}
