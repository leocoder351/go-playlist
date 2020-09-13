package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// Go 语言不支持自增运算
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestStromg(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))

}
