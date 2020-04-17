package decorator

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestNewLowCaseReader(t *testing.T) {
	text := "TEST! test! TesT! OnLY tESt! IT's also TeSt? YEap!"
	expected := strings.ToLower(text)

	textBuffer := bytes.NewBuffer([]byte(text))
	var reader io.Reader = bufio.NewReader(textBuffer)
	reader = NewLowCaseReader(reader)

	result := make([]byte, 0)
	buffer := make([]byte, 8)

	for {
		l, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		result = append(result, buffer[:l]...)
	}

	if string(result) != expected {
		t.Errorf("Получен неверный результат. Ожидалось %q, получено %q.", expected, result)
	}
}
