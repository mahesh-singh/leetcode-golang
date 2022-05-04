package max_profit

import "testing"

func TestMaxProfit(t *testing.T) {
	t.Run("Input [7, 1, 5, 3, 6, 4]", func(t *testing.T) {
		prices := [6]int{7, 1, 5, 3, 6, 4}
		got := MaxProfit(prices[:])
		want := 5
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, prices)
		}
	})

	t.Run("Input [7,6,4,3,1]", func(t *testing.T) {
		prices := [5]int{7, 6, 4, 3, 1}
		got := MaxProfit(prices[:])
		want := 0
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, prices)
		}
	})
}
