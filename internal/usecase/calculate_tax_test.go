package usecase

import (
    "capital-gains/internal/domain"
    "encoding/json"
    "testing"
)

func asOps(jsonStr string) []domain.Operation {
    var ops []domain.Operation
    json.Unmarshal([]byte(jsonStr), &ops)
    return ops
}

func asFloatSlice(results []TaxResult) []float64 {
    out := []float64{}
    for _, r := range results {
        out = append(out, r.Tax)
    }
    return out
}

func equalSlices(a, b []float64) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TestCase1(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy",  "unit-cost":10.00, "quantity":100},
        {"operation":"sell", "unit-cost":15.00, "quantity":50},
        {"operation":"sell", "unit-cost":15.00, "quantity":50}
    ]
    `)
    res := CalculateTax(ops)
    expected := []float64{0, 0, 0}
    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 1 expected %v, got %v", expected, asFloatSlice(res))
    }
}

func TestCase2(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy", "unit-cost":10.00, "quantity":10000},
        {"operation":"sell","unit-cost":20.00, "quantity":5000},
        {"operation":"sell","unit-cost":5.00,  "quantity":5000}
    ]
    `)

    res := CalculateTax(ops)
    expected := []float64{0, 10000, 0}

    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 2 expected %v, got %v", expected, asFloatSlice(res))
    }
}

func TestCase3(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy", "unit-cost":10.00, "quantity":10000},
        {"operation":"sell","unit-cost":5.00,  "quantity":5000},
        {"operation":"sell","unit-cost":20.00, "quantity":3000}
    ]
    `)

    res := CalculateTax(ops)
    expected := []float64{0, 0, 1000}

    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 3 expected %v, got %v", expected, asFloatSlice(res))
    }
}

func TestCase4(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy", "unit-cost":10.00, "quantity":10000},
        {"operation":"buy", "unit-cost":25.00, "quantity":5000},
        {"operation":"sell","unit-cost":15.00, "quantity":10000}
    ]
    `)

    res := CalculateTax(ops)
    expected := []float64{0, 0, 0}

    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 4 expected %v, got %v", expected, asFloatSlice(res))
    }
}

func TestCase5(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy",  "unit-cost":10.00, "quantity":10000},
        {"operation":"buy",  "unit-cost":25.00, "quantity":5000},
        {"operation":"sell", "unit-cost":15.00, "quantity":10000},
        {"operation":"sell", "unit-cost":25.00, "quantity":5000}
    ]
    `)

    res := CalculateTax(ops)
    expected := []float64{0, 0, 0, 10000}

    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 5 expected %v, got %v", expected, asFloatSlice(res))
    }
}

func TestCase6(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy",  "unit-cost":10.00, "quantity":10000},
        {"operation":"sell", "unit-cost":2.00,  "quantity":5000},
        {"operation":"sell", "unit-cost":20.00, "quantity":2000},
        {"operation":"sell", "unit-cost":20.00, "quantity":2000},
        {"operation":"sell", "unit-cost":25.00, "quantity":1000}
    ]
    `)

    res := CalculateTax(ops)
    expected := []float64{0, 0, 0, 0, 3000}

    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 6 expected %v, got %v", expected, asFloatSlice(res))
    }
}

func TestCase7(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy",  "unit-cost":10.00, "quantity":10000},
        {"operation":"sell", "unit-cost":2.00,  "quantity":5000},
        {"operation":"sell", "unit-cost":20.00, "quantity":2000},
        {"operation":"sell", "unit-cost":20.00, "quantity":2000},
        {"operation":"sell", "unit-cost":25.00, "quantity":1000},
        {"operation":"buy",  "unit-cost":20.00, "quantity":10000},
        {"operation":"sell", "unit-cost":15.00, "quantity":5000},
        {"operation":"sell", "unit-cost":30.00, "quantity":4350},
        {"operation":"sell", "unit-cost":30.00, "quantity":650}
    ]
    `)

    res := CalculateTax(ops)
    expected := []float64{0, 0, 0, 0, 3000, 0, 0, 3700, 0}

    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 7 expected %v, got %v", expected, asFloatSlice(res))
    }
}

func TestCase8(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy",  "unit-cost":10.00, "quantity":10000},
        {"operation":"sell", "unit-cost":50.00, "quantity":10000},
        {"operation":"buy",  "unit-cost":20.00, "quantity":10000},
        {"operation":"sell", "unit-cost":50.00, "quantity":10000}
    ]
    `)

    res := CalculateTax(ops)
    expected := []float64{0, 80000, 0, 60000}

    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 8 expected %v, got %v", expected, asFloatSlice(res))
    }
}

func TestCase9(t *testing.T) {
    ops := asOps(`
    [
        {"operation":"buy",  "unit-cost":5000.00, "quantity":10},
        {"operation":"sell", "unit-cost":4000.00, "quantity":5},
        {"operation":"buy",  "unit-cost":15000.00, "quantity":5},
        {"operation":"buy",  "unit-cost":4000.00,  "quantity":2},
        {"operation":"buy",  "unit-cost":23000.00, "quantity":2},
        {"operation":"sell", "unit-cost":20000.00, "quantity":1},
        {"operation":"sell", "unit-cost":12000.00, "quantity":10},
        {"operation":"sell", "unit-cost":15000.00, "quantity":3}
    ]
    `)

    res := CalculateTax(ops)
    expected := []float64{0, 0, 0, 0, 0, 0, 1000, 2400}

    if !equalSlices(asFloatSlice(res), expected) {
        t.Errorf("Case 9 expected %v, got %v", expected, asFloatSlice(res))
    }
}
