package even

import "testing"

/*
正常的用例
反面的用例（错误的输入，如用负数或字母代替数字，没有输入等）
边界检查用例（如果参数的取值范围是 0 到 1000，检查 0 和 1000 的情况）
*/

func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log("10 must be even")
		t.Fail()
	}
	if Even(7) {
		
		t.Log("7 is not even")
		t.Fail()
	}
}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log(" 11 must be odd!")
		t.Fail()
	}
	if Odd(10) {
		t.Log(" 10 is not odd!")
		t.Fail()
	}
}
