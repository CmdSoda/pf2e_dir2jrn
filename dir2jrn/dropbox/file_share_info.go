package dropbox

import "time"

type FileShareInfo struct {
	AccessType struct {
		Tag string `json:".tag"`
	} `json:"access_type"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	PathDisplay string `json:"path_display"`
	PathLower   string `json:"path_lower"`
	Policy      struct {
		ACLUpdatePolicy struct {
			Tag string `json:".tag"`
		} `json:"acl_update_policy"`
		SharedLinkPolicy struct {
			Tag string `json:".tag"`
		} `json:"shared_link_policy"`
		ViewerInfoPolicy struct {
			Tag string `json:".tag"`
		} `json:"viewer_info_policy"`
	} `json:"policy"`
	PreviewURL  string    `json:"preview_url"`
	TimeInvited time.Time `json:"time_invited"`
}
