package main

func groupAnagrams(words []string) [][]string {
	m := make(map[[26]int][]string)

	res := [][]string{}
	for i := 0; i < len(words); i++ {
		alp := [26]int{}
		for k := 0; k < len(words[i]); k++ {
			alp[words[i][k]-'a']++
		}
		m[alp] = append(m[alp], words[i])
	}

	for _, v := range m {
		res = append(res, v)
	}
	return res
}
