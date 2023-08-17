package encode

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

func Sha512WithSalt(str, salt string) string {
	val := strings.ToLower(str)
	val = salt + val
	fmt.Println("to lower:", val)
	hash := sha512.New()
	hash.Write([]byte(val))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
