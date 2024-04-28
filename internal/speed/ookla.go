package speed

import (
	"errors"
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
	"hurryup/internal/tools"
	"time"
)

type Ookla struct {
	client  *speedtest.Speedtest
	closest *speedtest.Server

	runAllFinished bool
}

func NewOokla() (*Ookla, error) {
	client := speedtest.New()
	servers, err := client.FetchServers()

	if err != nil {
		return nil, fmt.Errorf("could not fetch servers %w", err)
	}

	if servers.Len() == 0 {
		return nil, errors.New("could not find any servers")
	}

	return &Ookla{
		client:         client,
		closest:        servers[0],
		runAllFinished: false,
	}, nil
}

func (s *Ookla) RunAll() (<-chan Ping, <-chan DownloadRate, <-chan UploadRate) {
	pings := tools.NewChan[Ping]()
	downloadRate := tools.NewChan[DownloadRate]()
	uploadRate := tools.NewChan[UploadRate]()

	defer func() {
		s.runAllFinished = true
	}()

	go func() {
		_ = s.closest.PingTest(func(latency time.Duration) {
			pings.Send(Ping(latency))
		})
		pings.Close()

		td := s.closest.Context.CallbackDownloadRate(func(rate float64) {
			downloadRate.Send(DownloadRate(rate))
		})
		_ = s.closest.DownloadTest()
		td.Stop()
		downloadRate.Close()

		tu := s.closest.Context.CallbackUploadRate(func(rate float64) {
			uploadRate.Send(UploadRate(rate))
		})
		_ = s.closest.UploadTest()
		tu.Stop()
		uploadRate.Close()
	}()

	return pings.Data, downloadRate.Data, uploadRate.Data
}

func (s *Ookla) GetTesterData() TesterData {
	s.client.Manager.Reset()

	if !s.runAllFinished {
		_ = s.closest.PingTest(nil)
		_ = s.closest.DownloadTest()
		_ = s.closest.UploadTest()
	}

	return TesterData{
		Latency:      Ping(s.closest.Latency),
		DownloadRate: DownloadRate(s.closest.DLSpeed),
		UploadRate:   UploadRate(s.closest.ULSpeed),
	}
}

func (s *Ookla) GetServerData() ServerData {
	return ServerData{
		Name:    s.closest.Name,
		Country: s.closest.Country,
	}
}
