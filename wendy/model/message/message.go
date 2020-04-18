package message

import (
	"strings"
)

type Message struct {
	sb *strings.Builder
}

func New(str string) *Message {
	sb := strings.Builder{}
	sb.WriteString(str)

	return &Message{
		sb: &sb,
	}
}

func (t *Message) GetString() string {
	return t.sb.String()
}

func (t *Message) AddLine(str string) *Message {
	t.sb.WriteString(str)
	return t
}

func (t *Message) AddNewLine(str string) *Message {
	t.sb.WriteString("\n")
	t.sb.WriteString(str)
	return t
}
