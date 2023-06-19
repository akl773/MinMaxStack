package stack

// Stat holds the minimum and maximum values.
type Stat struct {
	Min int
	Max int
}

// MinMaxStack represents a stack that tracks minimum and maximum values.
type MinMaxStack struct {
	values []int
	stats  []Stat
}

// NewMinMaxStack creates a new MinMaxStack with an initial value.
func NewMinMaxStack(val int) *MinMaxStack {
	return &MinMaxStack{
		values: []int{val},
		stats:  []Stat{{val, val}},
	}
}
