package rand

import (
	"crypto/rand"
	"encoding/base64"
)
const RememberTokenByte = 32
/*
** Bytes will generate n random  bytes
** or return an error if there was one
 */
 func Byte(n uint)([]byte, error){
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil{
        return nil, err
    }    
    return b, nil
 }
 

 /*
 ** String will generate a byte slice of size numBytes
 ** and return a string that is the base64 url encoded
 */
 func String(numBytes uint)(string, error){
     b, err := Byte(numBytes)
     if err != nil {
         return "", err
     }
 
     return base64.URLEncoding.EncodeToString(b), nil 
 }

 func  NBytes( base64String  string) (int, error){
  b, err := base64.URLEncoding.DecodeString(base64String)
  if err != nil{
      return -1, err
  }
    return len(b), nil
 }
/*
** RememberToken is function to generate token
*/
 func RememberToken()(string, error){
     return String(RememberTokenByte)
 }
