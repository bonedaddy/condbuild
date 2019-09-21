.PHONY: linux
linux:
	GOOS=linux go build cmd/condbuildlinux/main.go

.PHONY: darwin
darwin:
	GOOS=darwin go build cmd/condbuilddarwin/main.go