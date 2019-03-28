# Uiza

## Welcome to your new GoLang!

...

## Installation
Install api-wrapper-go with:
```golang
go get -u github.com/uizaio/api-wrapper-go

```

### Requirements

- GoLang 10.0+.

## Usage

The library needs to be configured with your account's `WorkspaceAPIDomain` and `Key` (API key).
See details [here](https://docs.uiza.io/#authentication).
Set `Uiza.AppId` and `Uiza.Authorization` with your values:

## Init

```golang

import (
    Uiza "github.com/uizaio/api-wrapper-go"
)

func init() {
  Uiza.AppID = "your-AppId"
  Uiza.Authorization = "your-API-key"
}

```

## Custom Backend Config

You can custom your Backend Config by:

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
)

func init() {
  Uiza.AppID = "your-AppId"
  Uiza.Authorization = "your-API-key"

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
		&Uiza.BackendConfig{
			HTTPClient: httpClient,   // Your HttpClient.
			Logger:   Uiza.Logger,    // Log by your logger
			LogLevel: 3,              // Level debug log
		},
	)
	Uiza.SetBackend(Uiza.APIBackend, uizaBackendConfig)
}
```

## Entity

These below APIs used to take action with your media files (we called Entity)
See details [here](https://docs.uiza.io/?go#video).

## Category

Category has been splits into 3 types: folder, playlist and tag. These will make the management of entity more easier.
See details [here](https://docs.uiza.io/?go#category).

## Storage

You can add your storage (FTP, AWS S3) with UIZA. After synced, you can select your content easier from your storage to create entity.
See details [here](https://docs.uiza.io/?go#storage).

## Live

These APIs used to create and manage live streaming event.
See details [here](https://docs.uiza.io/?go#live-streaming).

## Callback

Callback used to retrieve an information for Uiza to your server, so you can have a trigger notice about an entity is upload completed and .
See details [here](https://docs.uiza.io/?go#callback).

## Development

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/uizaio/api-wrapper-go. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the Uiza project�s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/uizaio/api-wrapper-go/blob/master/CODE_OF_CONDUCT.md).
