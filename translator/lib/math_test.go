package lib

import (
	"testing"
)

func TestAnglePointToPoint(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0,0_5,5",
			args: args{
				p1: Point{X: 0, Y: 0},
				p2: Point{X: 5, Y: 5},
			},
			want: 45,
		},
		{
			name: "0,0_0,5",
			args: args{
				p1: Point{X: 0, Y: 0},
				p2: Point{X: 0, Y: 5},
			},
			want: 90,
		},
		{
			name: "0,0_5,0",
			args: args{
				p1: Point{X: 0, Y: 0},
				p2: Point{X: 5, Y: 0},
			},
			want: 0,
		},
		{
			name: "0,0_-5,0",
			args: args{
				p1: Point{X: 0, Y: 0},
				p2: Point{X: -5, Y: 0},
			},
			want: 180,
		},
		{
			name: "0,0_-5,-5",
			args: args{
				p1: Point{X: 0, Y: 0},
				p2: Point{X: -5, Y: -5},
			},
			want: 225,
		},
		{
			name: "0,0_0,-5",
			args: args{
				p1: Point{X: 0, Y: 0},
				p2: Point{X: 0, Y: -5},
			},
			want: 270,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnglePointToPoint(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("LengthPointToPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthPointToPoint(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0,0_3,4",
			args: args{
				p1: Point{X: 0, Y: 0},
				p2: Point{X: 3, Y: 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthPointToPoint(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("AnglePointToPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
