package rest

type CopyOptions struct {
}

type GetByProjectOptions struct {
	Project string `url:"project"`
}

type GetByProjectObject struct {
	QualityGate QualityGateGetByProject `json:"qualityGate"`
}

type QualityGateGetByProject struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Default bool   `json:"default,omitempty"`
}

type SearchObject struct {
	More    bool           `json:"more,omitempty"`
	Results []SearchResult `json:"results,omitempty"`
	Paging  SearchPaging   `json:"paging,omitempty"`
}

type SearchResult struct {
	Name     string `json:"name,omitempty"`
	Selected bool   `json:"selected,omitempty"`
	Key      string `json:"key,omitempty"`
}

type SearchPaging struct {
	PageIndex int `json:"pageIndex,omitempty"`
	PageSize  int `json:"pageSize,omitempty"`
	Total     int `json:"total,omitempty"`
}

type SearchOptions struct {
	GateName string `url:"gateName"`
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"pageSize,omitempty"`
	Query    string `url:"query,omitempty"`
	Selected string `url:"selected,omitempty"`
}

type ListObject struct {
	Qualitygates []QualityGatesList `json:"qualitygates,omitempty"`
	Default      int                `json:"default,omitempty"`
	Actions      ActionsList        `json:"actions,omitempty"`
}

type QualityGatesList struct {
	Id        string      `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	IsDefault bool        `json:"isDefault,omitempty"`
	IsBuiltIn bool        `json:"isBuiltIn,omitempty"`
	Actions   ActionsList `json:"actions,omitempty"`
}

type ActionsList struct {
	Create            bool `json:"create,omitempty"`
	Rename            bool `json:"rename,omitempty"`
	SetAsDefault      bool `json:"setAsDefault,omitempty"`
	Copy              bool `json:"copy,omitempty"`
	AssociateProjects bool `json:"associateProjects,omitempty"`
	Delete            bool `json:"delete,omitempty"`
	ManageConditions  bool `json:"manageConditions,omitempty"`
	Delegate          bool `json:"delegate,omitempty"`
}

type ProjectStatusObject struct {
	ProjectStatus ProjectStatus `json:"projectStatus"`
}
type ProjectStatus struct {
	Status            string       `json:"status"`
	IgnoredConditions bool         `json:"ignoredConditions"`
	CaycStatus        string       `json:"caycStatus"`
	Conditions        []Conditions `json:"conditions"`
	Periods           []Period     `json:"periods"`
	Period            Period       `json:"period"`
}

type Conditions struct {
	Status         string `json:"status,omitempty"`
	MetricKey      string `json:"metricKey,omitempty"`
	Comparator     string `json:"comparator,omitempty"`
	PeriodIndex    int    `json:"periodIndex,omitempty"`
	ErrorThreshold string `json:"errorThreshold,omitempty"`
	ActualValue    string `json:"actualValue,omitempty"`
	Id             string `json:"id,omitempty"`
	Metric         string `json:"metric,omitempty"`
	Op             string `json:"op,omitempty"`
	Error          string `json:"error,omitempty"`
}

type Period struct {
	Index     int    `json:"index,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Date      string `json:"date,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

type ProjectStatusOptions struct {
	AnalysisId  string `url:"analysisId,omitempty"`
	Branch      string `url:"branch,omitempty"`
	ProjectId   string `url:"projectId,omitempty"`
	ProjectKey  string `url:"projectKey,omitempty"`
	PullRequest int    `url:"pullRequest,omitempty"`
}
type ShowObject struct {
	Id         int          `json:"id"`
	Name       string       `json:"name"`
	Conditions []Conditions `json:"conditions"`
	IsBuiltIn  bool         `json:"isBuiltIn"`
	Actions    ActionsList  `json:"actions"`
}
type ShowOptions struct {
	Id   string `url:"id,omitempty"`
	Name string `url:"name,omitempty"`
}
