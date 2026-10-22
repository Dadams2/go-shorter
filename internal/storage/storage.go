package storage

import (
	"errors"

	"github.com/dadams2/go-shorter/internal/token"
)

// global bad no?
var m = make(map[string]string)

func StoreURL(url string) (*string, error) {
	str, err := token.GenerateToken();
	if err != nil {
		return nil, err;
	}
	m[*str] = url;
	return str, nil;
}

func RetrieveURL(token string) (*string, error) {
	url, ok := m[token];
	if !ok {
		return nil, errors.New("url not found");
	}
	return &url, nil;
}
