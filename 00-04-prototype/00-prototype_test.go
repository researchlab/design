package prototype

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKeywords_Clone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := Keywords{
		"testA": &Keyword{
			word:     "testA",
			visit:    1,
			UpdateAt: &updateAt,
		},
		"testB": &Keyword{
			word:     "testB",
			visit:    2,
			UpdateAt: &updateAt,
		},
		"testC": &Keyword{
			word:     "testC",
			visit:    3,
			UpdateAt: &updateAt,
		},
	}
	now := time.Now()
	updateWords := []*Keyword{
		{
			word:     "testB",
			visit:    10,
			UpdateAt: &now,
		},
	}
	got := words.Clone(updateWords)
	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updateWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
