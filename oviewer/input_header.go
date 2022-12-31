package oviewer

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

// setHeaderMode sets the inputMode to Header.
func (root *Root) setHeaderMode() {
	input := root.input
	input.value = ""
	input.cursorX = 0
	input.Event = newHeaderEvent()
}

// eventHeader represents the goto input mode.
type eventHeader struct {
	value string
	tcell.EventTime
}

// newHeaderEvent returns headerEvent.
func newHeaderEvent() *eventHeader {
	return &eventHeader{}
}

// Mode returns InputMode.
func (e *eventHeader) Mode() InputMode {
	return Header
}

// Prompt returns the prompt string in the input field.
func (e *eventHeader) Prompt() string {
	return "Header length:"
}

// Confirm returns the event when the input is confirmed.
func (e *eventHeader) Confirm(str string) tcell.Event {
	e.value = str
	e.SetEventNow()
	return e
}

// Up returns strings when the up key is pressed during input.
func (e *eventHeader) Up(str string) string {
	n, err := strconv.Atoi(str)
	if err != nil {
		return "0"
	}
	return strconv.Itoa(n + 1)
}

// Down returns strings when the down key is pressed during input.
func (e *eventHeader) Down(str string) string {
	n, err := strconv.Atoi(str)
	if err != nil || n <= 0 {
		return "0"
	}
	return strconv.Itoa(n - 1)
}
