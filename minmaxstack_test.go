package minmaxstack

import (
	"testing"
)

func TestNewMinMaxStack(t *testing.T) {
	mms := NewMinMaxStack(5)
	if mms.Len() != 1 {
		t.Errorf("NewMinMaxStack: expected length 1, got %v", mms.Len())
	}
	if mms.GetMin() != 5 || mms.GetMax() != 5 {
		t.Errorf("NewMinMaxStack: expected min 5 and max 5, got min %v and max %v", mms.GetMin(), mms.GetMax())
	}
}

func TestMinMaxStack_Push(t *testing.T) {
	mms := NewMinMaxStack(5).Push(3).Push(7).Push(2)

	if mms.Len() != 4 {
		t.Errorf("Push: expected length 4, got %v", mms.Len())
	}
	if mms.GetMin() != 2 || mms.GetMax() != 7 {
		t.Errorf("Push: expected min 2 and max 7, got min %v and max %v", mms.GetMin(), mms.GetMax())
	}
}

func TestMinMaxStack_Pop(t *testing.T) {
	mms := NewMinMaxStack(5).Push(3).Push(7).Push(2)
	mms.Pop()

	if mms.Len() != 3 {
		t.Errorf("Pop: expected length 3, got %v", mms.Len())
	}
	if mms.GetMin() != 3 || mms.GetMax() != 7 {
		t.Errorf("Pop: expected min 3 and max 7, got min %v and max %v", mms.GetMin(), mms.GetMax())
	}
}

func TestMinMaxStack_PopEmpty(t *testing.T) {
	mms := NewMinMaxStack(5)
	mms.Pop()

	defer func() {
		if r := recover(); r == nil {
			t.Error("PopEmpty: expected panic but didn't get one")
		}
	}()

	mms.Pop()
}

func TestMinMaxStack_GetMin(t *testing.T) {
	mms := NewMinMaxStack(5).Push(3).Push(7).Push(2)

	if mms.GetMin() != 2 {
		t.Errorf("GetMin: expected 2, got %v", mms.GetMin())
	}
}

func TestMinMaxStack_GetMax(t *testing.T) {
	mms := NewMinMaxStack(5).Push(3).Push(7).Push(2)

	if mms.GetMax() != 7 {
		t.Errorf("GetMax: expected 7, got %v", mms.GetMax())
	}
}
