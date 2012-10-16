/*
   Copyright Liam (liamzebedee) Edwards-Playne 2012

   This file is part of sublogic-go.

   sublogic-go is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   sublogic-go is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with sublogic-go.  If not, see <http://www.gnu.org/licenses/>.
*/
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

