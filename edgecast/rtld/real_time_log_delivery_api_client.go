// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package rtld

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/url"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"

	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/lookups"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_cdn"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_rl"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/settings_internal"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/rtld"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

//New creates a new real time log delivery API client
func New(config edgecast.SDKConfig) (*RtldService, error) {

	apiURL, err := url.Parse(config.BaseAPIURL.String() + DefaultBasePath)
	if err != nil {
		return nil, fmt.Errorf("RtldService.New(): %v", err)
	}

	// OAuth2 authentication
	authProvider, err := ecauth.NewIDSAuthorizationProvider(config.BaseIDSURL, ecauth.OAuth2Credentials(config.IDSCredentials))
	if err != nil {

		//Token authentication
		authTokenProvider, err := ecauth.NewTokenAuthorizationProvider(config.APIToken)
		if err != nil {
			return nil, fmt.Errorf("RtldService.New(): %v", err)
		}
		c := ecclient.New(ecclient.ClientConfig{
			BaseAPIURL:   *apiURL,
			UserAgent:    config.UserAgent,
			Logger:       config.Logger,
			AuthProvider: authTokenProvider,
		})

		return &RtldService{
			client:           c,
			Logger:           config.Logger,
			Lookups:          lookups.New(c, c.Config.BaseAPIURL.String()),
			ProfilesCdn:      profiles_cdn.New(c, c.Config.BaseAPIURL.String()),
			ProfilesRl:       profiles_rl.New(c, c.Config.BaseAPIURL.String()),
			ProfilesWaf:      profiles_waf.New(c, c.Config.BaseAPIURL.String()),
			SettingsInternal: settings_internal.New(c, c.Config.BaseAPIURL.String()),
		}, nil

	} else {

		c := ecclient.New(ecclient.ClientConfig{
			BaseAPIURL:   *apiURL,
			UserAgent:    config.UserAgent,
			Logger:       config.Logger,
			AuthProvider: authProvider,
		})

		return &RtldService{
			client:           c,
			Logger:           config.Logger,
			Lookups:          lookups.New(c, c.Config.BaseAPIURL.String()),
			ProfilesCdn:      profiles_cdn.New(c, c.Config.BaseAPIURL.String()),
			ProfilesRl:       profiles_rl.New(c, c.Config.BaseAPIURL.String()),
			ProfilesWaf:      profiles_waf.New(c, c.Config.BaseAPIURL.String()),
			SettingsInternal: settings_internal.New(c, c.Config.BaseAPIURL.String()),
		}, nil
	}
}

// RtldService is a client for real time log delivery API
type RtldService struct {
	Lookups lookups.ClientService

	ProfilesCdn profiles_cdn.ClientService

	ProfilesRl profiles_rl.ClientService

	ProfilesWaf profiles_waf.ClientService

	SettingsInternal settings_internal.ClientService

	client ecclient.APIClient

	clientConfig ecclient.ClientConfig

	Logger eclog.Logger
}
