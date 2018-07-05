// package main

// import (
// 	"encoding/base64"
// 	"fmt"
// )

// func main() {
// 	test := base64.StdEncoding.EncodeToString([]byte("1"))
// 	fmt.Println(test)
// 	test2, err := base64.StdEncoding.DecodeString(test)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(test2))

// }
package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	uuid, err := newUUID()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("%s\n", uuid)
}

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
