package main

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func greedy(gas []int, cost []int) int {
	n := len(gas)

	idx := 0
	max := 0

	for i := range n {
		r := gas[i] - cost[i]
		if r > max {
			max = r
			idx = i
		}
	}
	return idx
}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	if sum(gas) < sum(cost) {
		return -1
	}

	greedyStart := greedy(gas, cost)

	// fmt.Println("GREEDY START", greedyStart)

	for i := range n {

		start := (i + greedyStart) % n
		totalGas := 0

		for step := range n {

			pos := (start + step) % n
			nextPos := (start + step + 1) % n

			totalGas += gas[pos]
			totalGas -= cost[pos]

			// fmt.Println(pos, totalGas, nextPos, start)

			if totalGas < 0 {
				break
			}

			if nextPos == start {
				return start
			}

		}

		// fmt.Println("=====")

	}

	return -1
}
