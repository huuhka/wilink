package adotool

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"strings"
)

func NewOauthConnection(organizationUrl string, accessToken string) *azuredevops.Connection {
	authorizationString := "Bearer " + accessToken
	normalizedUrl := normalizeUrl(organizationUrl)

	return &azuredevops.Connection{
		AuthorizationString:     authorizationString,
		BaseUrl:                 normalizedUrl,
		SuppressFedAuthRedirect: true,
	}
}

func normalizeUrl(url string) string {
	return strings.ToLower(strings.TrimRight(url, "/"))
}