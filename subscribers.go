package staytus

import (
	"encoding/json"
)

//GetSubscriber returns a subscriber info. Requires subscriber's email OR verification token
func (api *Staytus) GetSubscriber(email, verificationToken string) (*Subscriber, error) {
	request := &request{
		SubscriberEmail:             email,
		SubscriberVerificationToken: verificationToken,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	data, err := api.post("api/v1/subscribers/info", body)
	if err != nil {
		return nil, err
	}
	var subscriber Subscriber
	if err := json.Unmarshal(data, &subscriber); err != nil {
		return nil, err
	}
	return &subscriber, nil
}

//AddSubscriber adds a subscriber as a notifications receiver. Requires subscriber's email
//
//Set verified true to add subscriber's email as already verified
func (api *Staytus) AddSubscriber(email string, verified bool) (*Subscriber, error) {
	request := &request{
		SubscriberEmail: email,
	}
	if verified {
		request.SubscriberVerified = 1
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	data, err := api.post("api/v1/subscribers/add", body)
	if err != nil {
		return nil, err
	}
	var subscriber Subscriber
	if err := json.Unmarshal(data, &subscriber); err != nil {
		return nil, err
	}
	return &subscriber, nil
}

//VerifySubscriber verifies an already added subscriber. Requires subscriber's email
func (api *Staytus) VerifySubscriber(email string) (bool, error) {
	request := &request{
		SubscriberEmail: email,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return false, err
	}
	data, err := api.post("api/v1/subscribers/verify", body)
	if err != nil {
		return false, err
	}
	var result bool
	if err := json.Unmarshal(data, &result); err != nil {
		return false, err
	}
	return result, nil
}

//SendVerificationEmail sends a verification email to already added subscriber. Requires subscriber's email
func (api *Staytus) SendVerificationEmail(email string) (*Message, error) {
	request := &request{
		SubscriberEmail: email,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	data, err := api.post("api/v1/subscribers/send_verification_email", body)
	if err != nil {
		return nil, err
	}
	var message Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, err
	}
	return &message, nil
}

//DeleteSubscriber deletes an already added subscriber. Requires subscriber's email
func (api *Staytus) DeleteSubscriber(email string) (bool, error) {
	request := &request{
		SubscriberEmail: email,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return false, err
	}
	data, err := api.post("api/v1/subscribers/delete", body)
	if err != nil {
		return false, err
	}
	return string(data) != "[]", nil
}
