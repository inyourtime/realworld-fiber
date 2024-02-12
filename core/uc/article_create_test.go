package uc

import (
	"realworld-go-fiber/core/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterTagNotExist(t *testing.T) {
	existingTags := []domain.Tag{
		{Name: "tag1"},
		{Name: "tag2"},
		{Name: "tag3"},
	}

	incomingTags := []string{
		"tag1",
		"tag4",
	}

	result := FilterTagNotExist(existingTags, incomingTags)

	assert.Len(t, result, 1)
	assert.Equal(t, "tag4", result[0].Name)
}
