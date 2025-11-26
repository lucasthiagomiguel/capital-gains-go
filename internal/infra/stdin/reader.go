package stdin

import (
	"bufio"
	"io"
	"strings"
)

func ReadLinesWithReader(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	var lines []string
	var buffer strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// ignore começo vazio
		if line == "" && buffer.Len() == 0 {
			continue
		}

		buffer.WriteString(line)

		// quando fecha o array, finaliza uma entrada
		if strings.HasSuffix(line, "]") {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}

	// se sobrou algo sem fechar (raro)
	if buffer.Len() > 0 {
		lines = append(lines, buffer.String())
	}

	return lines
}
