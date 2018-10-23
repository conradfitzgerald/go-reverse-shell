# go-reverse-shell
A reverse shell implementation in Go.

---

This is a little project I've decided to start working on. It's an attempt to build a reverse shell in Go, using UDP for C&C.

Build requirements:

```
Go toolchain
Packages {
  bufio
  io/ioutil
  fmt
  strings
  os
}
A Linux OS
```

### Installation/Usage:
`go run main.go` OR `go build main.go`

## TODO:
```
-- Change commandline struct names
-- Implement cat()
-- Implement help()
-- Implement exec()
-- Implement download()
-- Set up server/client
-- Windows Testing
-- Verify Build Requirement versions
```

---

Made with <3.
