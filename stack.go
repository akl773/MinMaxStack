package minmaxstack

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

// Push adds a value to the stack.
func (m *MinMaxStack) Push(val int) {
	m.values = append(m.values, val)
	m.updateStat()
}

// Pop removes the topmost value from the stack.
func (m *MinMaxStack) Pop() {
	if m.Len() == 0 {
		panic("pop called on empty stack")
	}
	lastIdx := m.Len() - 1
	m.values = m.values[:lastIdx]
	m.stats = m.stats[:lastIdx]
}

// Len returns the number of elements in the stack.
func (m *MinMaxStack) Len() int {
	return len(m.values)
}

// GetMin returns the minimum value in the stack.
func (m *MinMaxStack) GetMin() int {
	return m.stats[m.Len()-1].Min
}

// GetMax returns the maximum value in the stack.
func (m *MinMaxStack) GetMax() int {
	return m.stats[m.Len()-1].Max
}

// updateStat recalculates the minimum and maximum values.
func (m *MinMaxStack) updateStat() {
	e := m.values[m.Len()-1]
	prevStat := m.stats[m.Len()-2]
	currStat := Stat{Min: e, Max: e}

	if e > prevStat.Max {
		currStat.Max = e
	} else {
		currStat.Max = prevStat.Max
	}

	if e < prevStat.Min {
		currStat.Min = e
	} else {
		currStat.Min = prevStat.Min
	}

	m.stats = append(m.stats, currStat)
}
