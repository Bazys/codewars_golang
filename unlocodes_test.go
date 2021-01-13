/**
 * Write function for convert string like '-2345---'
 * to integer value = 60. 60 bitmask calculated
 * as 2^2 + 2^3 + 2^4 + 2^5
 * 0	Function not known, to be specified	          = 0
 * 1	Port, as defined in Rec 16	                  = 1
 * 2	Rail Terminal	                                = 2
 * 3	Road Terminal	                                = 3
 * 4	Airport	                                      = 4
 * 5	Postal Exchange Office                        = 5
 * 6	Multimodal Functions (ICDs, etc.)	            = 6
 * 7	Fixed Transport Functions (e.g. Oil platform)	= 7
 * 8	Inland Port	                                  = 8
 * B	Border Crossing	                              = 10
**/
package kata

import "testing"

func TestUnlocodeFunction(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantRes uint64
	}{
		{
			name:    "test1",
			args:    "-2345---",
			wantRes: 60,
		},
		{
			name:    "test2",
			args:    "0-------",
			wantRes: 1,
		},
		{
			name:    "test3",
			args:    "--3----B",
			wantRes: 1032,
		},
		{
			name:    "test4",
			args:    "--X----B  ",
			wantRes: 1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := UnlocodeFunction(tt.args); gotRes != tt.wantRes {
				t.Errorf("UnlocodeFunction() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
