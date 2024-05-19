package main

import (
	"testing"

	"go.uber.org/zap"
)

func TestMatchingSystem(_ *testing.T) {
	m := NewMatchingSystem(zap.NewExample())

	_, _ = m.AddSinglePersonAndMatch("Alice", 160, GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Ivy", 190, GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Ember", 180, GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Jenny", 160, GenderFemale, 2)

	_, _ = m.AddSinglePersonAndMatch("Bob", 180, GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Edward", 180, GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Michael", 190, GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Jack", 160, GenderMale, 1)
}
