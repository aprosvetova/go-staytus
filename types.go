package staytus

import "time"

//Staytus is the main type that allows you to interact with Staytus API
type Staytus struct {
	BaseURL string
	Token   string
	Secret  string
}

//Service is a type that represents a service object
type Service struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Permalink   string    `json:"permalink"`
	Position    int       `json:"position"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      Status    `json:"status"`
}

//Status is a type that represents a service status object
type Status struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Permalink  string    `json:"permalink"`
	Color      string    `json:"color"`
	StatusType string    `json:"status_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//Issue is a type that represents an issue object
type Issue struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	State      string    `json:"state"`
	Identifier string    `json:"identifier"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Notify     bool      `json:"notify"`
	User       struct {
		ID    int    `json:"id"`
		Email string `json:"email_address"`
		Name  string `json:"name"`
	} `json:"user"`
	ServiceStatus Status        `json:"service_status"`
	Updates       []IssueUpdate `json:"updates"`
}

//IssueUpdate is a type that represents an issue update object
type IssueUpdate struct {
	ID         int       `json:"id"`
	State      string    `json:"state"`
	Text       string    `json:"text"`
	Identifier string    `json:"identifier"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Notify     bool      `json:"notify"`
	User       struct {
		ID    int    `json:"id"`
		Email string `json:"email_address"`
		Name  string `json:"name"`
	} `json:"user"`
	ServiceStatus Status `json:"service_status"`
}

//Subscriber is a type that represents a email subscriber object
type Subscriber struct {
	ID                int       `json:"id"`
	Email             string    `json:"email_address"`
	VerificationToken string    `json:"verification_token"`
	VerifiedAt        time.Time `json:"verified_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

//Message is a type that represents an email message sent for subscriber verification
type Message struct {
	MessageID string `json:"message_id"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

type request struct {
	ServicePermalink            string `json:"service,omitempty"`
	StatusPermalink             string `json:"status,omitempty"`
	IssueID                     int    `json:"issue,omitempty"`
	SubscriberEmail             string `json:"email_address,omitempty"`
	SubscriberVerificationToken string `json:"verification_token,omitempty"`
	SubscriberVerified          int    `json:"verified"`
}

type issueRequest struct {
	ID            int      `json:"id,omitempty"`
	Text          string   `json:"text,omitempty"`
	Title         string   `json:"title,omitempty"`
	InitialUpdate string   `json:"initial_update,omitempty"`
	State         string   `json:"state"`
	Services      []string `json:"services,omitempty"`
	Status        string   `json:"status"`
	Notify        bool     `json:"notify"`
}
