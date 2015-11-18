package riffle

import (
	"math/rand"
	"time"
)

const (
	maxId int64 = 1 << 53
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewID generates a random WAMP uint.
func newID() uint {
	return uint(rand.Int63n(maxId))
}

// func PprintMap(m interface{}) {
// 	if b, err := jSON.MarshalIndent(m, "", "  "); err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		//log.Printf("%s", string(m))
// 	}
// }
