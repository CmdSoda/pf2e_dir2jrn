package dropbox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteFileShareInfoCommand(t *testing.T) {
	file := "/pathfinder/Bilder/Books/Grundregelwerk/Alles/0.jpg.webp"
	fsi, err := ExecuteFileShareInfoCommand(file)
	assert.Nil(t, err)
	assert.NotNil(t, fsi)
	assert.NotEmpty(t, fsi.PreviewURL)
}

func TestExecuteFileShareInfoCommand2(t *testing.T) {
	file := "/pathfinder/Bilder/Books/Grundregelwerk/Alles/gibtsnicht.jpg.webp"
	_, err := ExecuteFileShareInfoCommand(file)
	assert.NotNil(t, err)
}

func TestExecuteFileShareInfoCommand3(t *testing.T) {
	file := "/pathfinder/Bilder/Books/Grundregelwerk/Alles/38.jpg.webp"
	fsi, err := ExecuteFileShareInfoCommand(file)
	assert.Nil(t, err)
	assert.NotNil(t, fsi)
	assert.NotEmpty(t, fsi.PreviewURL)
}

func TestExecuteSharedItemList(t *testing.T) {
	sil, err := ExecuteSharedItemsCommand()
	assert.Nil(t, err)
	assert.NotNil(t, sil)
}

func TestExecuteFolderCommand(t *testing.T) {
	folder := "/pathfinder/Bilder/Books/Grundregelwerk/Alles/"
	fil, err := ExecuteFolderCommand(folder)
	assert.Nil(t, err)
	assert.NotNil(t, fil)
	fmt.Println(fil)
}
