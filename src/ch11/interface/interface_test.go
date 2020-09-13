package interface_test

import "testing"

type Programmer interface {
	WhiteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WhiteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WhiteHelloWorld())
}
