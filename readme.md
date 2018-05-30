# Printfcoder GoUtils

<a href="https://godoc.org/github.com/printfcoder/goutils"><img src="https://godoc.org/github.com/printfcoder/goutils?status.svg" alt="GoDoc"></a>

Printfcoder GoUtils, a package of Golang common utility. Its inspiration comes from [apache commons lang][apache_commons_lang]

## Documentation

More information can be found on [godoc][go_doc] which can be browsed.

## How to use

```shell
$go get -u github.com/printfcoder/goutils/...

```

### example

```golang
import (
  "fmt"

  "github.com/printfcoder/goutils/stringutils"
)

func demo(){
    out := stringutils.IndexOf("0123456789", "2")
    fmt.Println(out)
}
```

## Contributing

Just [email me][email]

## License

Code is under the [Apache Licence v2](https://www.apache.org/licenses/LICENSE-2.0.txt).

[apache_commons_lang]: https://github.com/apache/commons-lang
[go_doc]: https://godoc.org/github.com/printfcoder/goutils
[email]: mailto:printfcoder@gmail.com
