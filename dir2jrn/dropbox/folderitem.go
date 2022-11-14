package dropbox

type DropboxFolderItem struct {
	Tag         string `json:".tag"`
	PathDisplay string `json:"path_display"`
	//Name                 string    `json:"name"`
	//PathLower            string    `json:"path_lower"`
	//ParentSharedFolderID string    `json:"parent_shared_folder_id"`
	//ID                   string    `json:"id"`
	//ClientModified       time.Time `json:"client_modified"`
	//ServerModified       time.Time `json:"server_modified"`
	//Rev                  string    `json:"rev"`
	//Size                 int       `json:"size"`
	/*
		SharingInfo          struct {
			ReadOnly             bool   `json:"read_only"`
			ParentSharedFolderID string `json:"parent_shared_folder_id"`
			ModifiedBy           string `json:"modified_by"`
		} `json:"sharing_info"`
		IsDownloadable bool   `json:"is_downloadable"`
		ContentHash    string `json:"content_hash"`
	*/
}
