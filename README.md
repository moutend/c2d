# c2d

The `c2d` is a command-line tool designed for finding the description of a given character.

## Prerequisites

- Go 1.20+

## Install

```console
go install github.com/moutend/c2d/cmd/c2d@latest
```

## Usage

The `c2d` executable does not include dictionary files. Please run the `download` subcommand first:

```console
c2d download
```

Now, you can find the description of a character by executing the following command:

```console
c2d find --languages 'ja' 'あ'
```

## LICENSE

MIT
