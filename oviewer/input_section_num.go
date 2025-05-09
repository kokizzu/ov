package oviewer

import (
	"context"

	"github.com/gdamore/tcell/v2"
)

// inputSectionNum sets the inputMode to SectionNum.
func (root *Root) inputSectionNum(context.Context) {
	input := root.input
	input.reset()
	input.Event = newSectionNumEvent()
}

// eventSectionNum represents the section num input mode.
type eventSectionNum struct {
	tcell.EventTime
	value string
}

// newSectionNumEvent returns Event.
func newSectionNumEvent() *eventSectionNum {
	return &eventSectionNum{}
}

// Mode returns InputMode.
func (*eventSectionNum) Mode() InputMode {
	return SectionNum
}

// Prompt returns the prompt string in the input field.
func (*eventSectionNum) Prompt() string {
	return "Section Num:"
}

// Confirm returns the event when the input is confirmed.
func (e *eventSectionNum) Confirm(str string) tcell.Event {
	e.value = str
	e.SetEventNow()
	return e
}

// Up returns strings when the up key is pressed during input.
func (*eventSectionNum) Up(str string) string {
	return upNum(str)
}

// Down returns strings when the down key is pressed during input.
func (*eventSectionNum) Down(str string) string {
	return downNum(str)
}
