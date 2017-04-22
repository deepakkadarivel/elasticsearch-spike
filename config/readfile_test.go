package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetEssMapForValidFilePath(t *testing.T) {
	data := ReadFile("./mappings.json")
	assert.NotNil(t, data)
}

func TestReadFileForInvalidFilePath(t *testing.T) {
	assert.Panics(t, func() {
		ReadFile(MappingsJson)
	}, "Read File Panics on Invalid Path.")
}
