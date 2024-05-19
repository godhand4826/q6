package service

import (
	"q6/pkg/entity"
	"testing"

	"go.uber.org/zap"
)

func TestMatchingSystem(_ *testing.T) {
	m := NewMatchingSystem(zap.NewExample())

	_, _ = m.AddSinglePersonAndMatch("Alice", 160, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Ivy", 190, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Ember", 180, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Jenny", 160, entity.GenderFemale, 2)

	_, _ = m.AddSinglePersonAndMatch("Bob", 180, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Edward", 180, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Michael", 190, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Jack", 160, entity.GenderMale, 1)
}
