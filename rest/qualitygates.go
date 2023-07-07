package rest

import (
	"fmt"
	"net/http"
	"sonarqube-goclient/rest/qualitygates"
)

type QualityGates struct {
	client *Client
}

func (q *QualityGates) GetByProject(options *qualitygates.GetByProjectOptions) (*qualitygates.GetByProjectObject, error) {
	formatUrl := fmt.Sprintf("%s/%s?project=%s", q.client.GetSonarqubeUrl(), qualitygates.Endpoint, options.Project)
	req, err := http.NewRequest(http.MethodGet, formatUrl, nil)
	if err != nil {
		return nil, err
	}
	res := qualitygates.GetByProjectObject{}
	if err := q.client.SendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (q *QualityGates) Search(options *qualitygates.SearchOptions) (*qualitygates.SearchObject, error) {

}
