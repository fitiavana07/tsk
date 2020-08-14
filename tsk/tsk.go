// tsk is the package providing the entrypoint
package tsk

import (
	"bytes"
	"fmt"
)

func TskMain(out *bytes.Buffer) {
	fmt.Fprint(out, "No task, good news!")
}
