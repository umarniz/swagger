// @APIVersion 1.0.0
// @APITitle Swagger Example API
// @APIDescription Swagger Example API
// @BasePath http://127.0.0.1:3000/
// @Contact varyous@gmail.com
// @TermsOfServiceUrl http://umarniz.com/
// @License BSD
// @LicenseUrl http://umarniz.com/
package main

import (
	"net/http"

	"github.com/umarniz/swagger/example"
)

func main() {
	router := example.InitRouter()
	http.ListenAndServe("localhost:3000", router)
}
