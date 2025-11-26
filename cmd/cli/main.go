package main

import (
    "capital-gains/internal/domain"
    "capital-gains/internal/infra/stdin"
    "capital-gains/internal/usecase"
    "encoding/json"
    "fmt"
    "io"
    "os"
)

func run(r io.Reader, w io.Writer) {
    lines := stdin.ReadLinesWithReader(r)

    for _, line := range lines {
        var ops []domain.Operation

        if err := json.Unmarshal([]byte(line), &ops); err != nil {
            continue
        }

        results := usecase.CalculateTax(ops)

        out, _ := json.Marshal(results)
        fmt.Fprintln(w, string(out))
    }
}

func main() {
    run(os.Stdin, os.Stdout)
}
