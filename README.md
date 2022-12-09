# Embedded development services HTTP client

[![Go Badge](https://img.shields.io/badge/go-v1.19.3-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Overview
 
This repository contains an HTTP client written in Golang for communicating with Arm Embedded development services.
It is actually generated from our Web APIs contract which follows [OpenAPI 3](https://swagger.io/specification/) using the [GO generator](https://github.com/OpenAPITools/openapi-generator)

*Maintainers:* @ARM-software/embedded-development-services-client-admin 
 
## Using this library

To use this library, add the following line to your `go.mod`:
```
require (
    github.com/ARM-software/embedded-development-services-client/utils latest
    ...
)
```


## Releases

For release notes and a history of changes of all **production** releases, please see the following:

- [Changelog](CHANGELOG.md)

## Project Structure

The follow described the major aspects of the project structure:

- `docs/` - Code reference documentation.
- `utils/` - Go project source files.
- `changes/` - Collection of news files for unreleased changes.


## Getting Help

- For interface definition and usage documentation, please see [GitHub Pages](https://arm-software.github.io/embedded-development-services-client).
- For a list of known issues and possible workarounds, please see [Known Issues](KNOWN_ISSUES.md).
- To raise a defect or enhancement please use [GitHub Issues](https://github.com/ARM-software/embedded-development-services-client/issues) or [GitHub discussions](https://github.com/ARM-software/embedded-development-services-client/discussions).

## Contributing

- We are committed to fostering a welcoming community, please see our
  [Code of Conduct](CODE_OF_CONDUCT.md) for more information.
- For ways to contribute to the project, please see the [Contributions Guidelines](CONTRIBUTING.md)
- For a technical introduction into developing this package, please see the [Development Guide](DEVELOPMENT.md)
