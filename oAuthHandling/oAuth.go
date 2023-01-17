package oauthhandling

type GoogleAuth struct {
	client_id string `json:"client_id"`
	scope string `json:"client_secret"`
	redirect_uri string `json:"redirect_uri"`
	response_type string `json:"response_type"`
	state string `json:"state"`
}