package main

import (
	"amartha/models"
	"container/list"
	"fmt"
)

func main() {
}

func findPath(m int, n int, input [][]int) bool {

	visited := make(map[string]*models.Location)
	//stack is to store routes
	stack := list.New()

	//initialize src as the initial location and dest as the target location
	src := &models.Location{X: 0, Y: 0}
	dest := &models.Location{X: m - 1, Y: n - 1}

	//push the initial root route to stack
	stack.PushBack(src)

	var findNextPath func(src, dest *models.Location) (bool, bool)
	var checkIfVisitedAndFindNextPath func(current, dest *models.Location) (bool, bool)

	checkIfVisitedAndFindNextPath = func(current, dest *models.Location) (bool, bool) {
		//check if this current route has been visited
		//if already visited, find another route
		if _, ok := visited[current.String()]; !ok {

			//marked the visited route
			visited[current.String()] = current
			//push current route to stack
			stack.PushBack(current)

			//check if we have come to the dest
			if current.Compare(dest) {
				return true, true
			}

			//if not in the dest, find next path
			if finish, status := findNextPath(current, dest); finish {
				return true, status
			}
		}
		return false, false
	}

	findNextPath = func(src, dest *models.Location) (bool, bool) {
		//check coundaries
		checkBoundaries := src.CheckBoundaries(m, n)

		if checkBoundaries.InboundDown {
			current := src.Down()
			if input[current.X][current.Y] == 0 {
				if finish, status := checkIfVisitedAndFindNextPath(current, dest); finish {
					return true, status
				}
			}
		}

		if checkBoundaries.InboundRight {
			current := src.Right()
			if input[current.X][current.Y] == 0 {
				if finish, status := checkIfVisitedAndFindNextPath(current, dest); finish {
					return true, status
				}
			}
		}

		if checkBoundaries.InboundUp {
			current := src.Up()
			if input[current.X][current.Y] == 0 {
				if finish, status := checkIfVisitedAndFindNextPath(current, dest); finish {
					return true, status
				}
			}
		}

		if checkBoundaries.InboundLeft {
			current := src.Left()
			if input[current.X][current.Y] == 0 {
				if finish, status := checkIfVisitedAndFindNextPath(current, dest); finish {
					return true, status
				}
			}
		}

		//pop the last route since we must go back to find another route
		stack.Remove(stack.Back())

		return false, false
	}

	_, status := findNextPath(src, dest)

	//print the route we found if any
	for e := stack.Front(); e != nil; e = e.Next() {
		v := e.Value
		fmt.Printf("%v -> ", v)
	}
	fmt.Println(status)

	return status
}
