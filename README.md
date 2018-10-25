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
  os/exec
  runtime
}
A Linux OS
```

### Does it work on Windows?
This thing will compile/run on Windows. But God forbid it'll actually work.
Something to do with weird `bufio.NewReader(os.Stdin)` input streams and `\r\n` newlines over regular `\n`.
If someone wants to help make a Windows branch, feel free to do so.

### Installation/Usage:
`go run main.go` OR `go build main.go`

### Warranty
Read the license. There is none. If something really breaks your computer and you can't unfuck it, good luck.
I wrote 80% of this shell while absolutely hammered. Read the code yourself and determine whether it's safe or not.

I didn't make this to break anyone's computer or as some weird prank. Everything built-in to this program is tested by me before it's pushed to the main branch. You should still be testing it in a safe environment before trying to unleash it into the wild.

### Issues
There's no warranty on this, again. But if something is getting weird behaviour and you can't figure out why, then open up an issue here. Don't email me. Don't dox me and contact me from there. Open the issue here, if I feel like it's worth my time, I'll help you fix it.

Go isn't a difficult language. You can probably fix most issues yourself (I have like 50 lines of single-line comments, don't cry about documentation.)

---

## TODO:
```
[X] -- Change commandline struct names
[X] -- Implement cat()
[X] -- Implement help()
[X] -- Implement exec()
[X] -- Implement rm/rf()
[ ] -- Enumeration scripts?
[ ] -- Implement download()
[ ] -- Implement upload()
[X] -- Implement remote()
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
