package lc

func findSubStringInWarproundString(p string) int {
	dp := [26]int{}
	k := 1
	dp[p[0]-'a'] = 1
	pLength := len(p)
	for i := 1; i < pLength; i++ {

		cur := p[i]
		pre := p[i-1]
		isSerial := (cur-pre+26)%26 == 1 //1 or -25

		if i > 0 && isSerial {
			k++
		} else {
			k = 1
		}

		if k > dp[cur-'a'] {
			dp[cur-'a'] = k
		}
	}

	result := 0
	for i := 0; i < 26; i++ {
		result += dp[i]
	}

	return result
}
