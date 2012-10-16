package sublogic

import (
)

func outOfRange(val float32) bool {
	if (val < 0) || (val > 1) {
		return true
	}
	return false
}