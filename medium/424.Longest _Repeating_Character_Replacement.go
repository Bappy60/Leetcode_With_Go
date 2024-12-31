package main
func characterReplacement(s string, k int) int {
	start, end := 0, 0
	sChars := []rune(s)
	freqTable := make([]int, 26)
	maxFreqCnt, ans := 0, 0
	for end < len(s) {
		curChar := sChars[end]
		freqTable[curChar-'A']++
		if freqTable[curChar-'A'] > maxFreqCnt {
			maxFreqCnt = freqTable[curChar-'A']
		}
		end += 1
		if !isWindowValid(start, end, maxFreqCnt, k) {
			frontChar := sChars[start]
			freqTable[frontChar-'A']--
			start += 1
		}
		if end-start > 0 {
			ans = end - start
		}
	}
	return ans
}

func isWindowValid(start, end, maxFreq, k int) bool {
	return end-start-maxFreq <= k
}