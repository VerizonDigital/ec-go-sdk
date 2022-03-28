# Go SDK for EdgeCast CDN
This is the official Go SDK for EdgeCast. It contains services that interact 
with the EdgeCast APIs.

## Dependencies
- Go 1.16

## Installing
### $GOPATH
To install the SDK into your $GOPATH:
```
go get -u github.com/EdgeCast/ec-sdk-go
```

### Go Modules
```
go get github.com/EdgeCast/ec-sdk-go
```

## Using the SDK
Simply import the SDK and provide the API credentials provided to you. They may 
be an API Token or OAuth 2.0 Credentials. Examples are listed below for each 
feature.

### Customer management ###

Our Customer Management service provides administrative operations to manage 
Customer accounts and Customer User accounts for a Partner. These operations 
allows a partner to automate administrative tasks on their customers and 
customer user accounts. 

To use this Customer Management service, use the API Token provided to you.

#### Customer

Customer Account Management Operations allows management of customer accounts

```
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/customer"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	customerService, err := customer.New(sdkConfig)
	newCustomer := customer.Customer{
		// ...
	}
	addParams := customer.NewAddCustomerParams()
	addParams.Customer = newCustomer
	resp, err := customerService.AddCustomer(*addParams)
	// ...
}
```

#### Customer User

Customer User Account Management Operations allows management of user accounts 
under a (parent) customer.

```
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/customer"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	customerService, err := customer.New(sdkConfig)
	customerUser := customer.CustomerUser{
		// ...
	}
	addParams := customer.NewAddCustomerUserParams()
	addParams.CustomerUser = customerUser
	// ...
	resp, err := customerService.AddCustomerUser(*addParams)
	// ...
}
```

### Edge CNAME ###

Our User-Friendly URL, also known as Edge CNAME, takes advantage of an Edge CNAME 
configuration and a CNAME record to provide a friendlier alternative to a CDN URL. 
An edge CNAME URL is specific to the platform from which it was configured.

For more information, please read the [official documentation](https://docs.edgecast.com/cdn/index.html#Origin_Server_-_File_Storage/Creating_an_Alias_for_a_CDN_URL.htm)

To use the edge CNAME service, use the API Token providede to you.

```
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/edgecname"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	edgecnameService, err := edgecname.New(sdkConfig)
	cname := edgecname.EdgeCname{
		// ...
	}
	addParams := edgecname.NewAddEdgeCnameParams()
	addParams.EdgeCname = cname
	// ...
	resp, err := edgecnameService.AddEdgeCname(*addParams)
// ...
}
```

### Customer Origins 

Our Customer Origin Service allows to serve content stored or generated by 
third-party web servers (e.g., web hosting) via the CDN by:

- Creating a customer origin configuration. This type of configuration maps one 
or more servers to a CDN URL.
- (Optional) Creating an edge CNAME configuration that allows you to serve traffic 
via the CDN without having to update your links. This type of configuration maps 
a customer origin configuration to a CNAME record.

For more information, please read the [official documentation](https://docs.edgecast.com/cdn/index.html#Origin_Server_-_File_Storage/Customer_Origin_Server.htm)

To use the Customer Origins service, use the API Token provided to you.

```
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/edgecname"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	originService, err := origin.New(sdkConfig)
	newOrigin := origin.Origin{
		// ...
	}
	addParams := origin.NewAddOriginParams()
	// ...
	addParams.Origin = newOrigin
	resp, err := originService.AddOrigin(*addParams)
// ...
}
```

### Route DNS - BETA ###

Our Route (DNS) solution is a reliable, high performance, and secure DNS service 
that provides capabilities such as:

- Load balance traffic for a CNAME record or a subdomain for a primary zone hosted 
on another DNS system.
- Establish a failover system for a CNAME record or a subdomain for a primary zone 
hosted on another DNS system.
- Create a standard DNS zone. Optionally, load balance or failover requests to 
that zone.
- Import a secondary DNS zone by creating a master server group and a secondary 
zone group.
- Verify a server's capability to fulfill requests through health checks performed 
from around the world.

For more information, please read the [official documentation](https://docs.whitecdn.com/cdn/index.html#Route/Route_DNS.htm%3FTocPath%3DRoute%2520(DNS)%7C_____0)

To use the Route(DNS) service, use the API Token provided to you.

#### Master Server Group

A master server group allows quick and easy management of master name servers, 
while a secondary zone group defines the secondary zones that will be imported 
from servers defined in a master server group and any TSIG keys that should be 
used for the zone transfer.

```
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/routedns"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	routeDNSService, err := routedns.New(sdkConfig)
	masterServerGroup := buildMasterServerGroup()
	addParams := routedns.NewAddMasterServerGroupParams()
	addParams.MasterServerGroup = masterServerGroup
	// ...
	resp, err := routeDNSService.AddMasterServerGroup(*addParams)
// ...
}
```

### Real Time Log Delivery (RTLD)

Our Real-Time Log Delivery (RTLD) servicee delivers log data in near real-time 
to a variety of destinations.

For more information, please read the [official documentation](https://docs.edgecast.com/cdn/index.html#RTLD/RTLD.htm%3FTocPath%3DReports%2520and%2520Log%2520Data%7CReal-Time%2520Log%2520Delivery%7C_____0).

To use the Rules Engine service, use OAuth 2.0 Credentials.

#### Real-Time Log Delivery Rate Limiting (RTLD Rate Limiting)

Delivers log data that describes requests for which Web Application Firewall (WAF) 
enforced a rate limit as defined through a rate rule.

```
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = "IDS CREDENTIALS"
	rtldService, err := rtld.New(sdkConfig)
	addParams := profiles_rl.NewProfilesRateLimitingAddCustomerSettingParams()
	addParams.SettingDto = &rtldmodels.RateLimitingProfileDto{
		// ...
	}
	addResp, err :=
		rtldService.ProfilesRl.ProfilesRateLimitingAddCustomerSetting(addParams)
// ...
}
```

### Rules Engine

Our Rules Engine allows the customization of requests handled by our CDN. 
Sample customizations that may be performed are:

Override or define a custom cache policy
Secure or deny requests for sensitive content
Redirect requests to a different URL

For more information, please read the [official documentation](https://docs.whitecdn.com/cdn/index.html#HRE/HRE.htm).

To use the Rules Engine service, use OAuth 2.0 Credentials.
A Policy should be constructed as a JSON object passed as a string.

```
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)
// ...
	policyString := `{
		// ...
	}`
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = "IDS CREDENTIALS"
	rulesengineService, err := rulesengine.New(sdkConfig)
	addParams := rulesengine.NewAddPolicyParams()
	addParams.PolicyAsString = policyString
	addPolicyResp, err := rulesengineService.AddPolicy(*addParams)
// ...
}
```

### Web Application Firewall (WAF)
Our WAF service provides a layer of security between many security threats and 
your external web infrastructure. WAF increases security by monitoring, detecting, 
and preventing application layer attacks. It inspects inbound HTTP/HTTPS traffic 
against reactive and proactive security policies and blocks malicious activity 
in-band and on a real-time basis.

For more information, please read the [official documentation](https://docs.edgecast.com/cdn/index.html#Web-Security/Web-Security.htm%3FTocPath%3DSecurity%7CWeb%2520Application%2520Firewall%2520(WAF)%7C_____0).

To use the WAF service, use the API Token provided to you.

#### Access Rules
An access rule identifies legitimate traffic and threats by ASN, Cookie, Country, 
IP Address, Referrer, URL, User agent, HTTP method, Media type, File extension, 
and Request headers.
```
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	wafService, err := waf.New(sdkConfig)
	rule := waf.AccessRule{
		// ... 
	}
	resp, err := wafService.AddAccessRule(rule)
// ...
}
```

### Structure

```
.
├── edgecast
	package containing the main functionality for sdk.
	Please add new client and model folders for new services here.
│   ├── eclog
		defines the the implementation and helper methods for logging
│   ├──	internal
		package containing helper methods and shared functionality used in sdk
		please add new helper methods here
│   │	├── collectionhelper
			helper methods for working with aggregate/collection types
│   │	├── jsonhelper		
			helper methods for working with json
│   │	├── testhelper
			helper methods used in testing
│   │	├── ecauth
			authentication layer for oauth 2.0 and token based authentication
│   │	├── ecclient
			package client provides a base client implementation for interacting 
			with edgecast cdn apis.
			configuration and authentication types are also provided.
│   ├── customer
		client files for interacting with customer api
		model files for customer
│   ├── edgecnamee
		client files for interacting with edge cname api
		model files for edge cnamee
│   ├── origin
		client files for interacting with customer origin api
		model files for customer origin
│   ├── routedns
		client files for interacting with route (dns)
		model files for route (dns)
│   ├── rtld
		client files for interacting with real time log delivery api
│   ├── rtldmodels
		model files for real time log delivery
│   ├── rulesengine
		client files for interacting with rules engine api
		model files for rules engine
│   ├── waf
		client files for interacting with waf api
		model files for waf
│   ├── shared
		shared models and enums
│   ├── config
		defines the configuration of sdk services
│   ├── doc
		please add new docs here as needed
│   ├── version
		lists the latest version of sdk		
├── example
	example files to get started using the services
├── template
	template files used to generate client files and models using swagger api 
	documentation
├── Makefile
        This Makefile should contain all testing and building operations.

```

### Contributing

Please refer to the [contributing.md] (https://github.com/EdgeCast/ec-sdk-go/blob/main/Contributing.md) 
file for information about how to get involved. 
We welcome issues, questions, and pull requests.

### Maintainers

- Frank Contreras: frank.contreras@edgecast.com
- Hector Gray: hector.gray@edgecast.com
- Steven Paz: steven.paz@edgecast.com
- Shikha Saluja: shikha.saluja@edgecast.com

### Testing

To run all test files in the root folder
```
go test -v ./...
```
Tests should all pass, before and after any work that you do.  If they
do not, please reach out to the maintainers for help.

Separately, all test files are also run when a pull request is created.

## Resources
[CDN Reference Documentation](https://docs.edgecast.com/cdn/index.html) - This 
is a useful resource for learning about EdgeCast CDN. It is a good starting point 
before using this SDK.

[API Documentation](https://docs.edgecast.com/cdn/index.html#REST-API.htm%3FTocPath%3D_____8) - 
For developers that want to interact directly with the EdgeCast CDN API, 
refer to this documentation. It contains all of the available operations as well 
as their inputs and outputs.

[Examples](https://github.com/EdgeCast/ec-sdk-go/tree/main/example) - Examples 
to get started can be found here.

[Submit an Issue](https://github.com/EdgeCast/ec-sdk-go/issues) - Found a bug? 
Want to request a feature? Please do so here.