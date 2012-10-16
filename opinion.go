package sublogic

import (
	"errors"
)

type Opinion struct {
	Belief      float32
	Disbelief   float32
	Uncertainty float32
	Baserate    float32
}

// Creates a new opinion.
// 
// Opinion must satisfy the following:
// - belief + disbelief + uncertainty = 1
// - Each component must be in the range of 0 to 1
func NewOpinion(belief, disbelief, uncertainty, baserate float32) (*Opinion, error) {
	// Test for valid opinion
	if val := belief + disbelief + uncertainty; val != 1 {
		return nil, errors.New("The sum of all opinion components doesn't equal 1")
	}
	
	outOfRange := func(val float32) bool {
		if (val < 0) || (val > 1) {
			return true
		}
		return false
	}
	
	if outOfRange(belief) || outOfRange(disbelief) || outOfRange(uncertainty) || outOfRange(baserate) {
		return nil, errors.New("Opinion component isn't in range of 0 to 1")
	}
	
	
	opinion := Opinion{ belief, disbelief, uncertainty, baserate }
	return &opinion, nil
}

