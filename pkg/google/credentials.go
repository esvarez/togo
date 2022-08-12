package google

type Credentials struct {
	Installed `json:"installed"`
}

type Installed struct {
	ClientID                string   `json:"client_id"`
	ProjectID               string   `json:"project_id"`
	AuthURI                 string   `json:"auth_uri"`
	TokenURI                string   `json:"token_uri"`
	AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
	ClientSecret            string   `json:"client_secret"`
	RedirectURIs            []string `json:"redirect_uris"`
}

func getCredentials() *Credentials {
	return &Credentials{
		Installed{
			ClientID:                "11281438288-0ks28e321o628squ5bcn7che5ddqt84l.apps.googleusercontent.com",
			ProjectID:               "togo-cli",
			AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
			TokenURI:                "https://oauth2.googleapis.com/token",
			AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
			ClientSecret:            "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",
			RedirectURIs:            []string{"http://localhost"},
		},
	}
}
