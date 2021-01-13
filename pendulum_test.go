package kata

import (
	"reflect"
	"testing"
)

func TestPendulum(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"1", args{[]int{19, 30, 16, 19, 28, 26, 28, 17, 21, 17}}, []int{28, 26, 19, 17, 16, 17, 19, 21, 28, 30}},
		{"2", args{[]int{49, 40, 41, 41, 39, 43, 40, 46, 30, 47, 46, 40}}, []int{47, 46, 41, 40, 40, 30, 39, 40, 41, 43, 46, 49}},
		{"3", args{[]int{4, 10, 9}}, []int{10, 4, 9}},
		{"4", args{[]int{6, 6, 8, 5, 10}}, []int{10, 6, 5, 6, 8}},
		{"5", args{[]int{-7, -10, -1, -10, -8}}, []int{-1, -8, -10, -10, -7}},
		{"6", args{[]int{-2, -11, -6, -11, -4, -3, -5}}, []int{-2, -4, -6, -11, -11, -5, -3}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Pendulum(tt.args.values); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Pendulum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
