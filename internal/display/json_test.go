package display

import "hurryup/internal/speed"

func ExampleJSON() {
	s := speed.MockTester{}
	_ = JSON(s)
	// Output:
	// {"name":"Idaho","country":"Bumsville","latency":10,"download_rate":20,"upload_rate":30}
}
