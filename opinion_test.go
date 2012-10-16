package sublogic

import (
	"testing"
)

func Test_Discount(t *testing.T) () {
	opinionA, _ := NewOpinion(0.7, 0.0, 0.3, 0.5) // A's opinion of B
	opinionB, _ := NewOpinion(0.4, 0.1, 0.5, 0.5) // B's opinion of X
	opinionC := opinionA.Discount(opinionB) // A's derived opinion of X
	
	if (opinionC.Belief != 0.28) || (opinionC.Disbelief != 0.07) || (opinionC.Uncertainty != 0.65) {
		t.Error("Opinion discount calculation was not correct!")
	}
	
	/*
	Expectation:
	opinionA - 0.85
	opinionB - 0.65
	opinionC - 0.61
	*/
}