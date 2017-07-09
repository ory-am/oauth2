package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "foo", ToString("foo"))
	assert.Equal(t, "foo", ToString([]string{"foo"}))
	assert.Empty(t, ToString(1234))
	assert.Empty(t, ToString(nil))
}

func TestToTime(t *testing.T) {
	assert.Equal(t, time.Time{}, ToTime(nil))
	assert.Equal(t, time.Time{}, ToTime("1234"))

	now := time.Now().Round(time.Second)
	assert.Equal(t, now, ToTime(now))
	assert.Equal(t, now, ToTime(now.Unix()))
	assert.Equal(t, now, ToTime(float64(now.Unix())))
}
