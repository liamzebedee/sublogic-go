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
	
	if outOfRange(belief) || outOfRange(disbelief) || outOfRange(uncertainty) || outOfRange(baserate) {
		return nil, errors.New("Opinion component isn't in range of 0 to 1")
	}
	
	opinion := Opinion{ belief, disbelief, uncertainty, baserate }
	return &opinion, nil
}

func newEmptyOpinion() (*Opinion) {
	return &Opinion{ }
}

// Where A trusts B, and B trusts X, this returns A's transitive trust in X
func (A *Opinion) Discount(B *Opinion) (C *Opinion) {
	C = newEmptyOpinion()
	
	// TODO check A and B for consistency
	C.Belief = A.Belief * B.Belief
	C.Disbelief = A.Belief * B.Disbelief
	C.Uncertainty = (A.Disbelief + A.Uncertainty + A.Belief*B.Uncertainty)
	C.Baserate = B.Baserate
	
	return C
}