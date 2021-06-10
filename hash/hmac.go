package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"hash"
)

func NewHMAC(key string) HMAC{
    h := hmac.New(sha512.New,[]byte(key))
    return HMAC{hmac:h}
}

/*
** HMAC is crypto/hmac wrapper
** it little easier to use 
*/
type HMAC struct{
    hmac hash.Hash
}
/*
** Hash Function take input string and hash it
** Withe secret key when HMAC object was created
*/
func(h HMAC) Hash(data string) string{
    h.hmac.Reset()
    h.hmac.Write([]byte(data))
    hmacByte  := h.hmac.Sum(nil)
    return base64.URLEncoding.EncodeToString(hmacByte)
}