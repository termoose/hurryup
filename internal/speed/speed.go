package speed

import "time"

type DownloadRate float64
type UploadRate float64
type Ping time.Duration

type Tester interface {
	RunAll() (<-chan Ping, <-chan DownloadRate, <-chan UploadRate)
	GetTesterData() TesterData
	GetServerData() ServerData
}

type ServerData struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type TesterData struct {
	Latency      Ping         `json:"latency"`
	DownloadRate DownloadRate `json:"download_rate"`
	UploadRate   UploadRate   `json:"upload_rate"`
}

type Data struct {
	ServerData
	TesterData
}

type MockTester struct{}

func (m MockTester) RunAll() (<-chan Ping, <-chan DownloadRate, <-chan UploadRate) {
	return nil, nil, nil
}

func (m MockTester) GetTesterData() TesterData {
	return TesterData{
		Latency:      Ping(10),
		DownloadRate: DownloadRate(20),
		UploadRate:   UploadRate(30),
	}
}

func (m MockTester) GetServerData() ServerData {
	return ServerData{
		Name:    "Idaho",
		Country: "Bumsville",
	}
}
