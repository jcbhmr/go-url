# URL Standard for Go

💻 WHATWG URL Standard bindings for Go

<table align=center><td>

```go
x := url.NewURL("https://example.org/store/search?q=cool#results", nil)
fmt.Printf("origin=%s\n", x.Origin())
q, _ := x.SearchParams().Get("q")
fmt.Printf("q=%s\n", q)
fmt.Printf("pathname=%s\n", x.Pathname())
url.SetPathname("/menu/search")
fmt.Printf("href=%s\n", x.Href())
// Output:
// origin=https://example.org
// q=cool
// pathname=/store/search
// href=https://example.org/menu/search?q=cool#results
```

</table>

~~🌐 Has complete Go (no Cgo) implementation of the URL Standard~~<sup>not yet</sup> \
🟨 Uses JavaScript `URL` and `URLSearchParams` implementations when compiling for `js/wasm` \
✅ Preserves all the Web IDL idioms in Go \
📄 Interacts well with other web-related Go bindings

## Installation

```sh
go get github.com/jcbhmr/go-url
```

⚠️ This package currently only works with `js/wasm`. 

## Usage

```go
package main

import "github.com/jcbhmr/go-url"

func main() {
  x := url.NewURL("https://example.org/", nil)
}
```

<!-- ✅ These bindings conform to the Go ↔ Web IDL bindings recommendations outlined at [jcbhmr.me/WebIDL-Go](https://jcbhmr.me/WebIDL-Go/). -->

## Development

TODO
