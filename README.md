## vcs
Read the vcs information (BuildInfo data) embedded in the running go binary.

#### requirements
Note: Go BuildInfo's VCS data to be available in the running binary the following requirements must be satisfied:
- Binary must be built with Go 1.18 or greater
- When the binary is built, the build context folder must already be initialized by the vcs system (e.g. for git, `git init`)
- The `go build ...` invocation must include the whole directory and not just specific go files (e.g. `go build .`)

#### usage

```go
package main

import (
	"github.com/damdo/vcs"
	"fmt"
)

func main() {
	v, ok := vcs.ReadInfo()
	if ok {
		fmt.Println(v.Vcs, v.Revision, v.Time, v.Modified)
	}
}
```

or if you prefer to gather BuildInfo on your own:

```go
package main

import (
	"github.com/damdo/vcs"
	"fmt"
	"runtime/debug"
)

func main() {
	b, ok := debug.ReadBuildInfo()
	if ok {
		v, ok := vcs.FromBuildInfo(b)
		if ok {
			fmt.Println(v.Vcs, v.Revision, v.Time, v.Modified)
		}
	}
}
```
