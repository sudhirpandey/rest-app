package app

import (
    "fmt"
    "log"
	"net/http"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"rest-app/config"
	"rest-app/app/model"
	"rest-app/app/handler"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//creater a new mux ruter from the above routes
func (a *App) SetUpRouter(){

	var routes = Routes{
		Route{"Index", "GET", "/", Index},
		Route{"AllProdutcs", "GET", "/products", a.GetAllProducts},
		Route{"CreateProduct", "POST", "/product", a.CreateProduct},
		Route{"GetProduct", "GET", "/product/{name}", a.GetProduct},
		Route{"DeleteProduct", "DELETE", "/product/{name}", a.DeleteProduct},
		Route{"UpdateProduct", "PUT", "/product/{name}", a.UpdateProduct},
	}

	for _, route := range routes {
		a.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(prometheus.InstrumentHandler(route.Name,route.HandlerFunc))
	}
	a.Router.Handle("/metrics", promhttp.Handler())
}


// initialise the app to load the router as property 
func (app *App) Initialize(config *config.Config) {

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Name,
		config.DB.Charset)

	// println(config.DB.Host)
	db, err := gorm.Open(config.DB.Dialect, dbURI)

	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
    app.DB = model.DBMigrate(db)
	app.Router = mux.NewRouter().StrictSlash(true)
	app.SetUpRouter()
}

//listen and serve the app with endpoints from the router
func (app *App) Run(host string) {

  c := cors.New(cors.Options{
      AllowedOrigins: []string{"*"}, // All origins
      AllowedMethods: []string{"GET"}, // Allowing only get, just an example
  })
	println("Server up and listening....")
	log.Fatal(http.ListenAndServe(host, c.Handler(app.Router)))
}

//List of routes functions
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sample Rest Api endpint: v1.0")
    fmt.Println("Endpoint Hit: indexPage")
}

func (a *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	//prometheus.InstrumentHandler("webkv", s)
	handler.GetAllProducts(a.DB, w, r)
}

func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(a.DB, w, r)
}

func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	handler.GetProduct(a.DB, w, r)
}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	handler.DeleteProduct(a.DB, w, r)
}

func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	handler.UpdateProduct(a.DB, w, r)
}
