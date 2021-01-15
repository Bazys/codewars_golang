package kata

import (
	"reflect"
	"testing"
)

func TestIntToBitArray(t *testing.T) {
	tests := []struct {
		n       uint64
		wantRes []uint8
	}{
		{n: 16, wantRes: []uint8{4}},
		{n: 32, wantRes: []uint8{5}},
		{n: 1024, wantRes: []uint8{10}},
		{n: 8664, wantRes: []uint8{3, 4, 6, 7, 8, 13}},
		{n: 47612, wantRes: []uint8{2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 15}},
		{n: 47617, wantRes: []uint8{0, 9, 11, 12, 13, 15}},
	}
	for _, tt := range tests {
		if gotRes := IntToBitArray(tt.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
			t.Errorf("intToBitNoArray() = %v, want %v", gotRes, tt.wantRes)
		}
	}
}

func TestBitArrayToInt(t *testing.T) {
	tests := []struct {
		name string
		arr  []uint8
		want uint64
	}{
		{arr: []uint8{4}, want: uint64(16)},
		{arr: []uint8{5}, want: 32},
		{arr: []uint8{10}, want: 1024},
		{arr: []uint8{3, 4, 6, 7, 8, 13}, want: 8664},
		{arr: []uint8{2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 15}, want: 47612},
		{arr: []uint8{0, 9, 11, 12, 13, 15}, want: 47617},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BitArrayToInt(tt.arr); got != tt.want {
				t.Errorf("BitArrayToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
