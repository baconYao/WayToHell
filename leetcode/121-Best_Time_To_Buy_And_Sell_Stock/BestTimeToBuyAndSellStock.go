package besttimetobuyandsellstock

import "math"

func BestTimeToBuyAndSellStock(prices []int) int {
	minStock := math.MaxInt64
	maxProfit := 0

	for _, v := range prices {
		if v <= minStock {
			minStock = v
		}
		if v-minStock > maxProfit {
			maxProfit = v - minStock
		}
	}

	return maxProfit
}
