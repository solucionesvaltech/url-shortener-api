# URL Shortener API

Project that serves as a reference to work with Hexagonal architecture.

### [wire](https://github.com/google/wire)

Install wire for dependency injection management:

```shell
go install github.com/google/wire/cmd/wire@latest
```

use this in script in the [dependencies](internal/dependency) directory:

```shell
wire
```

### [golangci-lint](https://golangci-lint.run/)

Please don't miss the opportunity to improve our code, install this linter before adding new commits

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