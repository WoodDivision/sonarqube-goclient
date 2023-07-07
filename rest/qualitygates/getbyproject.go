package qualitygates

type GetByProjectOptions struct {
	Project string
}

type GetByProjectObject struct {
	QualityGate QualityGate `json:"qualityGate"`
}

type QualityGate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

const Endpoint string = "api/qualitygates/get_by_project"
