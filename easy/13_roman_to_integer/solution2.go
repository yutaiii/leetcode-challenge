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
	string := strings.Split(s, "")
	for i := 0; i < len(string); i++ {

		var in int
		in = romanMaster[string[i]]
		// sliceの最後でない場合
		if i < len(string)-1 {
			// sliceから抜き出す
			str := string[i]
			str2 := string[i+1]

			// もっといい方法はないか？
			if str == "I" && (str2 == "V" || str2 == "X") {
				in = romanMaster[string[i+1]] - romanMaster[string[i]]
				// 1つスキップ
				i++
			}

			if str == "X" && (str2 == "L" || str2 == "C") {
				in = romanMaster[string[i+1]] - romanMaster[string[i]]
				// 1つスキップ
				i++
			}

			if str == "C" && (str2 == "D" || str2 == "M") {
				in = romanMaster[string[i+1]] - romanMaster[string[i]]
				// 1つスキップ
				i++
			}

			returnNum += in
		} else {
			// sliceの最後の場合
			returnNum += romanMaster[string[i]]
		}
	}
	return returnNum
}