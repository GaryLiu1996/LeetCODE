package main

func main() {
	isValid("((")
}

func isValid(s string) bool {
	var array []string
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "(", "[", "{":
			array = append(array, string(s[i]))
		case ")":
			if len(array) == 0 {
				return false
			} else if array[len(array)-1] == "(" {
				array = array[0:(len(array) - 1)]
			} else {
				return false
			}
		case "]":
			if len(array) == 0 {
				return false
			} else if array[len(array)-1] == "[" {
				array = array[0:(len(array) - 1)]
			} else {
				return false
			}
		case "}":
			if len(array) == 0 {
				return false
			} else if array[len(array)-1] == "{" {
				array = array[0:(len(array) - 1)]
			} else {
				return false
			}
		default:
			return false
		}
	}
	if len(array) != 0 {
		return false
	}
	return true
}
