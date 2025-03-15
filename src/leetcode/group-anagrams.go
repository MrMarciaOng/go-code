package leetcode
import "sort"
func groupAnagrams(strs []string) [][]string {
	//create a slice for the result
	//create a map to store the anagrams
	anagramMap := make(map[string][]string)

	//iterate through the strings
	for _, str := range strs {
		//convert string to byte slice so we can sort it
		bytes := []byte(str)
		//sort the bytes
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		//convert back to string for map key
		sortedStr := string(bytes)
		//add the original string to the map using sorted string as key
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// Convert map values to result slice
	result := make([][]string, 0, len(anagramMap))
	for _, v := range anagramMap {
		result = append(result, v)
	}
	return result
}