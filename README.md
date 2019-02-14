# Uiza

## Welcome to your new GoLang!

...

## Installation

### Requirements

- GoLang 10.0+.

## Usage

The library needs to be configured with your account's `WorkspaceAPIDomain` and `Key` (API key).
See details [here](https://docs.uiza.io/#authentication).
Set `Uiza.WorkspaceAPIDomain` and `Uiza.Key` with your values:

## Init

```golang

import (
    Uiza "api-wrapper-go"
)

func init() {
  Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
  Uiza.Key = "your-API-key"
}

```

## Custom Backend Config

You can custom your Backend Config by:

```golang
import (
    Uiza "api-wrapper-go"
)

func init() {
  Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
  Uiza.Key = "your-API-key"

  // custom your transport of httpClient
  transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	uizaBackendConfig := Uiza.GetBackendWithConfig(
		Uiza.APIBackend,
		&uiza.BackendConfig{
			HTTPClient: httpClient,   // Your HttpClient.
			Logger:   uiza.Logger,    // Log by your logger
			LogLevel: 3,              // Level debug log
		},
	)
	Uiza.SetBackend(Uiza.APIBackend, uizaBackendConfig)
}
```

## Entity

These below APIs used to take action with your media files (we called Entity)
See details [here](https://docs.uiza.io/#video).

## Create entity

Create entity using full URL. Direct HTTP, FTP or AWS S3 link are acceptable.
See details [here](https://docs.uiza.io/#create-entity).

```golang

params = {
  name: "Sample Video",
  url: "https://example.com/video.mp4",
  inputType: "http",
  description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry"
}

import (
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

var typeHTTP = Uiza.InputTypeHTTP
params =  &uiza.EntityCreateParams{
					Name:      uiza.String("Sample Video"),
					URL:       uiza.String("https://example.com/video.mp4"),
          InputType: &typeHTTP,
          description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry"
        }

response, _ := entity.Retrieve(params)
log.Printf("%s\n", response)
```

## Retrieve entity

Get detail of entity including all information of entity.
See details [here](https://docs.uiza.io/#retrieve-an-entity).

```golang
import (
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

params := &Uiza.EntityRetrieveParams{ID: uiza.String("Your entity ID")}
response, _ := entity.Retrieve(params)
log.Printf("%s\n", response)
```

## Development

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/uizaio/api-wrapper-go. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the Uiza project�s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/uizaio/api-wrapper-go/blob/master/CODE_OF_CONDUCT.md).
