package tools

import "testing"

func TestChan(t *testing.T) {
	inElem := 1337
	var outElem int
	c := NewChan[int]()

	go func() {
		c.Send(inElem)
		c.Close()
	}()

	outElem = <-c.Data

	if outElem != inElem {
		t.Errorf("%d not equal %d", inElem, outElem)
	}
}
