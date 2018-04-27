package staytus

import "encoding/json"

//GetService returns a service info. Requires service's permalink as an input
func (api *Staytus) GetService(permalink string) (*Service, error) {
	request := &request{
		ServicePermalink: permalink,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	data, err := api.post("api/v1/services/info", body)
	if err != nil {
		return nil, err
	}
	var service Service
	if err := json.Unmarshal(data, &service); err != nil {
		return nil, err
	}
	return &service, nil
}

//GetServices gets a list of services
func (api *Staytus) GetServices() ([]Service, error) {
	data, err := api.post("api/v1/services/all", nil)
	if err != nil {
		return nil, err
	}
	var services []Service
	if err := json.Unmarshal(data, &services); err != nil {
		return nil, err
	}
	return services, nil
}

//SetServiceStatus sets the status for a service. Requires service's and status' permalink as an input
func (api *Staytus) SetServiceStatus(servicePermalink, statusPermalink string) (*Service, error) {
	request := &request{
		ServicePermalink: servicePermalink,
		StatusPermalink:  statusPermalink,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	data, err := api.post("api/v1/services/set_status", body)
	if err != nil {
		return nil, err
	}
	var service Service
	if err := json.Unmarshal(data, &service); err != nil {
		return nil, err
	}
	return &service, nil
}
