package entity

type ImagesResp struct {
	Start  int     `json:"start"`
	Count  int     `json:"count"`
	Total  int     `json:"total"`
	Images []Image `json:"images"`
}

type Image struct {
	Path string `json:"path"`
}
