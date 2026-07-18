package even

func Even(num int) bool {
	return num%2 == 0
}
func Odd(i int) bool { // Exported function
	return i%2 != 0
}
