# Notes

### Lesson 1

- `go build` creates a large executable. This executable contains the golang runtime
  - `ls -lh` shows the size of the executable: includes scheduler, GC, etc
- `time go run *.go` shows how much time it took to compile and run the code
- `GOOS=linux go build` will build the exe for a linux machine (`darwin` is macOS)
