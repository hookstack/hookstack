package convoy

import (
	"embed"
	"strings"
)

type HttpMethod string

type TaskName string

//go:embed VERSION
var f embed.FS

func (t TaskName) SetPrefix(prefix string) TaskName {
	var name strings.Builder
	delim := "-"

	name.WriteString(prefix)
	name.WriteString(delim)
	name.WriteString(string(t))

	return TaskName(name.String())
}

func ReadVersion() ([]byte, error) {
	data, err := f.ReadFile("VERSION")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetVersion() string {
	v := "0.1.0"

	f, err := ReadVersion()
	if err != nil {
		return v
	}

	v = strings.TrimSuffix(string(f), "\n")
	return v
}

const (
	EventProcessor      TaskName = "EventProcessor"
	DeadLetterProcessor TaskName = "DeadLetterProcessor"
)
