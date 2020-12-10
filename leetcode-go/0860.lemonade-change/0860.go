package _860_lemonade_change

//Mysoluion
//func lemonadeChange(bills []int) bool {
//	billsMap := make(map[int]int, 3)
//	for i := 0; i < len(bills); i++ {
//		if bills[i] == 5 {
//			billsMap[5]++
//		}
//		if bills[i] == 10 {
//			if billsMap[5] > 0 {
//				billsMap[5]--
//				billsMap[10]++
//			} else {
//				return false
//			}
//		}
//		if bills[i] == 20 {
//			if billsMap[5] > 0 && billsMap[10] > 0 {
//				billsMap[5]--
//				billsMap[10]--
//				billsMap[20]++
//			} else if billsMap[5] > 2 {
//				billsMap[5] -= 3
//				billsMap[20]++
//			} else {
//				return false
//			}
//		}
//	}
//	return true
//}
//Mysoluion改进版本
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

//官方版本
//func lemonadeChange(bills []int) bool {
//    five, ten := 0, 0
//    for _, bill := range bills {
//        if bill == 5 {
//            five++
//        } else if bill == 10 {
//            if five == 0 {
//                return false
//            }
//            five--
//            ten++
//        } else {
//            if five > 0 && ten > 0 {
//                five--
//                ten--
//            } else if five >= 3 {
//                five -= 3
//            } else {
//                return false
//            }
//        }
//    }
//    return true
//}
