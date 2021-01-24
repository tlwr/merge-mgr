package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type teamodel struct {
	cursor  int
	choice  chan *GHOpenPR
	choices map[int]GHOpenPR
}

func NewTUI(prs []GHOpenPR) (chan *GHOpenPR, *tea.Program) {
	choices := map[int]GHOpenPR{}
	for i, pr := range prs {
		choices[i] = pr
	}

	result := make(chan *GHOpenPR, 1)
	tui := tea.NewProgram(teamodel{
		cursor:  0,
		choice:  result,
		choices: choices,
	})

	return result, tui
}

func (m teamodel) Init() tea.Cmd {
	return nil
}

func (m teamodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			close(m.choice)
			return m, tea.Quit

		case "enter":
			choice := m.choices[m.cursor]
			m.choice <- &choice
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}

	}

	return m, nil
}

func (m teamodel) View() string {
	s := strings.Builder{}
	s.WriteString("choose a PR to merge\n\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		s.WriteString(m.choices[i].Display())
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}
