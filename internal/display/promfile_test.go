package display

import "hurryup/internal/speed"

func ExamplePromFile() {
	s := speed.MockTester{}
	_ = PromFile(s)

	// Output:
	// hurryup_download_speed 20.0
	// hurryup_upload_speed 30.0
	// hurryup_latency 10
	// hurryup_location Idaho
	// hurryup_country Bumsville
}
