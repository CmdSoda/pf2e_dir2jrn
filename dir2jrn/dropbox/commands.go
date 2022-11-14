package dropbox

import (
	"encoding/json"
	"os/exec"
	"strings"
)

func ExecuteFileShareInfoCommand(filepath string) (*FileShareInfo, error) {
	cmd := exec.Command("tbx", "file", "share", "info", "-path", filepath, "-output", "json")
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	var fsi FileShareInfo
	err = json.Unmarshal(out, &fsi)

	if err != nil {
		return nil, err
	}

	return &fsi, nil
}

func ExecuteSharedItemsCommand() (SharedItemList, error) {
	cmd := exec.Command("tbx", "sharedlink", "list", "-output", "json")
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	var sil SharedItemList = SharedItemList{}

	output := string(out)
	outputLines := strings.Split(strings.ReplaceAll(output, "\r\n", "\n"), "\n")
	for _, line := range outputLines {
		if line != "" {
			var si SharedItem
			err = json.Unmarshal([]byte(line), &si)
			if err == nil {
				sil = append(sil, si)
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return sil, nil
}

func ExecuteFolderCommand(folder string) ([]DropboxFolderItem, error) {
	cmd := exec.Command("tbx", "file", "list", "-path", folder, "-output", "json")
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	var fil []DropboxFolderItem = []DropboxFolderItem{}

	output := string(out)
	outputLines := strings.Split(strings.ReplaceAll(output, "\r\n", "\n"), "\n")
	for _, line := range outputLines {
		if line != "" {
			var fi DropboxFolderItem
			err = json.Unmarshal([]byte(line), &fi)
			if err == nil {
				fil = append(fil, fi)
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return fil, nil
}
