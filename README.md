# status-go

When you call to `status-go`, it'll check response codes for 5 minutes in a tight loop. There's some simple flood
protection to make sure it doesn't hit the webserver more than once a second.

If you provide an argument to `status-go`, it'll assume that's the website you're testing. If you don't provide an
argument, it defaults to https://gitlab.com.

```
status-go https://matthew.nz
=> 200 22ms
=> 200 22ms
... 5 minutes pass
=> 200 21ms
=> 200 21ms

status-go
200 507ms
200 515ms
... 5 minutes pass
200 501ms
200 507ms
```

## Installation

Make sure you have go installed, and that you're exporting $GOPATH/bin into your path. If you've got these, just:

```
go get -u github.com/heymatthew/status-go
```
