[![Build Status](https://travis-ci.com/mekilis/purify.svg?branch=master)](https://travis-ci.com/mekilis/purify)

# Purify

This is a simple profanity filter API written in `Go`. The API automatically
removes content of profanity, obscenity or other undesirable content in a
given text and modifies it as needed. It has no authorization as yet.

Supports only English for now.

## Installation

`go get github.com/mekilis/purify/...`

Purify would automatically be installed at the already defined `$GOPATH` for the
current user. Otherwise use the `go` tool to build and run at any other desired
location.

### Usage

The service can be started via a terminal:

```bash
$ purify
```

or

```bash
$ /path/to/purify
```

depending on the mode of installation. By default, the server runs on port
`:9002`. This can be changed by passing the `-p` or `--port` argument, for
example:

```bash
$ purify --port 12345
```

Then a message is sent via `POST` method to the root endpoint
`http://localhost:9002`.

#### Example

Request:
```json
{
	"message": "sh!t"
}
```
Response:

```json
{
    "status_code": 1,
    "status": "successfully filtered",
    "message": "s**t"
}
```


## Supported Go Versions
Tested only on Go version 1.10 Linux/AMD64.

## License
Apache License 2.0
