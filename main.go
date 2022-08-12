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
	/* Test 3
	var inputGenerated []string = []string{"Xeljanz", "Trisenox", "Afinitor", "Nityr", "Mylotarg", "Phesgo", "Ampyra", "Odomzo", "Delstrigo", "Enhertu", "Thiola", "Gelnique", "Nardil", "Cardura", "Cortef", "Gleevec", "Daypro", "Evista", "Myobloc", "Treanda", "Lumoxiti", "Bosulif", "Levoxyl", "Piqray", "Ciprodex", "Accupril", "Gilotrif", "Cipro", "Anafranil", "Gardasil", "Caplyta", "Imovax", "BeneFIX", "Fycompa", "Iclusig", "Thymoglobulin", "Pristiq", "Korlym", "Jardiance", "Botox"}
	*/
	var inputGenerated []string = []string{"Xanax", "Ativan", "Viagra"}
	N = len(inputGenerated)

	// Creation of a list to contain drugMap of each drugName
	var listDrugMap []drugMap = make([]drugMap, 0, N)

	for i := 0; i < N; i++ {
		// scanner.Scan()
		// s := scanner.Text()
		s := inputGenerated[i]

		// Create actual drugMap
		drugMap_t := createMapFromDrug(strings.ToLower(s))

		// Append drugMap to list
		listDrugMap = append(listDrugMap, drugMap_t)
	}

	fmt.Fprintln(os.Stderr, "N", N)
	fmt.Fprintln(os.Stderr, "listDrugMap", listDrugMap)

	var counter int = 0

	var listCompatible [][]int = make([][]int, 0, N)

	// Liste des compatibles
	for i := 0; i < N; i++ {

		var array_t []int = make([]int, 0, N-1)
		listCompatible = append(listCompatible, array_t)

		for j := 0; j < N; j++ {
			if j == i {
				continue
			}

			if compareDrugMap(listDrugMap[i], listDrugMap[j]) {
				listCompatible[i] = append(listCompatible[i], j)
			}
		}
	}

	fmt.Fprintln(os.Stderr, listCompatible, "listCompatible")

	for i := 0; i < N; i++ {
		var i_t int = maxInteractionInSet(listCompatible, listCompatible[i])
		fmt.Fprintln(os.Stderr, i, listCompatible[i], i_t, "maxInSet")
	}

	fmt.Println(counter) // Write answer to stdout
}

func maxInteractionInSet(listCompatible [][]int, set []int) int {
	var max int = 0

	// Parcours du set
	for _, v := range set {

		var count int = 0

		// Parcours des éléments de l'élément en cours du set initial
		for _, subElemt := range listCompatible[v] {
			if in(subElemt, set) {
				count++
			}
		}

		if count > max {
			max = count
		}

	}

	return max
}

func in(x int, list []int) bool {
	for _, v := range list {
		if v == x {
			return true
		}
	}
	return false
}
