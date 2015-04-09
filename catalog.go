package main

// taken from https://github.com/cloudfoundry-community/types-cf
// had to convert pointer refs in struct to values
// investigate using the package

type Catalog struct {
	Services []Service `json:"services"`
}

// Service describes Cloud Foundry service
type Service struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Bindable    bool        `json:"bindable"`
	Tags        []string    `json:"tags,omitempty"`
	Metadata    ServiceMeta `json:"metadata,omitempty"`
	Plans       []Plan      `json:"plans"`
	// Requires    []string    `json:"requires,omitempty"`
	// Dashboard   *Dashboard  `json:"dashboard_client,omitempty"`
}

// ServiceMeta describers Cloud Foundry service meta-data
type ServiceMeta struct {
	DisplayName         string `json:"displayName"`
	ImageURL            string `json:"imageUrl"`
	Description         string `json:"longDescription"`
	ProviderDisplayName string `json:"providerDisplayName"`
	DocURL              string `json:"documentationUrl"`
	SupportURL          string `json:"supportUrl"`
}

// Plan describes Cloud Foundry plan structure
type Plan struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Metadata    PlanMeta `json:"metadata,omitempty"`
	Free        bool     `json:"free,omitempty"`
}

// PlanMeta describers Cloud Foundry plan meta-data
type PlanMeta struct {
	Bullets     []string `json:"bullets"`
	Costs       string   `json:"costs"`
	DisplayName string   `json:"displayName"`
}

// Cost describers Cloud Foundry plan Cost
type Cost struct {
	Amount Amount `json:"amount"`
	Unit   string `json:"unit"`
}

// Amount describers Cloud Foundry cost amount
type Amount struct {
	Usd float32 `json:"usd"`
}

// Dashboard describes Cloud Foundry dashboard
type Dashboard struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	URI    string `json:"redirect_uri"`
}
