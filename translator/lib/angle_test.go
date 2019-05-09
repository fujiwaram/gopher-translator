package lib

import "testing"

func Test_Angle_RightDiff(t *testing.T) {
	type args struct {
		dst Angle
	}
	tests := []struct {
		name string
		src  Angle
		args args
		want Angle
	}{
		{
			name: "90-89",
			src:  90,
			args: args{dst: 89},
			want: 1,
		},
		{
			name: "90-0",
			src:  90,
			args: args{dst: 0},
			want: 90,
		},
		{
			name: "90-225",
			src:  90,
			args: args{dst: 225},
			want: 225,
		},
		{
			name: "0-350",
			src:  0,
			args: args{dst: 350},
			want: 10,
		},
		{
			name: "45-345",
			src:  45,
			args: args{dst: 345},
			want: 60,
		},
		{
			name: "45-60",
			src:  45,
			args: args{dst: 60},
			want: 345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.src.RightDiff(tt.args.dst); got != tt.want {
				t.Errorf("Angle.RightDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
