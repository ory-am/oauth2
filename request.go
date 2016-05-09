package fosite

import (
	"net/url"
	"time"
)

// Request is an implementation of Requester
type Request struct {
	RequestedAt   time.Time
	Client        Client
	Scopes        Arguments
	GrantedScopes Arguments
	Form          url.Values
	Session       interface{}
}

func NewRequest() *Request {
	return &Request{
		Client: &DefaultClient{},
		Scopes: Arguments{},
		Form:   url.Values{},
	}
}

func (a *Request) GetRequestForm() url.Values {
	return a.Form
}

func (a *Request) GetRequestedAt() time.Time {
	return a.RequestedAt
}

func (a *Request) GetClient() Client {
	return a.Client
}

func (a *Request) GetScopes() Arguments {
	return a.Scopes
}

func (a *Request) SetScopes(s Arguments) {
	a.Scopes = s
}

func (a *Request) GetGrantedScopes() Arguments {
	return a.GrantedScopes
}

func (a *Request) GrantScope(scope string) {
	a.GrantedScopes = append(a.GrantedScopes, scope)
}

func (a *Request) SetSession(session interface{}) {
	a.Session = session
}

func (a *Request) GetSession() interface{} {
	return a.Session
}

func (a *Request) Merge(request Requester) {
	for _, scope := range request.GetScopes() {
		a.Scopes = append(a.Scopes, scope)
	}
	for _, scope := range request.GetGrantedScopes() {
		a.GrantedScopes = append(a.GrantedScopes, scope)
	}
	a.RequestedAt = request.GetRequestedAt()
	a.Client = request.GetClient()
	a.Session = request.GetSession()

	for k, v := range request.GetRequestForm() {
		a.Form[k] = v
	}
}
