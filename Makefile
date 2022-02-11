tagregex="v[0-9]*.[0-9]*.[0-9]*"
version=`git describe --tags --dirty --always --long --match=${tagregex} 2>/dev/null`

default:
	go build -o mcrepo main.go

fmt:
	@gofmt -s -w .

clean:
	rm -f mcrepo

.PHONY: default fmt clean
