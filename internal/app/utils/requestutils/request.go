package requestutils

import (
	"errors"
	"fmt"
	"net/http"
)

func GetCookieValue(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)
	if err != nil {
		return "", err
	}

	if cookie == nil {
		return "", errors.New(fmt.Sprintf("%s is nil", key))
	}

	return cookie.Value, nil
}
