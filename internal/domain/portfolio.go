package domain

type Portfolio struct {
    Quantity        int
    AvgPrice        float64
    AccumulatedLoss float64
}

func NewPortfolio() *Portfolio {
    return &Portfolio{
        Quantity:        0,
        AvgPrice:        0.0,
        AccumulatedLoss: 0.0,
    }
}

func (p *Portfolio) ApplyBuy(unit float64, qty int) {
    if p.Quantity == 0 {
        p.AvgPrice = unit
        p.Quantity = qty
        return
    }
    totalCost := p.AvgPrice*float64(p.Quantity) + unit*float64(qty)
    p.Quantity += qty
    p.AvgPrice = totalCost / float64(p.Quantity)
}

func (p *Portfolio) ApplySell(unit float64, qty int) (profit float64, totalValue float64) {
    totalValue = float64(qty) * unit
    profit = float64(qty) * (unit - p.AvgPrice)
    p.Quantity -= qty
    return
}
