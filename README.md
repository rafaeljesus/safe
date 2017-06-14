## Safe

* Thread-safe for any value.

## Installation
```bash
go get -u github.com/rafaeljesus/safe
```

## Usage
```go
package main

import "github.com/rafaeljesus/safe"

func main() {
  data := map[string]interface{}{"key": "value"}
  s := New(data)
  result, ok := s.Get().(map[string]interface{})
  if !ok {
    // wrong type
  }
  // do something
}
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![Build Status](https://circleci.com/gh/rafaeljesus/safe.svg?style=svg)](https://circleci.com/gh/rafaeljesus/safe)
[![Go Report Card](https://goreportcard.com/badge/github.com/rafaeljesus/safe)](https://goreportcard.com/report/github.com/rafaeljesus/safe)
[![Go Doc](https://godoc.org/github.com/rafaeljesus/safe?status.svg)](https://godoc.org/github.com/rafaeljesus/safe)

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Medium [@_jesus_rafael](https://medium.com/@_jesus_rafael) &nbsp;&middot;&nbsp;
> Twitter [@_jesus_rafael](https://twitter.com/_jesus_rafael)
