package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Scanner struct {
	reader *bufio.Reader
	writer io.Writer
}

func NewScanner(input io.Reader, output io.Writer) *Scanner {
	return &Scanner{
		reader: bufio.NewReader(input),
		writer: output,
	}
}

func (s *Scanner) ReadFloat() (float64, error) {
	text, err := s.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	text = strings.TrimSpace(text)
	return strconv.ParseFloat(text, 64)
}

func (s *Scanner) ReadInt() (int, error) {
	text, err := s.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	text = strings.TrimSpace(text)
	return strconv.Atoi(text)
}

func (s *Scanner) Println(a ...any) {
	fmt.Fprintln(s.writer, a...)
}

func (s *Scanner) Print(a ...any) {
	fmt.Fprint(s.writer, a...)
}
