package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	sessionType    SessionType
	timeRemaining  int
	isRunning      bool
	completedFocus int // Track completed focus sessions for cycle management
	width          int
	height         int
}

// Timer tick message
type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func initialModel() model {
	return model{
		sessionType:    Focus,
		timeRemaining:  FocusDuration,
		isRunning:      true,
		completedFocus: 0,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("TermiPom"),
		tickCmd(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Resize window event
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	// Key press event
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			// Toggle pause/resume
			m.isRunning = !m.isRunning
			return m, nil
		case "s":
			// Skip current session and move to next
			return m.nextSession(), nil
		case "r":
			// Reset current session
			return m.resetCurrentSession(), nil
		}

	// Timer tick event
	// 1 frame per second
	case tickMsg:
		if m.isRunning && m.timeRemaining > 0 {
			m.timeRemaining--
			if m.timeRemaining == 0 {
				// Session completed, move to next session and start new tick
				return m.nextSession(), tickCmd()
			}
		}
		// Move to next tick (keeps ticking even when paused)
		return m, tickCmd()
	}

	return m, nil
}

// Move to the next session in the pomodoro cycle
func (m model) nextSession() model {
	switch m.sessionType {
	case Focus:
		m.completedFocus++
		if m.completedFocus%4 == 0 {
			// Time for long break after 4 focus sessions
			m.sessionType = LongBreak
			m.timeRemaining = LongBreakDuration
		} else {
			// Time for short break
			m.sessionType = ShortBreak
			m.timeRemaining = ShortBreakDuration
		}
	case ShortBreak:
		fallthrough
	case LongBreak:
		// Back to focus session
		m.sessionType = Focus
		m.timeRemaining = FocusDuration
	}
	m.isRunning = true
	return m
}

// Reset the current session to its full duration
func (m model) resetCurrentSession() model {
	switch m.sessionType {
	case Focus:
		m.timeRemaining = FocusDuration
	case ShortBreak:
		m.timeRemaining = ShortBreakDuration
	case LongBreak:
		m.timeRemaining = LongBreakDuration
	}
	return m
}

func (m model) View() string {
	// Define styles based on session type
	var bgStyle lipgloss.Style
	var sessionName string

	switch m.sessionType {
	case Focus:
		bgStyle = lipgloss.NewStyle().Background(lipgloss.Color(FocusBgColor))
		sessionName = "FOCUS"
	case ShortBreak:
		bgStyle = lipgloss.NewStyle().Background(lipgloss.Color(ShortBreakBgColor))
		sessionName = "SHORT BREAK"
	case LongBreak:
		bgStyle = lipgloss.NewStyle().Background(lipgloss.Color(LongBreakBgColor))
		sessionName = "LONG BREAK"
	}

	// Format time as MM:SS
	minutes := m.timeRemaining / 60
	seconds := m.timeRemaining % 60
	timeStr := fmt.Sprintf("%02d:%02d", minutes, seconds)

	// Create ASCII timer display
	timerDisplay := m.buildASCIITimer(timeStr)

	// Create content
	content := strings.Builder{}

	// Session type header
	sessionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Bold(true).
		MarginBottom(2)
	content.WriteString(sessionStyle.Render(sessionName))
	content.WriteString("\n\n")

	// Timer display
	timerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
	content.WriteString(timerStyle.Render(timerDisplay))
	content.WriteString("\n\n")

	// Status indicator
	status := "RUNNING"
	if !m.isRunning {
		status = "PAUSED"
	}
	statusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffff00")).
		Bold(true).
		MarginTop(1)
	content.WriteString(statusStyle.Render(status))

	// Help footer
	helpText := "[space] pause/resume  [s] skip  [r] reset  [q] quit"
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		MarginTop(3)

	// Center all content
	centeredContent := lipgloss.Place(
		m.width, m.height-2, // content size
		lipgloss.Center, lipgloss.Center,
		content.String(),
	)

	// Position help at bottom
	helpCentered := lipgloss.Place(
		m.width, 1, // help text size
		lipgloss.Center, lipgloss.Center,
		helpStyle.Render(helpText),
	)

	// Combine content and help
	fullContent := centeredContent + "\n" + helpCentered

	// Apply background to entire screen
	return bgStyle.Width(m.width).Height(m.height).Padding(1, 1).Render(fullContent)
}

func (m model) buildASCIITimer(timeStr string) string {
	lines := make([]string, 7) // ASCII numbers are 7 lines tall

	for i, char := range timeStr {
		charStr := string(char)
		if asciiArt, exists := asciiNumbers[charStr]; exists {
			for lineIdx, artLine := range asciiArt {
				if i > 0 {
					lines[lineIdx] += " " // Add space between characters
				}
				lines[lineIdx] += artLine
			}
		}
	}

	return strings.Join(lines, "\n")
}
