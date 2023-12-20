package main

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func cntNext(root *Employee, eMap map[int]*Employee) int {
	importance := root.Importance
	for _, id := range root.Subordinates {
		importance += cntNext(eMap[id], eMap)
	}
	return importance
}

func getImportance(employees []*Employee, id int) int {
	eMap := map[int]*Employee{}
	for _, employee := range employees {
		eMap[employee.Id] = employee
	}
	return cntNext(eMap[id], eMap)
}

func main() {
	fmt.Println(getImportance([]*Employee{{1, 5, []int{2, 3}},
		{2, 3, []int{}}, {3, 3, []int{}}}, 1))
}
