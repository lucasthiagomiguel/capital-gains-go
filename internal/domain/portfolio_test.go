package domain

import "testing"

func TestApplyBuy(t *testing.T) {
    p := NewPortfolio()

    p.ApplyBuy(10.0, 100)
    if p.Quantity != 100 {
        t.Errorf("expected 100, got %d", p.Quantity)
    }

    if p.AvgPrice != 10.0 {
        t.Errorf("expected avg 10.0, got %.2f", p.AvgPrice)
    }

    p.ApplyBuy(20.0, 100)
    expectedAvg := ((10*100) + (20*100)) / 200.0

    if p.AvgPrice != expectedAvg {
        t.Errorf("avg expected %.2f, got %.2f", expectedAvg, p.AvgPrice)
    }
}

func TestApplySell(t *testing.T) {
    p := NewPortfolio()
    p.ApplyBuy(10, 100)

    profit, total := p.ApplySell(15, 50)

    if p.Quantity != 50 {
        t.Errorf("expected 50 left, got %d", p.Quantity)
    }

    if total != 15*50 {
        t.Errorf("total expected %d, got %.2f", 15*50, total)
    }

    expectedProfit := float64(50) * (15 - 10)
    if profit != expectedProfit {
        t.Errorf("profit expected %.2f, got %.2f", expectedProfit, profit)
    }
}
