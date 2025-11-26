package main

import (
    "bytes"
    "encoding/json"
    "capital-gains/internal/usecase"
    "strings"
    "testing"
)

func TestCLI_Case2(t *testing.T) {
    input := `[{"operation":"buy","unit-cost":10.00,"quantity":10000},
{"operation":"sell","unit-cost":20.00,"quantity":5000},
{"operation":"sell","unit-cost":5.00,"quantity":5000}]
`

    var out bytes.Buffer
    run(strings.NewReader(input), &out)

    expectedJSON := []usecase.TaxResult{
        {Tax: 0},
        {Tax: 10000},
        {Tax: 0},
    }

    expected, _ := json.Marshal(expectedJSON)

    if !bytes.Contains(out.Bytes(), expected) {
        t.Errorf("\nExpected: %s\nGot: %s", expected, out.String())
    }
}

func TestCLI_Case1(t *testing.T) {
    input := `[{"operation":"buy","unit-cost":10.00,"quantity":100},
{"operation":"sell","unit-cost":15.00,"quantity":50},
{"operation":"sell","unit-cost":15.00,"quantity":50}]
`

    var out bytes.Buffer
    run(strings.NewReader(input), &out)

    expected := `[{"tax":0},{"tax":0},{"tax":0}]`

    if !bytes.Contains(out.Bytes(), []byte(expected)) {
        t.Errorf("Expected %s, got %s", expected, out.String())
    }
}
