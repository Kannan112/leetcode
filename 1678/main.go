package main

func main() {
	str := "G()()()()(al)"
	interpret(str)

}

//	func interpret(command string) string {
//		result := ""
//		for index, Asc := range command {
//			if Asc == 40 {
//				if command[index+1] == 41 {
//					result += "O"
//				}else if
//			}\
//		}
//		return "sdf"
//	}
func interpret(command string) string {
	output := ""
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			output = output + "G"
		} else if command[i] == '(' && command[i+1] == ')' {
			output = output + "o"
			i++
		} else {
			output = output + "al"
		}
	}
	return output
}
