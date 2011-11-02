include ${GOROOT}/src/Make.inc

TARG = fuzz
GOFILES = fuzz.go parse.go format.go

include ${GOROOT}/src/Make.pkg

fmt:
	gofmt -w *.go

