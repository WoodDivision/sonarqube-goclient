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

func (q *QualityGates) List() (*ListObject, error) {
	u := q.client.SetUrlOpt("api/qualitygates/list", nil)
	obj := ListObject{}
	if err := q.client.SendRequest(http.MethodGet, u, &obj); err != nil {
	}
	return &obj, nil
}

func (q *QualityGates) ProjectStatus(opt *ProjectStatusOptions) (*ProjectStatusObject, error) {
	u := q.client.SetUrlOpt("api/qualitygates/project_status", &opt)
	obj := ProjectStatusObject{}
	if err := q.client.SendRequest(http.MethodGet, u, &obj); err != nil {
	}
	return &obj, nil
}

func (q *QualityGates) Show(opt *ShowOptions) (*ShowObject, error) {
	u := q.client.SetUrlOpt("api/qualitygates/show", &opt)
	obj := ShowObject{}
	if err := q.client.SendRequest(http.MethodGet, u, &obj); err != nil {
	}
	return &obj, nil
}
