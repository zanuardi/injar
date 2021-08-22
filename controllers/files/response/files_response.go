package response

type FileResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Path string `json:"path"`
}