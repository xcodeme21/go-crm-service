package models

import (
	"database/sql"
	"time"
)

// DashboardStorage :nodoc
type DashboardStorage struct {
	ID               int          `json:"id"`
	URL              string       `json:"url"`
	Service          string       `json:"service"`
	Source           string       `json:"source"`
	Category         string       `json:"category"`
	OriginalFileName string       `json:"original_file_name"`
	StorageFileName  string       `json:"storage_file_name"`
	FileType         string       `json:"file_type"`
	FileExtension    string       `json:"file_extension"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DeletedAt        sql.NullTime `json:"deleted_at"`
}

type GoogleCloudCredential struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

type FilterDashboardStorage struct {
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	SortBy   string `json:"sort_by"`
	SortDir  string `json:"sort_dir"`
	Search   string `json:"search"`
	Source   string `json:"source"`
	Category string `json:"category"`
}
