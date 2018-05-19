package logintest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/astenmies/lychee/server/gqlResolver"

	"github.com/astenmies/lychee/server/gqlSchema"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var starwarsSchema = graphql.MustParseSchema(gqlSchema.GetRootSchema(), &gqlResolver.Resolver{})

type test_struct struct {
	Query string
}

func parseRequest(w http.ResponseWriter, r *http.Request) {

	////////////////0//////////////////
	decoder := json.NewDecoder(r.Body)

	var t test_struct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Query)
}
func TestServeHTTP(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", strings.NewReader(`{"query":"{ post(slug: \"second\") { title } }", "operationName":"", "variables": null}`))

	parseRequest(w, r)

	h := relay.Handler{Schema: starwarsSchema}

	h.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Fatalf("Expected status code 200, got %d.", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("Invalid content-type. Expected [application/json], but instead got [%s]", contentType)
	}

	expectedResponse := `{"data":{"post":{"title":"Hello second post"}}}`
	actualResponse := w.Body.String()
	if expectedResponse != actualResponse {
		t.Fatalf("Invalid response. Expected [%s], but instead got [%s]", expectedResponse, actualResponse)
	}
}
