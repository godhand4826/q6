package util_test

import (
	"q6/lib/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsRef(t *testing.T) {
	assert.Equal(t, 1, *util.AsRef(1))
	assert.Equal(t, "string", *util.AsRef("string"))
	assert.Equal(t, struct{}{}, *util.AsRef(struct{}{}))
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	res := util.Map(func(v int) int { return v + 1 }, arr)
	assert.Equal(t, []int{2, 3, 4}, res)
}
