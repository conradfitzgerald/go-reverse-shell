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

### Does it work on Windows?
This thing will compile/run on Windows. But God forbid it'll actually work.
Something to do with weird `bufio.NewReader(os.Stdin)` input streams and `\r\n` newlines over regular `\n`.
If someone wants to help make a Windows branch, feel free to do so.

### Installation/Usage:
`go run main.go` OR `go build main.go`

---

## TODO:
```
[X] -- Change commandline struct names
[X] -- Implement cat()
[X] -- Implement help()
[X] -- Implement exec()
[ ] -- Implement rm/rf()
[ ] -- Enumeration scripts?
[ ] -- Implement download()
[ ] -- Implement upload()
[ ] -- Implement remoteDownload()
[ ] -- Set up server/client
[X] -- Windows Testing
[ ] -- Windows working
[ ] -- Verify Build Requirement versions
[ ] -- Errors should not kill the shell
[ ] -- Create header for shell functions
[ ] -- Create header for struct and program functions
```

---

Made with <3.
