package rest

import (
	"net/http"
)

type QualityGates struct {
	client *Client
}

func (q *QualityGates) GetByProject(opt *GetByProjectOptions) (*GetByProjectObject, error) {
	u := q.client.SetUrlOpt("api/qualitygates/get_by_project", &opt)
	obj := GetByProjectObject{}
	if err := q.client.SendRequest(http.MethodGet, u, &obj); err != nil {
	}
	return &obj, nil
}

func (q *QualityGates) Search(opt *SearchOptions) (*SearchObject, error) {
	u := q.client.SetUrlOpt("api/qualitygates/search", &opt)
	obj := SearchObject{}
	if err := q.client.SendRequest(http.MethodGet, u, &obj); err != nil {
	}
	return &obj, nil
}
