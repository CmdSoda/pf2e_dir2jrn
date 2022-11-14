package dropbox

type SharedItem struct {
	Tag       string `json:".tag"`
	URL       string `json:"url"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	PathLower string `json:"path_lower"`
}

type SharedItemList []SharedItem
