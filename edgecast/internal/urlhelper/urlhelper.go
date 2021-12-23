// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

// Package urlhelper provides helper methods for working with URLs
package urlhelper

import "fmt"

// FormatURLAddPartnerID is a utility function for adding the optional partner ID query string param
func FormatURLAddPartnerID(originalURL string, partnerID int) string {
	if partnerID != 0 {
		return originalURL + fmt.Sprintf("&partnerid=%d", partnerID)
	}
	return originalURL
}
