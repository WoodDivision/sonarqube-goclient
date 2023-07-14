package rest

type CopyOptions struct {
}

type GetByProjectOptions struct {
	Project string `url:"project"`
}

type GetByProjectObject struct {
	QualityGate QualityGate `json:"qualityGate"`
}

type QualityGate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

type SearchObject struct {
	More    bool           `json:"more"`
	Results []SearchResult `json:"results"`
	Paging  SearchPaging   `json:"paging"`
}

type SearchResult struct {
	Name     string `json:"name"`
	Selected bool   `json:"selected"`
	Key      string `json:"key"`
}

type SearchPaging struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}

type SearchOptions struct {
	GateName string `url:"gateName"`
	Page     int    `url:"page"`
	PageSize int    `url:"pageSize"`
	Query    string `url:"query"`
	Selected string `url:"selected"`
}
