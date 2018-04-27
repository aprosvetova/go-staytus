package staytus

import "encoding/json"

//CreateIssue creates a new issue. Services' and status' permalinks should be used as IDs
func (api *Staytus) CreateIssue(title, initialUpdate, state string, services []string, status string, notify bool) (*Issue, error) {
	request := &issueRequest{
		Title:         title,
		InitialUpdate: initialUpdate,
		State:         state,
		Services:      services,
		Status:        status,
		Notify:        notify,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	data, err := api.post("api/v1/issues/create", body)
	if err != nil {
		return nil, err
	}
	var issue Issue
	if err := json.Unmarshal(data, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

//UpdateIssue updates existing issue. Services' and status' permalinks should be used as IDs
func (api *Staytus) UpdateIssue(id int, text, state, status string, notify bool) (*IssueUpdate, error) {
	request := &issueRequest{
		ID:     id,
		Text:   text,
		State:  state,
		Status: status,
		Notify: notify,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	data, err := api.post("api/v1/issues/update", body)
	if err != nil {
		return nil, err
	}
	var issueUpdate IssueUpdate
	if err := json.Unmarshal(data, &issueUpdate); err != nil {
		return nil, err
	}
	return &issueUpdate, nil
}

//GetIssue returns an issue info. Requires issue's ID as an input
func (api *Staytus) GetIssue(id int) (*Issue, error) {
	request := &request{
		IssueID: id,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	data, err := api.post("api/v1/issues/info", body)
	if err != nil {
		return nil, err
	}
	var issue Issue
	if err := json.Unmarshal(data, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

//GetIssues gets a list of issues
func (api *Staytus) GetIssues() ([]Issue, error) {
	data, err := api.post("api/v1/issues/all", nil)
	if err != nil {
		return nil, err
	}
	var issues []Issue
	if err := json.Unmarshal(data, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}
