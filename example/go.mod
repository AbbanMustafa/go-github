module github.com/google/go-github/v47/example

go 1.17

require (
	github.com/bradleyfalzon/ghinstallation/v2 v2.0.4
	github.com/google/go-github/v47 v47.0.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	google.golang.org/appengine v1.6.7
)

require (
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/google/go-github/v41 v41.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
)

// Use version at HEAD, not the latest published.
replace github.com/google/go-github/v47 => ../
