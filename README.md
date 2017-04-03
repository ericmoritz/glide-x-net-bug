When I try to use glide, I get a different version of x/net/context
than what I get when I use `go get`.

``` shell
GCI-EMORITZ2-M:glide-x-net-bug emoritz$ go version
go version go1.8 darwin/amd64
GCI-EMORITZ2-M:glide-x-net-bug emoritz$ glide install
[INFO]	Lock file (glide.lock) does not exist. Performing update.
[INFO]	Downloading dependencies. Please wait...
[INFO]	No references set.
[INFO]	Resolving imports
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for golang.org/x/net
[INFO]	Setting references for remaining imports
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting golang.org/x/net
[INFO]	Replacing existing vendor dependencies
[INFO]	Project relies on 1 dependencies.
GCI-EMORITZ2-M:glide-x-net-bug emoritz$ go run main.go
# github.com/GannettDigital/glide-x-net-bug/vendor/golang.org/x/net/context
vendor/golang.org/x/net/context/go18.go:17: syntax error: unexpected = in type declaration
vendor/golang.org/x/net/context/go18.go:22: syntax error: unexpected = in type declaration
GCI-EMORITZ2-M:glide-x-net-bug emoritz$
```

With `go get`

``` shell
GCI-EMORITZ2-M:glide-x-net-bug emoritz$ rm -rf vendor/
GCI-EMORITZ2-M:glide-x-net-bug emoritz$ go get golang.org/x/net/context
GCI-EMORITZ2-M:glide-x-net-bug emoritz$ go run main.go
bg = context.Background
```

Directory listing of vendor/golang.org/x/net/context:

``` shell
GCI-EMORITZ2-M:glide-x-net-bug emoritz$ ls vendor/golang.org/x/net/context/
context.go          context_test.go     ctxhttp             go17.go             go18.go             pre_go17.go         pre_go18.go         withtimeout_test.go
```

Directory listing $GOPATH/src/goland.org/x/net/context:

``` shell
GCI-EMORITZ2-M:glide-x-net-bug emoritz$ ls $GOPATH/src/golang.org/x/net/context/
context.go          context_test.go     ctxhttp             go17.go             pre_go17.go         withtimeout_test.go
```
