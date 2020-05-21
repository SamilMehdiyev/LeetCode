type StockSpanner struct {
    prices []int
    spans []int
    days int
}


func Constructor() StockSpanner {
    return StockSpanner{[]int{}, []int{}, 0}
}


func (this *StockSpanner) Next(price int) int {
    spans := 1
    
    if this.days > 1 && price < this.prices[len(this.prices) - 1] {
        fmt.Println(this.prices[len(this.prices) - 1])
        this.prices = append(this.prices, price)
        this.spans = append(this.spans, spans)
        return spans
    }
    
    lastSpans := 1
    if len(this.spans) > 0 {
        lastSpans = this.spans[len(this.spans) - 1]
        spans = lastSpans
    }
    for i := len(this.prices) - lastSpans; i >= 0; i-- {
        if this.prices[i] <= price {
            spans++
        } else {
            break
        }
    }
    
    this.days++
    this.prices = append(this.prices, price)
    this.spans = append(this.spans, spans)
    
    return spans
}