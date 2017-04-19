package chef

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
)

const stubPolicyLock = "test/policy.lock.json"

func TestPolicyFromJSONDecoder(t *testing.T) {
	if file, err := os.Open(stubPolicyLock); err != nil {
		t.Error("unexpected error", err, "during os.Open on ", stubPolicyLock)
	} else {
		decoder := json.NewDecoder(file)
		var p Policy
		if err := decoder.Decode(&p); err == io.EOF {
			log.Println(p)
		} else if err != nil {
			log.Fatal(err)
		}
	}

}
