package leetcode

func FindJudge(n int, trust [][]int) int {
	//town judge have trust issue
	//everyone else trust judge besdies judge himself

	//build an map of person by index and couunt of how many ppl trust that person
	if n == 1 {
		return 1
	}
	//shit count start at 1 theres not person zero
	var trustMap = make(map[int]int)
	var trustIssue = make([]int, n+1)
	for i := range trustIssue {
		trustIssue[i] = 1
	}

	for _, value := range trust {
		var a, b = value[0], value[1]
		trustIssue[a] = 0 //set the person who trust to 0
		//inreaset the person trust count
		trustMap[b] = trustMap[b] + 1
	}

	for i, v := range trustMap {
		//everyone else the person
		if v == n-1 {
			//the person has trust issue
			if trustIssue[i] == 1 {
				return i
			}
		}
	}
	return -1
}
