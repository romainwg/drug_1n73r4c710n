package main

import (
	"reflect"
	"testing"
)

func TestCreateMapFromDrug(t *testing.T) {

	var err error = nil

	var expectedMap drugMap
	var drugMapping drugMap
	var drug string
	var eq bool
	var want bool

	// Test 1 : ok
	drug = "XanaxT"
	expectedMap = make(map[string]int)
	expectedMap["a"] = 2
	expectedMap["x"] = 2
	expectedMap["n"] = 1
	expectedMap["t"] = 1
	want = true

	drugMapping = createMapFromDrug(drug)

	// Deep compare ; type must be the same
	// for this function : map[string]int != drugMap
	eq = reflect.DeepEqual(drugMapping, expectedMap)
	if eq != want || err != nil {
		t.Fatalf(`createMapFromDrug("%s") = %v, %v, want match for %v, nil`, drug, drugMapping, err, expectedMap)
	}

	// Test 2 : too much letter in expected
	drug = "XanaxTd"
	expectedMap = make(map[string]int)
	expectedMap["a"] = 2
	expectedMap["x"] = 2
	expectedMap["n"] = 1
	expectedMap["t"] = 1
	want = false

	drugMapping = createMapFromDrug(drug)

	eq = reflect.DeepEqual(drugMapping, expectedMap)
	if eq != want || err != nil {
		t.Fatalf(`createMapFromDrug("%s") = %v, %v, want NOT match for %v, nil`, drug, drugMapping, err, expectedMap)
	}

}

func TestCompareDrugMap(t *testing.T) {

	var err error = nil

	var drug1, drug2 string
	var drugMap1, drugMap2 drugMap
	var want, eq bool

	// Test 1 : ok - good interaction
	drug1 = "Xanax"
	drug2 = "Viagra"
	want = true

	drugMap1 = createMapFromDrug(drug1)
	drugMap2 = createMapFromDrug(drug2)

	eq = compareDrugMap(drugMap1, drugMap2)

	if eq != want || err != nil {
		t.Fatalf(`compareDrugMap(%v, %v) = %t, %v, want match for %t, nil`, drugMap1, drugMap2, eq, err, want)
	}

	// Test 2 : ko - bad interaction
	drug1 = "Ativan"
	drug2 = "Viagra"
	want = false

	drugMap1 = createMapFromDrug(drug1)
	drugMap2 = createMapFromDrug(drug2)

	eq = compareDrugMap(drugMap1, drugMap2)

	if eq != want || err != nil {
		t.Fatalf(`compareDrugMap(%v, %v) = %t, %v, want match for %t, nil`, drugMap1, drugMap2, eq, err, want)
	}

	// Test 3 : ok - good interaction
	drug1 = "aaAabbBbcd"
	drug2 = "abefgijklmopqrstuvwxyz"
	want = true

	drugMap1 = createMapFromDrug(drug1)
	drugMap2 = createMapFromDrug(drug2)

	eq = compareDrugMap(drugMap1, drugMap2)

	if eq != want || err != nil {
		t.Fatalf(`compareDrugMap(%v, %v) = %t, %v, want match for %t, nil`, drugMap1, drugMap2, eq, err, want)
	}

	// Test 4 : ko - bad interaction
	drug1 = "ABCdef"
	drug2 = "abc"
	want = false

	drugMap1 = createMapFromDrug(drug1)
	drugMap2 = createMapFromDrug(drug2)

	eq = compareDrugMap(drugMap1, drugMap2)

	if eq != want || err != nil {
		t.Fatalf(`compareDrugMap(%v, %v) = %t, %v, want match for %t, nil`, drugMap1, drugMap2, eq, err, want)
	}
}
