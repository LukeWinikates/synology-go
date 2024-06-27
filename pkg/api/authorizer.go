package api

import "net/url"

type Authorizer interface {
	Apply(values url.Values)
}

type NoopAuthorizer struct {
}

func (noop NoopAuthorizer) Apply(_ url.Values) {

}
