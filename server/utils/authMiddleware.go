package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Read the content
		buf, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			log.Print("bodyErr ", bodyErr.Error())
			http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
			return
		}

		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		// Restore the to its original state
		r.Body = rdr2
		// manipulate rd1 only
		// log.Printf("BODY: %q", rdr1)
		decoder := json.NewDecoder(rdr1)

		var t test_struct
		err := decoder.Decode(&t)

		if err != nil {
			panic(err)
		}
		log.Println("json quergoogleOAuth2 ---", t.Query)
		///////////////
		ctx := r.Context()
		token := r.Header.Get("Authorization")

		// utils.ParseRequest(w, r)

		jwt, err := CheckToken(token)
		if err != nil {
			fmt.Println(err)
		} else {
			// WARNING: the token was also added within index.html to simulate
			// requests with an Authorization header
			cookie := &http.Cookie{Name: "lychee", Value: token, HttpOnly: false}
			http.SetCookie(w, cookie) // TODO: set the cookie from somewhere more useful
			// https://upgear.io/blog/golang-tip-wrapping-http-response-writer-for-middleware/
		}

		ctx = context.WithValue(ctx, "jwt", jwt)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
