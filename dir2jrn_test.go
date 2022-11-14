package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortFileNameAscend(t *testing.T) {
	localpath := "S:\\Data\\Bilder\\Books\\Grundregelwerk\\"

	files, err := GetFiles(localpath)
	assert.Nil(t, err)

	SortFileNameAscend(files)

}
