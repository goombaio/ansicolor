# ansicolor

ANSI color string colorizer for Golang.

## Badges

[![License][License-Image]][License-URL]
[![CircleCI Status][CircleCI-Image]][CircleCI-URL]
[![Coverage Report][Coverage-Image]][Coverage-URL]
[![Go Report Card][GoReportCard-Image]][GoReportCard-URL]
[![CII Best Practices][CII-Image]][CII-URL]
[![GoDoc][GoDoc-Image]][GoDoc-URL]

## Install

```bash
go get github.com/goombaio/ansicolor
```

You can also update an already installed version:

```bash
go get -u github.com/goombaio/ansicolor
```

## Example of use

```go
package main

import (
    "github.com/goombaio/ansicolor"
)

func main() {
    str := "This is an string"

    result := str
    fmt.Println(result)

    result = ansicolor.Color8Colors(str, 31, 40)
    fmt.Println(result)

    result = ansicolor.Color16Colors(str, 93, 100)
    fmt.Println(result)

    result = ansicolor.Color256Colors(str, 4, 238)
    fmt.Println(result)

    result = ansicolor.ColorTrueColors(str, 76, 168, 67, 55, 55, 55)
    fmt.Println(result)
}
```

## License

Copyright (c) 2018 Goomba project Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

[License-Image]: https://img.shields.io/badge/License-Apache-blue.svg
[License-URL]: http://opensource.org/licenses/Apache
[CircleCI-Image]: https://circleci.com/gh/goombaio/ansicolor.svg?style=svg
[CircleCI-URL]: https://circleci.com/gh/goombaio/ansicolor
[Coverage-Image]: https://codecov.io/gh/goombaio/ansicolor/branch/master/graph/badge.svg
[Coverage-URL]: https://codecov.io/gh/goombaio/ansicolor
[GoReportCard-Image]: https://goreportcard.com/badge/github.com/goombaio/ansicolor
[GoReportCard-URL]: https://goreportcard.com/report/github.com/goombaio/ansicolor
[CII-Image]: https://bestpractices.coreinfrastructure.org/projects/2234/badge
[CII-URL]: https://bestpractices.coreinfrastructure.org/projects/2234
[GoDoc-Image]: https://godoc.org/github.com/goombaio/ansicolor?status.svg
[GoDoc-URL]: http://godoc.org/github.com/goombaio/ansicolor