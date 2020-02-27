package leetcode

func DefangIpaddr(address string) string{
	return defangIPaddr(address)
}

func defangIPaddr(address string) string {
	var ret []rune
	for _, letter := range address{
		if letter == '.'{
			ret = append(ret, '[')
			ret = append(ret, letter)
			ret = append(ret, ']')
		}else{
			ret = append(ret, letter)
		}
		
	}
    return string(ret[:])
}