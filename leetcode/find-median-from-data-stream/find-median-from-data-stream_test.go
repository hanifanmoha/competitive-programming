package findmedianfromdatastream

import "testing"

// Note : need to balance the tree

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name    string
		actions []string
		values  [][]int
		want    []float64
	}{
		{
			name:    "test1",
			actions: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			values:  [][]int{{}, {3}, {1}, {}, {2}, {}, {4}, {}},
			want:    []float64{2.0, 2.0, 2.5},
		},
		{
			name:    "test2",
			actions: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			values:  [][]int{{}, {1}, {2}, {}, {3}, {}},
			want:    []float64{1.5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			medianFinder := Constructor()
			var got float64
			wantIdx := 0
			for i, action := range tt.actions {
				switch action {
				case "addNum":
					medianFinder.AddNum(tt.values[i][0])
				case "findMedian":
					got = medianFinder.FindMedian()
					if got != tt.want[wantIdx] {
						t.Errorf("medianFinder.FindMedian() = %v, want %v", got, tt.want[wantIdx])
					}
					wantIdx++
				}
			}
		})
	}
}
