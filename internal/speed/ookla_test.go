package speed

import "testing"

func TestOokla(t *testing.T) {
	ookla, err := NewOokla()

	if err != nil {
		t.Errorf("could not create ookla instance")
	}

	ping, down, up := ookla.RunAll()

	pings := 0
	for range ping {
		pings++
	}

	downs := 0
	for range down {
		downs++
	}

	ups := 0
	for range up {
		ups++
	}

	if pings == 0 {
		t.Errorf("no ping data received")
	}

	if downs == 0 {
		t.Errorf("no download data received")
	}

	if ups == 0 {
		t.Errorf("no upload data received")
	}
}
