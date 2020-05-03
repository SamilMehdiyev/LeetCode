package solutions

// DestCity function is a solution for the
// Question #1 - Destination City
// from leetcode.com
func DestCity(paths [][]string) string {

	roadMap := make(map[string]string)

	for i := 0; i < len(paths); i++ {
		source := paths[i][0]
		destination := paths[i][1]

		roadMap[source] = destination
		if _, found := roadMap[destination]; found {
			continue
		}
		roadMap[destination] = "noway"
	}

	destination := ""
	for k, v := range roadMap {
		if v == "noway" {
			return k
		}
	}

	return destination
}
