package SkyLine_Backend

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"strings"
)

var HashMapper = map[string]func(string) string{
	"MD5": func(dt string) string {
		return fmt.Sprintf("%x", md5.Sum([]byte(dt)))
	},
	"SHA1": func(dt string) string {
		return fmt.Sprintf("%x", sha1.Sum([]byte(dt)))
	},
	"SHA256": func(dt string) string {
		return fmt.Sprintf("%x", sha256.Sum256([]byte(dt)))
	},
	"SHA512": func(dt string) string {
		return fmt.Sprintf("%x", sha512.Sum512([]byte(dt)))
	},
}

func Crypto_Hasher(args ...Object) Object {
	if len(args) != 2 {
		return NewError("Sorry but crypt.hash requires 2 positional based arguments, you gave %d but the function wants 2 ", len(args))
	}
	var hashtype, data string
	hashtype = strings.ToUpper(args[0].Inspect())
	data = args[1].Inspect()
	if HashMapper[hashtype] != nil {
		return &String{Value: HashMapper[hashtype](data)}
	} else {
		return NewError("crypt.hash requires a supported hash type, below are the supported hash types \n | -> MD5 \n | -> SHA1 \n | -> SHA256 \n | -> SHA512 ")
	}
}
