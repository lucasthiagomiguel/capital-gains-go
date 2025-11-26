package usecase

import (
    "capital-gains/internal/domain"
    "math"
)

type TaxResult struct {
    Tax float64 `json:"tax"`
}

func round2(v float64) float64 {
    return math.Round(v*100) / 100
}

func CalculateTax(ops []domain.Operation) []TaxResult {
    p := domain.NewPortfolio()
    results := make([]TaxResult, 0, len(ops))

    for _, op := range ops {
        if op.Operation == "buy" {
            p.ApplyBuy(op.UnitCost, op.Quantity)
            results = append(results, TaxResult{Tax: 0.0})
            continue
        }

        // SELL
        profit, totalValue := p.ApplySell(op.UnitCost, op.Quantity)

        // regra do limite <= 20.000
        if totalValue <= 20000 {
            if profit < 0 {
                p.AccumulatedLoss += -profit
            }
            results = append(results, TaxResult{Tax: 0.0})
            continue
        }

        if profit > 0 {
            if p.AccumulatedLoss > 0 {
                if p.AccumulatedLoss >= profit {
                    p.AccumulatedLoss -= profit
                    results = append(results, TaxResult{Tax: 0.0})
                    continue
                } else {
                    profit -= p.AccumulatedLoss
                    p.AccumulatedLoss = 0
                }
            }
            tax := round2(profit * 0.20)
            results = append(results, TaxResult{Tax: tax})
        } else {
            p.AccumulatedLoss += -profit
            results = append(results, TaxResult{Tax: 0.0})
        }
    }
    return results
}
