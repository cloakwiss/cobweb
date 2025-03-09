package ui

import (
	// "time"

	"time"

	"github.com/cloakwiss/cobweb/app"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	nothing struct{}
	quit    struct{}
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
			} else {
				return done{}
			}
		// case <-time.After(time.Millisecond * 4):
		default:
			return nothing{}
		}
	}
}

func NewModel(c chan app.ApMsg) UiState {
	return UiState{
		buf:           make([]string, 1000),
		pollChannel:   getMsg(c),
		spinnner:      spinner.New(),
		channelOpen:   true,
		receivedCount: 0,
		printedCount:  0,
	}
}

type UiState struct {
	buf           []string
	receivedCount uint
	printedCount  uint
	channelOpen   bool

	pollChannel tea.Cmd
	spinnner    spinner.Model
}

type printline struct{}

func (u UiState) Init() tea.Cmd {
	return nil
}

func printTick() tea.Cmd {
	return tea.Tick(time.Millisecond*8, func(_ time.Time) tea.Msg {
		return printline{}
	})
}

func relay(in tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return in
	}
}

func (u UiState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return u, tea.Quit
		}
	case spinner.TickMsg:
		u.spinnner, cmd = u.spinnner.Update(msg)
	case done:
		u.channelOpen = false
	case quit:
		return u, tea.Quit
	case app.ApMsg:
		u.buf[u.receivedCount] = msg.String()
		u.receivedCount += 1
	case printline:
		if u.printedCount < u.receivedCount {
			// if n := u.printedCount - 1; n >= 0 && u.receivedCount > n {
			line := u.buf[u.printedCount]
			cmd = tea.Printf("%s  %s", checkMark, line)
			// }
			u.printedCount += 1
		} else {
			if u.channelOpen {
				cmd = relay(nothing{})
			} else {
				cmd = relay(quit{})
			}
		}
	case nothing:
	}
	if u.channelOpen {
		return u, tea.Batch(u.pollChannel, u.spinnner.Tick, cmd, printTick())
	}
	return u, tea.Batch(u.spinnner.Tick, cmd, printTick())
}

var (
	checkMark = "🗸"
)

func (u UiState) View() (out string) {
	out = u.spinnner.View() + "    "
	// pr := strconv.FormatUint(uint64(u.printedCount), 10)
	// rec := strconv.FormatUint(uint64(u.receivedCount), 10)
	// out += pr + "/" + rec
	return
}
