package diff

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
)

func init() {
	os.RemoveAll("./out/diff")
	os.MkdirAll("./out/diff", 0755)
}

func Write(name string, old, new runtime.Object) {
	diff := cmp.Diff(old, new)
	if diff != "" {
		ioutil.WriteFile(fmt.Sprintf("./out/diff/%s-%s.txt", time.Now().Format("15:04:05.000000"), name), []byte(diff), 0644)
	}
}
