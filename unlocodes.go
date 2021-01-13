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

import (
	"strconv"
)

func UnlocodeFunction(s string) (res uint64) {
	for _, el := range s {
		if el == 'B' {
			res |= (1 << 10)
		}
		digit, err := strconv.ParseUint(string(el), 10, 16)
		if err != nil {
			continue
		}
		res |= (1 << digit)
	}
	return res
}
