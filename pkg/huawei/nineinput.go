package huawei

func NineInput(s string) []string {
	stringSet := []string{""}
	mappings := []string{
		"", "", "abc", "def",
		"ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz"}
	for _, c := range s {
		add := mappings[int(c-'0')]
		tmpSet := []string{}
		for _, a := range add {
			for _, o := range stringSet {
				tmpSet = append(tmpSet, o+string(a))
			}
		}
		stringSet = tmpSet
	}
	return stringSet
}
