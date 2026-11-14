func groupAnagrams(strs []string) [][]string {
    stringMap := make(map[string][]string)
    for _,v := range(strs){
        sortedString := sortString(v)
        _,ok := stringMap[sortedString]
        if ok{
            stringMap[sortedString] = append(stringMap[sortedString], v)
        }else{
            stringMap[sortedString] = []string{v}
        }
    }

    result := make([][]string,0)
    for _,group := range stringMap{
        result = append(result, group)
    }
    return result

}

func sortString(s string) string {
	chars := []byte(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}


