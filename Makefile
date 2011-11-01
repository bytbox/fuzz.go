include ${GOROOT}/src/Make.inc

TARG = fuzz
GOFILES = fuzz.go parse.go

include ${GOROOT}/src/Make.pkg

