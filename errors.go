package vector

import "errors"

var (
	// ErrNot3Dimensional is an error that is returned in functions that only
	// supports 3 dimensional vectors
	ErrNot3Dimensional = errors.New("vector is not 3 dimensional")
	// ErrNotSameDimensions is an error that is returned when functions need both
	// Vectors provided to be the same dimensionally
	ErrNotSameDimensions = errors.New("the two vectors provided aren't the same dimensional size")
)
