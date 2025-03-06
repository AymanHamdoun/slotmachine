package middleware

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
)

// BindInput is a middleware that binds request data to a struct
func BindInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Store the original body for the next handlers
		r.ParseForm()
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// BindInputData binds request data to the provided struct based on the request method and content type
func BindInputData(r *http.Request, inputStruct interface{}) error {
	if r.Method == http.MethodPost && r.Header.Get("Content-Type") == "application/json" {
		// Marshal the request JSON body to the input struct
		err := json.NewDecoder(r.Body).Decode(inputStruct)
		if err != nil {
			if err.Error() == "EOF" {
				return nil
			}
			return err
		}
	} else if r.Method == http.MethodPost {
		// Bind form data
		err := bindForm(r.PostForm, inputStruct)
		if err != nil {
			return err
		}
	} else {
		// Bind query parameters
		err := bindForm(r.URL.Query(), inputStruct)
		if err != nil {
			return err
		}
	}
	return nil
}

// bindForm binds url.Values to a struct using reflection
func bindForm(values url.Values, dst interface{}) error {
	val := reflect.ValueOf(dst)
	if val.Kind() != reflect.Ptr {
		return nil
	}
	val = val.Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := typ.Field(i)
		formKey := typeField.Tag.Get("form")
		if formKey == "" {
			formKey = typeField.Name
		}

		if values.Get(formKey) != "" && field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(values.Get(formKey))
				// Add more types as needed
			}
		}
	}
	return nil
}
