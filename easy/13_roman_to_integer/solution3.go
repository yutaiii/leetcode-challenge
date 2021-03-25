func romanToInt(s string) int {
	lastLetter := ""
	returnNum := 0
	romanMaster := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	takeAway := 2
	for _, val := range s {
		letter := string(val)
		if letter == "V" || letter == "X" {
			if lastLetter == "I" {
				returnNum += (romanMaster[letter] - takeAway)
				continue
			}
		}

		if letter == "L" || letter == "C" {
			if lastLetter == "X" {
				returnNum += (romanMaster[letter] - takeAway*10)
				continue
			}
		}

		if letter == "D" || letter == "M" {
			if lastLetter == "C" {
				returnNum += (romanMaster[letter] - takeAway*100)
				continue
			}
		}

		returnNum += romanMaster[letter]
		lastLetter = letter
	}
	return returnNum
}