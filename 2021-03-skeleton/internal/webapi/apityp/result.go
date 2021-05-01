package apityp

// ResultJSON 単数正常系レスポンス
type ResultJSON struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

// ResultsJSON 複数正常系レスポンス
type ResultsJSON struct {
	Results    interface{} `json:"results"`
	Size       int         `json:"size"`
	TotalSize  int64       `json:"total_size"`
	Page       int         `json:"page"`
	TotalPages int         `json:"total_pages"`
	Error      string      `json:"error,omitempty"`
}
