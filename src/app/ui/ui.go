package ui

import (
	"time"

	"github.com/cloakwiss/cobweb/app"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	nothing struct{}
	done    struct{}
)

func getMsg(input <-chan app.ApMsg) tea.Cmd {
	return func() tea.Msg {
		// this is blocking the thread
		// SO NEED TO DO SOMETHING
		select {
		case msg, ok := <-input:
			if ok {
				return msg
			}
		case <-time.After(time.Millisecond * 16):
			return nothing{}
		}
		return done{}
	}
}

func NewModel(c chan app.ApMsg) UiState {
	return UiState{
		buf:         [2]string{},
		pollingFunc: getMsg(c),
		count:       0,
		spinnner:    spinner.New(),
	}
}

type UiState struct {
	buf   [2]string
	count uint

	pollingFunc tea.Cmd
	spinnner    spinner.Model
}

func (u UiState) Init() tea.Cmd {
	return nil
}

func (u UiState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case app.ApMsg:
		u.buf[u.count%2] = msg.String()
		cmd = tea.Printf("%s  %s", checkMark, u.buf[u.count%2])
		u.count += 1
	case done:
		return u, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return u, tea.Quit
		}
	case spinner.TickMsg:
		u.spinnner, cmd = u.spinnner.Update(msg)
	case nothing:
	}
	return u, tea.Batch(u.pollingFunc, u.spinnner.Tick, cmd)
}

var (
	checkMark = "🗸"
)

func (u UiState) View() (out string) {
	out = u.spinnner.View()
	out += "  "
	out += u.buf[u.count%2]
	return
}
