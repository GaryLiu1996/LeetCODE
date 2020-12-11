package _649_dota2_senate

//func predictPartyVictory(senate string) string {
//	arr := []byte{}
//	r := 0
//	d := 0
//	for i:=0; i < len(senate); i++ {
//		if r>=0{
//			arr = append(arr, senate[i])
//			d--
//		}
//		if d>=0{
//			arr = append(arr, senate[i])
//
//		}
//	}
//	return ""
//}

//guanfang
func predictPartyVictory(senate string) string {
	var radiant, dire []int
	for i, s := range senate {
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
