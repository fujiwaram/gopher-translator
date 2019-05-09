package main

import (
	"testing"
)

// go run main.go '[[["moveTo",[[212,425]]],["curveTo",[[241,510],[266,587],[291,675]]],["lineTo",[[295,675]]],["curveTo",[[321,587],[345,510],[375,425]]],["lineTo",[[419,293]]],["lineTo",[[167,293]]],["closePath",[]],["moveTo",[[8,0]]],["lineTo",[[68,0]]],["lineTo",[[151,245]]],["lineTo",[[435,245]]],["lineTo",[[516,0]]],["lineTo",[[580,0]]],["lineTo",[[325,729]]],["lineTo",[[263,729]]],["closePath",[]]],[["moveTo",[[106,0]]],["lineTo",[[322,0]]],["curveTo",[[486,0],[594,72],[594,213]]],["curveTo",[[594,313],[532,372],[439,388]]],["lineTo",[[439,393]]],["curveTo",[[511,415],[550,477],[550,552]]],["curveTo",[[550,675],[454,729],[307,729]]],["lineTo",[[106,729]]],["closePath",[]],["moveTo",[[166,411]]],["lineTo",[[166,680]]],["lineTo",[[294,680]]],["curveTo",[[423,680],[490,644],[490,545]]],["curveTo",[[490,462],[432,411],[288,411]]],["closePath",[]],["moveTo",[[166,50]]],["lineTo",[[166,363]]],["lineTo",[[308,363]]],["curveTo",[[453,363],[535,315],[535,213]]],["curveTo",[[535,100],[449,50],[308,50]]],["closePath",[]]],[["moveTo",[[368,-13]]],["curveTo",[[462,-13],[530,26],[586,91]]],["lineTo",[[552,129]]],["curveTo",[[500,72],[445,41],[371,41]]],["curveTo",[[218,41],[123,168],[123,367]]],["curveTo",[[123,564],[220,688],[374,688]]],["curveTo",[[441,688],[493,659],[532,615]]],["lineTo",[[566,654]]],["curveTo",[[527,700],[460,742],[374,742]]],["curveTo",[[189,742],[60,598],[60,365]]],["curveTo",[[60,132],[188,-13],[368,-13]]],["closePath",[]]]]'

func Test_translate(t *testing.T) {
	type args struct {
		text    string
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ABC",
			args: args{
				text:    testInputABC,
				message: "ABC",
			},
			want:    testOutputABC,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translate(tt.args.text, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	testInputABC = `
	[
		[
			["moveTo", [[212, 425]]],
			["curveTo", [[241, 510], [266, 587], [291, 675]]],
			["lineTo", [[295, 675]]], 
			["curveTo", [[321, 587], [345, 510], [375, 425]]], 
			["lineTo", [[419, 293]]], 
			["lineTo", [[167, 293]]], 
			["closePath", []],
			["moveTo", [[8, 0]]], 
			["lineTo", [[68, 0]]],
			["lineTo", [[151, 245]]], 
			["lineTo", [[435, 245]]], 
			["lineTo", [[516, 0]]], 
			["lineTo", [[580, 0]]], 
			["lineTo", [[325, 729]]], 
			["lineTo", [[263, 729]]], 
			["closePath", []]
		], 
		[
			["moveTo", [[106, 0]]], 
			["lineTo", [[322, 0]]], 
			["curveTo", [[486, 0], [594, 72], [594, 213]]], 
			["curveTo", [[594, 313], [532, 372], [439, 388]]], 
			["lineTo", [[439, 393]]], 
			["curveTo", [[511, 415], [550, 477], [550, 552]]], 
			["curveTo", [[550, 675], [454, 729], [307, 729]]], 
			["lineTo", [[106, 729]]], 
			["closePath", []], 
			["moveTo", [[166, 411]]], 
			["lineTo", [[166, 680]]], 
			["lineTo", [[294, 680]]], 
			["curveTo", [[423, 680], [490, 644], [490, 545]]], 
			["curveTo", [[490, 462], [432, 411], [288, 411]]], 
			["closePath", []], 
			["moveTo", [[166, 50]]], 
			["lineTo", [[166, 363]]], 
			["lineTo", [[308, 363]]], 
			["curveTo", [[453, 363], [535, 315], [535, 213]]], 
			["curveTo", [[535, 100], [449, 50], [308, 50]]], 
			["closePath", []]
		],
		[
			["moveTo", [[368, -13]]], 
			["curveTo", [[462, -13], [530, 26], [586, 91]]], 
			["lineTo", [[552, 129]]], 
			["curveTo", [[500, 72], [445, 41], [371, 41]]], 
			["curveTo", [[218, 41], [123, 168], [123, 367]]], 
			["curveTo", [[123, 564], [220, 688], [374, 688]]], 
			["curveTo", [[441, 688], [493, 659], [532, 615]]], 
			["lineTo", [[566, 654]]], 
			["curveTo", [[527, 700], [460, 742], [374, 742]]], 
			["curveTo", [[189, 742], [60, 598], [60, 365]]], 
			["curveTo", [[60, 132], [188, -13], [368, -13]]], 
			["closePath", []]
		]
	]`

	testOutputABC = `draw mode
color off
right 303.000000
forward 7.883660
color black
right 76.000000
forward 0.898109
right 359.000000
forward 0.809568
right 358.000000
forward 0.914822
right 74.000000
forward 0.040000
right 74.000000
forward 0.917606
right 359.000000
forward 0.806536
right 358.000000
forward 0.901388
right 1.000000
forward 1.391402
right 108.000000
forward 2.520000
right 109.000000
forward 1.394597
color off
right 187.000000
forward 4.714244
color black
right 244.000000
forward 0.600000
right 289.000000
forward 2.586774
right 71.000000
forward 2.840000
right 72.000000
forward 2.580426
right 288.000000
forward 0.640000
right 251.000000
forward 7.723121
right 289.000000
forward 0.620000
right 290.000000
forward 7.723121
color off
right 250.000000
forward 6.780000
color black
right 0.000000
forward 2.160000
right 0.000000
forward 1.640000
right 327.000000
forward 1.297998
right 303.000000
forward 1.410000
right 0.000000
forward 1.000000
right 314.000000
forward 0.855862
right 326.000000
forward 0.943663
right 80.000000
forward 0.050000
right 74.000000
forward 0.752861
right 319.000000
forward 0.732462
right 327.000000
forward 0.750000
right 0.000000
forward 1.230000
right 300.000000
forward 1.101454
right 330.000000
forward 1.470000
right 0.000000
forward 2.010000
right 270.000000
forward 7.290000
color off
right 189.000000
forward 4.153565
color black
right 351.000000
forward 2.690000
right 90.000000
forward 1.280000
right 0.000000
forward 1.290000
right 29.000000
forward 0.760592
right 61.000000
forward 0.990000
right 0.000000
forward 0.830000
right 49.000000
forward 0.772334
right 41.000000
forward 1.440000
right 0.000000
forward 1.220000
color off
right 270.000000
forward 3.610000
color black
right 180.000000
forward 3.130000
right 90.000000
forward 1.420000
right 0.000000
forward 1.450000
right 31.000000
forward 0.950158
right 59.000000
forward 1.020000
right 0.000000
forward 1.130000
right 60.000000
forward 0.994786
right 30.000000
forward 1.410000
right 0.000000
forward 1.420000
color off
right 185.000000
forward 7.984892
color black
right 355.000000
forward 0.940000
right 331.000000
forward 0.783901
right 340.000000
forward 0.857963
right 278.000000
forward 0.509902
right 264.000000
forward 0.771557
right 18.000000
forward 0.631348
right 29.000000
forward 0.740000
right 0.000000
forward 1.530000
right 54.000000
forward 1.586001
right 36.000000
forward 1.990000
right 0.000000
forward 1.970000
right 39.000000
forward 1.574325
right 51.000000
forward 1.540000
right 0.000000
forward 0.670000
right 30.000000
forward 0.595399
right 19.000000
forward 0.587963
right 263.000000
forward 0.517397
right 278.000000
forward 0.603075
right 343.000000
forward 0.790759
right 327.000000
forward 0.860000
right 0.000000
forward 1.850000
right 312.000000
forward 1.933313
right 318.000000
forward 2.330000
right 0.000000
forward 2.330000
right 319.000000
forward 1.934141
right 311.000000
forward 1.800000
right 0.000000
forward 0.000000
color off
right 359.000000
forward 4.182021
color off
right 91.000000
forward 0.001000
say ABC
`
)
