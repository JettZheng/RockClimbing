package luckydraw

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",args{
				a: []int{1,2,3,4,5},
				x: 1,
			},-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
