<img src="assets/logo.png" alt="logo" width="50"/>

# URL Shortener API


The URL Shortener API enables users to shorten long URLs, generating a unique short identifier and redirecting to the original URL. Key functionalities include:

- Create Short Endpoint: Generates a unique identifier for a long URL.
- Redirection: Allows access to the original URL using the short identifier.
- URL Details: Retrieves information about a specific URL, including its status and original link.
- Toggle Status: Enables or disables a short URL, giving control over its accessibility.

### Solution

#### Pattern

The Hexagonal Architecture in this project separates business logic from implementation details, creating a modular and maintainable system.
Here’s a concise overview of the directory structure and its purpose:
</br>

- cmd: Contains the entry point (main.go), responsible for initializing the application.
</br>

- internal

  adapter
    - driven: Adapters that implement interactions with databases, external services, etc., acting as dependencies for the core.
    - driver: Entry adapters, such as HTTP controllers, that interact with core use cases.


  core
  - domain: Contains core entities that represent business concepts.
  - port: Interfaces defining communication points between business logic and adapters.
  - usecase: Holds the main business logic and workflows of the system.

  dependency: Central configuration for dependency injection.

- pkg: Shared utilities like configuration, logger, web server, metrics, and error handling.

Advantages
This approach decouples business logic from technical details, ensuring extensibility, maintainability, and easy testing.
Hexagonal architecture supports adaptability by allowing changes in technology within adapters without affecting the core application.

#### Structure

```
├── Dockerfile
├── assets
│   └── docker
├── cmd
│   └── main.go
├── internal
│   ├── adapter
│   │   ├── driven
│   │   └── driver
│   ├── core
│   │   ├── domain
│   │   ├── port
│   │   └── usecase
│   └── dependency
│       └──dependency.go
└── pkg
    ├── config
    ├── customerror
    ├── helper
    ├── log
    ├── metric
    └── server
```

### Endpoints

You can import the endpoints from this [collection](assets/url-shortener.postman_collection.json) file
  
- Create URL

  Type: POST </br>
  Endpoint: urls </br>
  Description: Creates a new short URL from a provided long URL in the request body. </br>
  Request Body:
  ```json
  {
    "url": "https://example.com"
  }
  ```
  Response Code: 201 Created </br>
  Response: Returns the unique ID that represents the shortened URL. </br>
</br>

- Get URL Details

  Method: GET </br>
  Endpoint: urls/:shortID </br>
  Description: Returns details about a specific short URL, including its status (enabled/disabled) and the original URL. </br>
  Path Parameter: shortID (The unique identifier for the shortened URL). </br>
  Response Code: 200 OK or 404 Not Found if the shortID does not exist. </br>
</br>

- Redirect (Get Short URL)

  Method: GET </br>
  Endpoint: :shortID </br>
  Description: Redirects the user to the original URL associated with the specified shortID. </br>
  Path Parameter: shortID (The unique identifier for the shortened URL). </br>
  Response Code: 302 Moved Permanently (Redirects to the original URL) or 404 Not Found if the shortID does not exist. </br>
</br>

- Update URL
  
  Method: PUT
  Endpoint: :shortID </br>
  Description: Updates the original URL associated with an existing shortID. </br>
  Request Body: </br>
  ```json
  {
    "url": "https://example.com"
  }
  ```
  Path Parameter: shortID (The unique identifier for the shortened URL). </br>
  Response Code: 200 OK or 404 Not Found if the shortID does not exist.  </br>
</br>

- Toggle URL Status

  Method: PATCH  </br>
  Endpoint: :shortID  </br>
  Description: Changes the status of the short URL to either enabled or disabled.  </br>
  Request Body:  </br>
  ```json
  {
    "enabled": true
  }
  ```
  Path Parameter: shortID (The unique identifier for the shortened URL). </br>
  Response Code: 200 OK or 404 Not Found if the shortID does not exist. </br>
</br>

- Health Check

  Method: GET  </br>
  Endpoint: health  </br>
  Description: Checks the health status of the API.  </br>
  Response Code: 200 OK (if operational)  </br>
</br>

- Metrics

  Type: GET </br>
  Endpoint: metrics </br>
  Description: Exposes Prometheus metrics to monitor API performance and usage. </br>
  Response Code: 200 OK </br>
</br>

### Local environment

#### Services

You can use the required services using the [docker-compose](https://docs.docker.com/compose/intro/features-uses/#key-benefits-of-docker-compose) file

```shell
# cd assets/docker

docker-compose -p url-shortener-api up -d

```

#### Variables

We use [configuration](pkg/config/configyaml/config.yaml) in [YAML](https://yaml.org/) format.
This values are converted to a [Config](pkg/config/config.go) type.
This approach offers a controlled way to managed and validate all the variables before the app starts.


### Git Hooks

This project use some githooks to improve the code quality. Make sure to run below command to replace .git/hooks
by **.githooks/**

```shell
chmod ug+x .githooks/*
git config core.hooksPath .githooks
```

### [wire](https://github.com/google/wire)

Install wire for dependency injection management:

```shell
go install github.com/google/wire/cmd/wire@latest
```

use this script in the [dependencies](internal/dependency) directory:

```shell
wire
```

### [golangci-lint](https://golangci-lint.run/)

Please don't miss the opportunity to improve the code, install this linter before adding new commits

Homebrew:

```bash
brew install golangci-lint@1.61.0
```

Shell:

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0
```

Run:

```bash
golangci-lint run -v --fix
```


### [arch-go](https://github.com/fdaines/arch-go)

Install arch-go to run architecture unit test cases:

```shell
go install -v github.com/fdaines/arch-go@1.5.4
```

Then you can run below command to check the arch unit compliance:

```shell
arch-go
```

And off course you can contribute editing the [arch-go.yml](arch-go.yml) file

### Testing

To ensure robust and reliable code, the project uses [Ginkgo](https://github.com/onsi/ginkgo) for behavior-driven testing and [mockgen](https://github.com/golang/mock) for generating mocks. 

### [mockgen](https://github.com/golang/mock)

Install `mockgen` for mock generation:

```shell
go install github.com/golang/mock/mockgen@v1.6.0
```

Generate mocks the mockgen command, here there is an example:

```shell
mockgen -source=path/yourfile.go -destination=path/mock.go
```


### [Ginkgo](https://github.com/onsi/ginkgo)

Is a testing framework designed to write expressive specs.

To ensure that you will be able to run your tests follow this steps:

- System installation
  ```bash
  go install github.com/onsi/ginkgo/v2/ginkgo@latest
  ```
- Plugin installation
  [install the plugin](https://plugins.jetbrains.com/plugin/17554-ginkgo)
- Setup
  In the IDE setup configuration, add Ginko for all tests
  ![img.png](/assets/ginko-config.png)

Run all tests:
```bash
ginkgo -v ./...
```
Or if you preffer the simple approach
```bash
go test ./...
```

Made with 💛 by Víctor Valenzuela

