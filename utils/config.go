package helper

import (
	"github.com/xcodeme21/go-crm-service/models"
	"os"
)

func GetGoogleCloudStorageCredentials() models.GoogleCloudCredential {
	var credentials models.GoogleCloudCredential

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
	privateKeyID := os.Getenv("GOOGLE_CLOUD_PRIVATE_KEY_ID")
	privateKey := os.Getenv("GOOGLE_CLOUD_PRIVATE_KEY")
	clientEmail := os.Getenv("GOOGLE_CLOUD_CLIENT_EMAIL")
	clientID := os.Getenv("GOOGLE_CLOUD_CLIENT_ID")
	clientX509Cert := os.Getenv("GOOGLE_CLOUD_CLIENT_X509_CERT_URL")

	// CAN'T READ FILE WHEN BUILD ON PROD
	//storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("ERASPACE-a08de72906cf.json"))
	credentials.Type = "service_account"
	credentials.ProjectID = projectID
	credentials.PrivateKeyID = privateKeyID
	credentials.PrivateKey = privateKey
	credentials.ClientEmail = clientEmail
	credentials.ClientID = clientID
	credentials.AuthURI = "https://accounts.google.com/o/oauth2/auth"
	credentials.TokenURI = "https://oauth2.googleapis.com/token"
	credentials.AuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
	credentials.ClientX509CertURL = clientX509Cert
	return credentials
}

func EndpointURL(path string) string {
	// sample:
	// https://api-repair.eratech.id/path/to/your
	endpointURL := os.Getenv("ENDPOINT_URL") + path
	return endpointURL
}
