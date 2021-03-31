package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotImgUrl(t *testing.T) {
	for _, url := range []string{
		"https://www.host/comments?post=74573",
		"https://youtu.be/qFdvwSyd_zs",
		"https://www.host/comments.csv",
		"",
	} {
		assert.Equal(t, false, checkUrl(url), "they should be equal")
	}
}

func TestImgUrl(t *testing.T) {
	for _, url := range []string{
		"https://i.imgur.com/dJJe2N7.gifv",
		"https://yo.be/qFdvwSyd_zs.jpg",
		"https://www.host/comments.png",
	} {
		assert.Equal(t, true, checkUrl(url), "they should be equal")
	}
}
