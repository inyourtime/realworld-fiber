package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterTagNotExist(t *testing.T) {
	existingTags := []Tag{
		{Name: "tag1"},
		{Name: "tag2"},
		{Name: "tag3"},
	}

	incomingTags := []string{
		"tag1",
		"tag4",
	}

	result := []Tag{}
	FilterTagNotExist(existingTags, incomingTags, &result)

	assert.Len(t, result, 1)
	assert.Equal(t, "tag4", result[0].Name)
}
