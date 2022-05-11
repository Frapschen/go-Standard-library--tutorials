package fmt_tutorial

import (
	"fmt"
	"os"
	"testing"
)

func TestPrntln(t *testing.T) {
	//默认formats
	fmt.Fprint(os.Stdout, "aaa")
}
