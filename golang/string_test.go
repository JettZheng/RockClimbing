package golang

import (
	"testing"
)
func BenchmarkAddStringsByAdd(b *testing.B) {
	type args struct {
		is []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				is: []string{"eix934u", ":", "comments"},
			},
			want: "eix934u:comments",
		},
	}
	for _, tt := range tests {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if got := AddStringsByAdd(tt.args.is...); got != tt.want {
					b.Errorf("error")
				}
			}
		})
	}
}

func BenchmarkAddStringsByByteBufferPool(b *testing.B) {
	type args struct {
		is []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				is: []string{"eix934u", ":", "comments"},
			},
			want: "eix934u:comments",
		},
	}
	for _, tt := range tests {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if got := AddStringsByByteBufferPool(tt.args.is...); got != tt.want {
					b.Errorf("error")
				}
			}
		})
	}
}


func BenchmarkAddStringsByByteBufferNoPool(b *testing.B) {
	type args struct {
		is []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				is: []string{"eix934u", ":", "comments"},
			},
			want: "eix934u:comments",
		},
	}
	for _, tt := range tests {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if got := AddStringsByByteBufferNoPool(tt.args.is...); got != tt.want {
					b.Errorf("error")
				}
			}
		})
	}
}



func BenchmarkAddStringsByStringBufferNoPool(b *testing.B) {
	type args struct {
		is []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				is: []string{"eix934u", ":", "comments"},
			},
			want: "eix934u:comments",
		},
	}

	for _, tt := range tests {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if got := AddStringsByStringBufferNoPool(tt.args.is...); got != tt.want {
					b.Errorf("error")
				}
			}
		})
	}
}

func BenchmarkAddStringsByStringBufferPool(b *testing.B) {
	type args struct {
		is []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				is: []string{"eix934u", ":", "comments"},
			},
			want: "eix934u:comments",
		},
	}

	for _, tt := range tests {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if got := AddStringsByStringBufferPool(tt.args.is...); got != tt.want {
					b.Errorf("error")
				}
			}
		})
	}
}