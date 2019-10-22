// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"fmt"
	"os"
)

// ClientCertificateCredential enables authentication of a service principal in to Azure Active Directory using a X509 certificate that is assigned to it's App Registration. More information
// on how to configure certificate authentication can be found here:
// https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-certificate-credentials#register-your-certificate-with-azure-ad
type ClientCertificateCredential struct {
	TokenCredential

	client aadIdentityClient

	// The Azure Active Directory tenant (directory) Id of the service principal
	tenantID string

	// The client (application) ID of the service principal
	clientID string

	// Path to the client certificate generated for the App Registration used to authenticate the client
	// CP: Across languages it is accepted to take the path to a non-password protected PEM file, can also take it from an env file
	clientCertificate string
}

// NewClientCertificateCredential creates an instance of ClientCertificateCredential with the details needed to authenticate against Azure Active Directory with the specified certificate.
// tenantID: The Azure Active Directory tenant (directory) Id of the service principal.
// clientID: The client (application) ID of the service principal.
// clientCertificate: The path to the client certificate that was generated for the App Registration used to authenticate the client.
// options: allow to configure the management of the requests sent to the Azure Active Directory service.
func NewClientCertificateCredential(tenantID string, clientID string, clientCertificate string, options *IdentityClientOptions) (ClientCertificateCredential, error) {
	// CP: Do this or pass *os.File or io.Reader? Could also make a struct to pass io.Reader + file size for fingerprint and private key funcs
	_, err := os.Stat(clientCertificate)
	if err != nil {
		return ClientCertificateCredential{}, fmt.Errorf("Certificate file not found in path: %s", clientCertificate)
	}

	if options == nil {
		options, err = newIdentityClientOptions()
		if err != nil {
			return ClientCertificateCredential{}, fmt.Errorf("newIdentityClientOptions: %w", err)
		}
	}
	var client aadIdentityClient = newAADIdentityClient(options)

	return ClientCertificateCredential{tenantID: tenantID, clientID: clientID, clientCertificate: clientCertificate, client: client}, nil
}

// GetToken obtains a token from the Azure Active Directory service, using the certificate in the file path to authenticate.
// scopes: The list of scopes for which the token will have access.
// ctx: controlling the request lifetime.
// Returns an AccessToken which can be used to authenticate service client calls.
func (c ClientCertificateCredential) GetToken(ctx context.Context, scopes []string) (*AccessToken, error) {
	return c.client.AuthenticateCertificate(ctx, c.tenantID, c.clientID, c.clientCertificate, scopes)
}
