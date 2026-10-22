
package token

import (
	"crypto/rand"
	"encoding/hex"
)


func GenerateToken() (*string, error) {
	b := make([]byte, 8);
	_, err := rand.Read(b);
	if err != nil {
		return nil, err;		
	}
	str := hex.EncodeToString(b);
	return &str, nil;
	
}

