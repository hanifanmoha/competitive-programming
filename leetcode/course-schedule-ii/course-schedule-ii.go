package coursescheduleii

func findOrder(numCourses int, prerequisites [][]int) []int {

	prereqMap := make(map[int]map[int]bool)
	for _, prereq := range prerequisites {
		c := prereq[0]
		p := prereq[1]
		if _, ok := prereqMap[c]; !ok {
			prereqMap[c] = make(map[int]bool)
		}
		if _, ok := prereqMap[c][p]; !ok {
			prereqMap[c][p] = true
		}
	}

	// wave function collapse
	flag := make([]bool, numCourses)

	for i := range flag {
		flag[i] = false
	}

	res := make([]int, 0)
	for _ = range numCourses {
		for c := range numCourses {
			prereqC := prereqMap[c]
			isAllPrerequisitesDone := true
			for p := range prereqC {
				isAllPrerequisitesDone = isAllPrerequisitesDone && flag[p]
			}
			if !flag[c] && isAllPrerequisitesDone {
				res = append(res, c)
				flag[c] = true
			}
		}
	}

	if len(res) < numCourses {
		return []int{}
	}

	return res
}
