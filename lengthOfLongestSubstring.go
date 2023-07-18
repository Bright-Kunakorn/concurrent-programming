func lengthOfLongestSubstring(s string) int {
    window := make(map[byte]bool)
    left := 0
    right := 0
    maxLength := 0

    for right < len(s) {
        if !window[s[right]] {
            window[s[right]] = true
            if maxLength < right-left+1 {
                maxLength = right - left + 1
            }
            right++
        } else {
            delete(window, s[left])
            left++
        }
    }

    return maxLength
}
