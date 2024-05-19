package main

import (
	"testing"

	"go.uber.org/zap"
)

func TestMatchingSystem(t *testing.T) {
	m := NewMatchingSystem(zap.NewExample())

	m.AddSinglePersonAndMatch("Alice", 160, GENDER_FEMALE, 1)
	m.AddSinglePersonAndMatch("Ivy", 190, GENDER_FEMALE, 1)
	m.AddSinglePersonAndMatch("Ember", 180, GENDER_FEMALE, 1)
	m.AddSinglePersonAndMatch("Jenny", 160, GENDER_FEMALE, 2)

	m.AddSinglePersonAndMatch("Bob", 180, GENDER_MALE, 1)
	m.AddSinglePersonAndMatch("Edward", 180, GENDER_MALE, 1)
	m.AddSinglePersonAndMatch("Michael", 190, GENDER_MALE, 1)
	m.AddSinglePersonAndMatch("Jack", 160, GENDER_MALE, 1)
}
