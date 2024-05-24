package projects

import "net/url"

type ServicePortalConfiguration struct {
	Enabled  bool
	Name     string
	Port     int
	Protocol string
}

func applyServicePortalOptions(_ *ServicePortalConfiguration, query url.Values) {
	query.Add("enable_service_portal", "false")
	query.Add("service_portal_name", `""`)
	query.Add("service_portal_port", "0")
	query.Add("service_portal_protocol", `"0"`)
}
