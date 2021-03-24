func romanToInt(s string) int {
	romanMaster := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	returnNum := 0
	skipFlg := false
	string := strings.Split(s, "")
	for i := 0; i < len(string); i++ {
		// 2つの数字を既に計算していたらSkip
		if skipFlg {
			skipFlg = false
			continue
		}

		var in int
		in = romanMaster[string[i]]
		if i < len(string)-1 {
			// sliceから抜き出す
			in = romanMaster[string[i]]
			str := string[i]
			str2 := string[i+1]

			// もっといい方法はないか？
			if str == "I" && (str2 == "V" || str2 == "X") {
				in = romanMaster[string[i+1]] - romanMaster[string[i]]
				skipFlg = true
			}

			if str == "X" && (str2 == "L" || str2 == "C") {
				in = romanMaster[string[i+1]] - romanMaster[string[i]]
				skipFlg = true
			}

			if str == "C" && (str2 == "D" || str2 == "M") {
				in = romanMaster[string[i+1]] - romanMaster[string[i]]
				skipFlg = true
			}

			returnNum += in
		} else {
			returnNum += romanMaster[string[i]]
		}
	}
	return returnNum
}