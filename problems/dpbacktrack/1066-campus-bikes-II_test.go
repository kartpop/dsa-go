package dbbacktrack

import "testing"

func TestAssignBikes(t *testing.T) {
	workers := [][]int{{815,60},{638,626},{6,44},{103,90},{591,880}}
	bikes := [][]int{{709,161},{341,339},{755,955},{172,27},{433,489}}
	ans := AssignBikes(workers, bikes)

	if ans != 1458 {
		t.Errorf("Expected %d, got %d", 1458, ans)
	}
}

