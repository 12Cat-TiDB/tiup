package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateAndParserD1agHeader(t *testing.T) {
	assert := require.New(t)

	mirror, comp, tag, err := ParseComponentVersion("tiup.io/playground:v1.2.3")
	assert.Nil(err)
	assert.EqualValues("tiup.io", mirror)
	assert.EqualValues("playground", comp)
	assert.EqualValues("v1.2.3", tag)
}
