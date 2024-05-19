package main

import (
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"
)

type MatchingSystem interface {
	// Add a new user to the matching system and find any possible matches for the new user.
	// Returns the ID of the user.
	AddSinglePersonAndMatch(name string, height int, gender Gender, dates int) (ID, error)
	// Remove a user from the matching system so that the user cannot be matched anymore.
	RemoveSinglePerson(id ID) error
	// Find the most N possible matched single people, where N is a request parameter.
	// Gender is required
	QuerySinglePeople(n int, gender Gender) ([]*MatchRequest, error)

	// - A single person has four input parameters: name, height, gender, and number of
	// wanted dates.
	// - Boys can only match girls who have lower height. Conversely, girls match boys who
	// are taller.
	// - Once the girl and boy match, they both use up one date. When their number of dates
	// becomes zero, they should be removed from the matching system.
}

var (
	ErrInvalidGender    = errors.New("invalid gender")
	ErrInvalidDateCount = errors.New("invalid date count")
	ErrPersonNotFound   = errors.New("person not found")
	ErrInvalidQuerySize = errors.New("invalid query size")
)

var _ MatchingSystem = (*matchingSystem)(nil)

type matchingSystem struct {
	logger           *zap.SugaredLogger
	mutex            sync.Mutex
	nextID           int
	requests         map[ID]*Item[MatchRequest]
	maxQueueByHeight *Queue[MatchRequest]
	minQueueByHeight *Queue[MatchRequest]
}

func NewMatchingSystem(logger *zap.Logger) MatchingSystem {
	return &matchingSystem{
		logger:   logger.Sugar(),
		requests: make(map[ID]*Item[MatchRequest]),
		maxQueueByHeight: NewQ(func(a, b *MatchRequest) bool {
			if a.Height != b.Height {
				return a.Height > b.Height
			} else {
				return a.CreatedAt.Before(b.CreatedAt)
			}
		}),
		minQueueByHeight: NewQ(func(a, b *MatchRequest) bool {
			if a.Height != b.Height {
				return a.Height < b.Height
			} else {
				return a.CreatedAt.Before(b.CreatedAt)
			}
		}),
	}
}

func (m *matchingSystem) AddSinglePersonAndMatch(name string, height int, gender Gender, dates int) (ID, error) {
	if gender != GENDER_MALE && gender != GENDER_FEMALE {
		return 0, ErrInvalidGender
	}
	if dates <= 0 {
		return 0, ErrInvalidDateCount
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.nextID += 1
	id := m.nextID

	item := NewItem(&MatchRequest{
		UserID:    id,
		Name:      name,
		Height:    height,
		Gender:    gender,
		Dates:     dates,
		CreatedAt: time.Now(),
	})

	m.requests[id] = item

	if gender == GENDER_MALE {
		m.maxQueueByHeight.PushItem(item)
	} else if gender == GENDER_FEMALE {
		m.minQueueByHeight.PushItem(item)
	}

	m.logger.Infow("Person added", "user", item.Value)

	if gender == GENDER_FEMALE {
		m.matchForFemale()
	} else if gender == GENDER_MALE {
		m.matchForMale()
	}

	return id, nil
}

func (m *matchingSystem) RemoveSinglePerson(id ID) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.removeSinglePerson(id)
}

func (m *matchingSystem) removeSinglePerson(id ID) error {
	item := m.requests[id]
	if item == nil {
		return ErrPersonNotFound
	}

	delete(m.requests, id)

	if item.Value.Gender == GENDER_MALE {
		m.maxQueueByHeight.Remove(item.Index)
	} else if item.Value.Gender == GENDER_FEMALE {
		m.minQueueByHeight.Remove(item.Index)
	}

	m.logger.Infow("Person removed", "user", item.Value)

	return nil
}

func (m *matchingSystem) QuerySinglePeople(n int, gender Gender) ([]*MatchRequest, error) {
	if n <= 0 {
		return nil, ErrInvalidQuerySize
	}
	if gender != GENDER_MALE && gender != GENDER_FEMALE {
		return nil, ErrInvalidGender
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	var q *Queue[MatchRequest]
	if gender == GENDER_MALE {
		q = m.maxQueueByHeight
	} else if gender == GENDER_FEMALE {
		q = m.minQueueByHeight
	}

	size := min(n, q.Len())

	items := q.PopItems(size)

	result := Map(GetItemValue, items)

	q.PushItems(items...)

	return result, nil
}

func (m *matchingSystem) matchForMale() {
	maleItem := m.maxQueueByHeight.PeekItem()

	var femaleItems []*Item[MatchRequest]
	for i := 0; i < maleItem.Value.Dates &&
		m.minQueueByHeight.Len() > 0 &&
		m.minQueueByHeight.PeekItem().Value.Height <= maleItem.Value.Height; i++ {
		femaleItems = append(femaleItems, m.minQueueByHeight.PopItem())
	}

	maleItem.Value.Dates -= len(femaleItems)
	for _, item := range femaleItems {
		item.Value.Dates -= 1
		m.logger.Infow("Matched", "male", maleItem.Value, "female", item.Value)
	}

	m.minQueueByHeight.PushItems(femaleItems...)

	if maleItem.Value.Dates == 0 {
		m.removeSinglePerson(maleItem.Value.UserID)
	}
	for m.minQueueByHeight.Len() > 0 && m.minQueueByHeight.PeekItem().Value.Dates == 0 {
		m.removeSinglePerson(m.minQueueByHeight.PeekItem().Value.UserID)
	}
}

func (m *matchingSystem) matchForFemale() {
	femaleItem := m.minQueueByHeight.PeekItem()

	var maleItems []*Item[MatchRequest]
	for i := 0; i < femaleItem.Value.Dates &&
		m.maxQueueByHeight.Len() > 0 &&
		m.maxQueueByHeight.PeekItem().Value.Height <= femaleItem.Value.Height; i++ {
		maleItems = append(maleItems, m.maxQueueByHeight.PopItem())
	}

	femaleItem.Value.Dates -= len(maleItems)
	for _, item := range maleItems {
		item.Value.Dates -= 1
		m.logger.Infow("Matched", "female", femaleItem.Value, "male", item.Value)
	}

	m.maxQueueByHeight.PushItems(maleItems...)

	if femaleItem.Value.Dates == 0 {
		m.removeSinglePerson(femaleItem.Value.UserID)
	}
	for m.maxQueueByHeight.Len() > 0 && m.maxQueueByHeight.PeekItem().Value.Dates == 0 {
		m.removeSinglePerson(m.maxQueueByHeight.PeekItem().Value.UserID)
	}
}
