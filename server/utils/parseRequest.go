package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type test_struct struct {
	Query string
}

func ParseRequest(w http.ResponseWriter, r *http.Request) {

	////////////////0//////////////////
	decoder := json.NewDecoder(r.Body)

	var t test_struct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Query)

	// ////////////////1/////////////////

	// buf, bodyErr := ioutil.ReadAll(r.Body)
	// if bodyErr != nil {
	// 	log.Print("bodyErr ", bodyErr.Error())
	// 	http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// log.Printf("BODY: %q", rdr1)

	// r.Body = rdr2
	// // fmt.Fprintf(w, "%q", dump)

	// ///////////////////////2///////////////////
	// // Read the content
	// var bodyBytes []byte
	// if r.Body != nil {
	// 	bodyBytes, _ = ioutil.ReadAll(r.Body)
	// }
	// // Restore the io.ReadCloser to its original state
	// r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// // Use the content
	// bodyString := string(bodyBytes)
	// log.Println(bodyString)

}
