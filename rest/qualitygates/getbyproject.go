package qualitygates

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sonarqube-goclient/client"
)

type Errors struct {
	Msg string `json:"msg"`
}

type ApiQualityGates struct {
	Client *client.Client
}

type QualityGateOptions struct {
	Project string
}

type QualityGates struct {
	QualityGate QualityGate `json:"qualityGate"`
}

type QualityGate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

const endpoint string = "api/qualitygates/get_by_project"

func (a *ApiQualityGates) GetByProject(options *QualityGateOptions) (*QualityGates, error) {
	formatUrl := fmt.Sprintf("%s/%s?project=%s", a.Client.GetSonarqubeUrl(), endpoint, options.Project)
	req, err := http.NewRequest("GET", formatUrl, nil)
	if err != nil {
		return nil, err
	}
	res := QualityGates{}
	if err := a.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *ApiQualityGates) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(a.Client.GetSonarqubeUser(), a.Client.GetSonarqubePass())

	res, err := a.Client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes Errors
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Msg)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}
	return nil
}
