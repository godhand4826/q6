package service_test

import (
	"q6/pkg/entity"
	"q6/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestInvalidInput(t *testing.T) {
	m := service.NewMatchingSystem(zap.NewNop())

	id, err := m.AddSinglePersonAndMatch("Alice", 160, entity.GenderUnknown, 1)
	assert.Equal(t, service.ErrInvalidGender, err)
	assert.Equal(t, 0, id)

	id, err = m.AddSinglePersonAndMatch("Alice", 160, entity.GenderMale, 0)
	assert.Equal(t, service.ErrInvalidDateCount, err)
	assert.Equal(t, 0, id)

	requests, err := m.QuerySinglePeople(10, entity.GenderUnknown)
	assert.Equal(t, service.ErrInvalidGender, err)
	assert.Len(t, requests, 0)

	requests, err = m.QuerySinglePeople(0, entity.GenderUnknown)
	assert.Equal(t, service.ErrInvalidQuerySize, err)
	assert.Len(t, requests, 0)

	err = m.RemoveSinglePerson(0)
	assert.Equal(t, service.ErrPersonNotFound, err)
}

func TestMatchingSystemRemove(t *testing.T) {
	m := service.NewMatchingSystem(zap.NewNop())

	id, err := m.AddSinglePersonAndMatch("Alice", 160, entity.GenderFemale, 1)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, id)

	requests, err := m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 1)

	err = m.RemoveSinglePerson(id)
	assert.NoError(t, err)

	requests, err = m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 0)
}

func TestMatchingSystemMatchTwoFemale(t *testing.T) {
	m := service.NewMatchingSystem(zap.NewNop())

	female, err := m.AddSinglePersonAndMatch("Alice", 160, entity.GenderFemale, 1)
	assert.NoError(t, err)

	requests, err := m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 1)
	assert.Equal(t, female, requests[0].UserID)

	female2, err := m.AddSinglePersonAndMatch("Ivy", 190, entity.GenderFemale, 1)
	assert.NoError(t, err)

	requests, err = m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 2)
	assert.Equal(t, female, requests[0].UserID)
	assert.Equal(t, female2, requests[1].UserID)

	_, err = m.AddSinglePersonAndMatch("Andy", 200, entity.GenderMale, 2)
	assert.NoError(t, err)

	requests, err = m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 0)

	requests, err = m.QuerySinglePeople(10, entity.GenderMale)
	assert.NoError(t, err)
	assert.Len(t, requests, 0)
}

func TestMatchingSystemBothTwoDates(t *testing.T) {
	m := service.NewMatchingSystem(zap.NewNop())

	female, err := m.AddSinglePersonAndMatch("Alice", 160, entity.GenderFemale, 2)
	assert.NoError(t, err)

	requests, err := m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 1)
	assert.Equal(t, female, requests[0].UserID)

	_, err = m.AddSinglePersonAndMatch("Andy", 200, entity.GenderMale, 2)
	assert.NoError(t, err)

	requests, err = m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 1)
	assert.Equal(t, 1, requests[0].Dates)

	requests, err = m.QuerySinglePeople(10, entity.GenderMale)
	assert.NoError(t, err)
	assert.Len(t, requests, 1)
	assert.Equal(t, 1, requests[0].Dates)
}

func TestMatchingSystemMany(t *testing.T) {
	m := service.NewMatchingSystem(zap.NewNop())

	_, _ = m.AddSinglePersonAndMatch("Alice", 160, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Ivy", 190, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Ember", 180, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Jenny", 160, entity.GenderFemale, 2)

	_, _ = m.AddSinglePersonAndMatch("Bob", 180, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Edward", 180, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Michael", 190, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Jack", 160, entity.GenderMale, 1)

	requests, err := m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 2)
	assert.Equal(t, "Ember", requests[0].Name)
	assert.Equal(t, "Ivy", requests[1].Name)

	requests, err = m.QuerySinglePeople(10, entity.GenderMale)
	assert.NoError(t, err)
	assert.Len(t, requests, 1)
	assert.Equal(t, "Jack", requests[0].Name)
}

func TestMatchingSystemMany2(t *testing.T) {
	m := service.NewMatchingSystem(zap.NewNop())

	_, _ = m.AddSinglePersonAndMatch("Alice", 160, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Jack", 160, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Bob", 180, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Edward", 180, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Ivy", 190, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Michael", 190, entity.GenderMale, 1)
	_, _ = m.AddSinglePersonAndMatch("Ember", 180, entity.GenderFemale, 1)
	_, _ = m.AddSinglePersonAndMatch("Jenny", 160, entity.GenderFemale, 2)

	requests, err := m.QuerySinglePeople(10, entity.GenderFemale)
	assert.NoError(t, err)
	assert.Len(t, requests, 2)
	assert.Equal(t, "Jenny", requests[0].Name)
	assert.Equal(t, "Ember", requests[1].Name)

	requests, err = m.QuerySinglePeople(10, entity.GenderMale)
	assert.NoError(t, err)
	assert.Len(t, requests, 2)
	assert.Equal(t, "Michael", requests[0].Name)
	assert.Equal(t, "Edward", requests[1].Name)
}
