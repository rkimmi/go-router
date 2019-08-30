package main

import ( 
"net/http"

m "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}    

// func main() {

// 	e := echo.New()

// 	// Middleware
// 	// Server header
// 	// e.Use(ServerHeader)
// 	// e.Use(CustomEchoLogger())
// 	// e.Use(middleware.Recover())
// 	// e.Use(middleware.RequestID())
// 	// e.Use(middleware.Gzip())
// 	// e.Use(middleware.Secure())

// 	// CORS
// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{
// 			echo.GET, echo.HEAD, echo.OPTIONS, echo.POST, echo.PATCH},
// 	}))

// 	e.GET("/", hello)

//    // http.HandleFunc("/", hello)
//     http.ListenAndServe(":8080", nil)
// }

// func hello(w http.ResponseWriter, r *http.Request) {
//  w.Header().Set("Access-Control-Allow-Origin", "*")
//    w.Write([]byte("hello!"))
// }
