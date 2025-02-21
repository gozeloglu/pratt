# pratt

The `pratt` command is a stripped-down expression parser based on the
[Pratt parser]. To help understand Pratt parsing, the command also
provides a naive, recursive, right-associative expression parser and a naive,
iterative, left-associative expression parser.

Try it with:

```
go run ./main.go '1 + 2 * 3'
```


[Pratt parser]: https://en.wikipedia.org/wiki/Pratt_parser

## `svg` command

The `svg` command is meant to be used with the `pratt` command. It generates
and opens an SVG image of the expression tree in the default SVG viewer.

Try it with:

```
go install ./...
pratt '1 * 2 + 3 * 4' | svg -o
```

## Installation

You can install by using the following command.

```shell
go install github.com/gozeloglu/pratt@latest
```

## Credit

Thanks to [@juliaogris](https://github.com/juliaogris) for this tool. Also, thanks to her for letting me improve new features.

## LICENSE

[MIT](LICENSE)