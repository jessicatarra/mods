package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type tamode struct {
	textarea textarea.Model
	err      error
	Config   *Config
}

func newTaMode(cfg *Config) tamode {
	ti := textarea.New()
	ti.Placeholder = "Ask your AI assistant..."
	ti.Focus()

	return tamode{
		textarea: ti,
		err:      nil,
		Config:   cfg,
	}
}

func (t tamode) Init() tea.Cmd {
	//return textarea.Blink
	return textarea.Blink
}

func (t tamode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if t.textarea.Focused() {
				t.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return t, tea.Quit
		case tea.KeyCtrlS:
			t.Config.TextAreaInput = t.textarea.Value()
			return t, tea.Quit
		default:
			if !t.textarea.Focused() {
				cmd = t.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	case errMsg:
		t.err = msg
		return t, nil
	}

	t.textarea, cmd = t.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return t, tea.Batch(cmds...)
}

func (t tamode) View() string {
	return fmt.Sprintf(
		"Enter your prompt...\n\n%s\n\n%s",
		t.textarea.View(),
		"(ctrl+c to quit) (ctrl+s to save)",
	) + "\n\n"
}
