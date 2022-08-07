package besttimetobuyandsellstock

// Method 1
// 概念是找有獲利的波段，並加總
func BestTimeToBuyAndSellStock2(prices []int) int {
	minStock := prices[0]
	tmpMaxProfit := 0
	totalProfit := 0

	for i := 1; i < len(prices); i++ {
		// 後面比前面小，要替換 minSotkc
		if prices[i] < prices[i-1] {
			minStock = prices[i]
			if tmpMaxProfit > 0 {
				totalProfit += tmpMaxProfit
			}
			tmpMaxProfit = 0
		} else {
			tmpMaxProfit = prices[i] - minStock
		}
	}

	// 最後還要加上 tmpMaxProfit 的原因在於
	// [4,6,10] 這種最後沒有變小的 case
	return totalProfit + tmpMaxProfit
}

// Method 2
// 也是加總獲利的部分，相較於 Method 1，更為簡潔
func BestTimeToBuyAndSellStock2_2(prices []int) int {
	profit := 0

	for i := range prices {
		if i == 0 {
			continue
		}

		//compare current with previous
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}

	return profit
}
