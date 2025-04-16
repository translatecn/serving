package diff

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"time"
)

func init() {
	os.RemoveAll("./diff")
	os.MkdirAll("./diff", 0755)
}

func Write(name string, old, new runtime.Object) {
	diff := cmp.Diff(old, new)
	if diff != "" {
		ioutil.WriteFile(fmt.Sprintf("./diff/%s-%s.txt", time.Now().Format("15:04:05.000000"), name), []byte(diff), 0644)
	}
}
