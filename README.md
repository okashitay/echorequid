[![Coverage Status](https://coveralls.io/repos/github/deadcheat/echorequid/badge.svg?branch=master)](https://coveralls.io/github/deadcheat/echorequid?branch=master)

# Echo-Requid

requid is provide as middleware of echo, which adding "X-Request-Id" to HTTP-Header.

## Usage

install like below.

```
go get github.com/deadcheat/echorequid
```

and you only just use it!

```
	// prepare echo
	e := echo.New()
	e.Pre(echorequid.AppendRequestUUID())
```

