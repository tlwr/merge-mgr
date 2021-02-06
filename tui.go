package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type teamodel struct {
	cursor   int
	choice   chan *GHOpenPR
	choices  map[int]GHOpenPR
	selected map[int]bool
}

func NewTUI(prs []GHOpenPR) (chan *GHOpenPR, *tea.Program) {
	choices := map[int]GHOpenPR{}
	for i, pr := range prs {
		choices[i] = pr
	}

	result := make(chan *GHOpenPR, 64)
	tui := tea.NewProgram(teamodel{
		cursor:   0,
		choice:   result,
		choices:  choices,
		selected: map[int]bool{},
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
			for i := 0; i < len(m.choices); i++ {
				if m.selected[i] {
					choice := m.choices[i]
					m.choice <- &choice
				}
			}
			close(m.choice)
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

		case " ":
			m.selected[m.cursor] = !m.selected[m.cursor]
		}

	}

	return m, nil
}

func (m teamodel) View() string {
	s := strings.Builder{}

	s.WriteString("choose a PR to merge\n\n")

	for i := 0; i < len(m.choices); i++ {
		var prefix = "  "
		if m.cursor == i {
			prefix = "> "
		}

		var symbol = " "
		if m.selected[i] {
			symbol = "•"
		}

		s.WriteString(fmt.Sprintf("%s(%s) %s \n", prefix, symbol, m.choices[i].Display()))
	}

	s.WriteString("\n(press space to select ; enter to proceed ; q to quit)\n")

	return s.String()
}
