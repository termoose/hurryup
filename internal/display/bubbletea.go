package display

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"hurryup/internal/speed"
	"time"
)

type model struct {
	ping         speed.Ping
	downloadRate speed.DownloadRate
	uploadRate   speed.UploadRate

	pings         <-chan speed.Ping
	downloadRates <-chan speed.DownloadRate
	uploadRates   <-chan speed.UploadRate

	name    string
	country string

	pingSpinner     spinner.Model
	downloadSpinner spinner.Model
	uploadSpinner   spinner.Model
}

func BubbleTea(tester speed.Tester) error {
	pings, downloadRate, uploadRate := tester.RunAll()
	serverData := tester.GetServerData()

	p := tea.NewProgram(model{
		pings:           pings,
		downloadRates:   downloadRate,
		uploadRates:     uploadRate,
		name:            serverData.Name,
		country:         serverData.Country,
		pingSpinner:     spinner.New(spinner.WithSpinner(spinner.MiniDot)),
		downloadSpinner: spinner.New(spinner.WithSpinner(spinner.MiniDot)),
		uploadSpinner:   spinner.New(spinner.WithSpinner(spinner.MiniDot)),
	})

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("could not start program: %w", err)
	}

	return nil
}

func processMeasurement[T any](data <-chan T, cancel bool) tea.Cmd {
	return func() tea.Msg {
		r, ok := <-data
		if ok {
			return r
		} else {
			if cancel {
				return tea.Quit()
			}

			return nil
		}
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		processMeasurement(m.pings, false),
		processMeasurement(m.downloadRates, false),
		processMeasurement(m.uploadRates, true),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case speed.Ping:
		m.ping = msg.(speed.Ping)
		m.pingSpinner, _ = m.pingSpinner.Update(m.pingSpinner.Tick())
		return m, processMeasurement(m.pings, false)

	case speed.DownloadRate:
		m.downloadRate = msg.(speed.DownloadRate)
		m.downloadSpinner, _ = m.downloadSpinner.Update(m.downloadSpinner.Tick())
		return m, processMeasurement(m.downloadRates, false)

	case speed.UploadRate:
		m.uploadRate = msg.(speed.UploadRate)
		m.uploadSpinner, _ = m.uploadSpinner.Update(m.uploadSpinner.Tick())
		return m, processMeasurement(m.uploadRates, true)

	case tea.QuitMsg:
		return m, tea.Quit

	case tea.KeyMsg:
		key := msg.(tea.KeyMsg)
		return m, handleKeyPress(key)

	default:
		return m, nil
	}
}

func (m model) View() string {
	s := fmt.Sprintf("🌎 %s, %s\n"+
		"%s Latency: %s\n"+
		"%s Download: %.1f\n"+
		"%s Upload: %.1f\n",
		m.name, m.country,
		m.pingSpinner.View(), time.Duration(m.ping).Round(time.Millisecond),
		m.downloadSpinner.View(), float64(m.downloadRate),
		m.uploadSpinner.View(), float64(m.uploadRate))
	return s
}

func handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch msg.Type {
	case tea.KeyCtrlC:
		return tea.Quit
	default:
		return nil
	}
}
