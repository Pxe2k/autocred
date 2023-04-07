package transport

import (
	"autocredit/cmd/api/middlewares"
	"net/http"
)

func (server *Server) InitializeRoutes() {
	fs := http.FileServer(http.Dir("./storage/"))
	server.Router.PathPrefix("/storage/").Handler(http.StripPrefix("/storage/", fs))

	authApi := server.Router.PathPrefix("/api/auth").Subrouter()

	authApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createUser)).Methods("POST")
	authApi.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.login)).Methods("POST")
	authApi.HandleFunc("/submit", middlewares.SetMiddlewareJSON(server.submit)).Methods("POST")

	clientApi := server.Router.PathPrefix("/api/client").Subrouter()

	clientApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createClient)).Methods("POST")
	clientApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allClients)).Methods("GET")
	clientApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.getClient)).Methods("GET")
	clientApi.HandleFunc("/upload-avatar/{id}", middlewares.SetMiddlewareJSON(server.uploadAvatar)).Methods("PUT")

	pledgeApi := server.Router.PathPrefix("/api/pledge").Subrouter()

	pledgeApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createPledge)).Methods("POST")
	pledgeApi.HandleFunc("/all/{id}", middlewares.SetMiddlewareJSON(server.allPledges)).Methods("GET")
	pledgeApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.getPledge)).Methods("GET")

	applicationApi := server.Router.PathPrefix("/api/application").Subrouter()

	applicationApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createApplication)).Methods("POST")
	applicationApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allApplications)).Methods("GET")
	applicationApi.HandleFunc("/get/{id}", middlewares.SetMiddlewareJSON(server.getApplication)).Methods("GET")

	bankApi := server.Router.PathPrefix("/api/bank").Subrouter()

	bankApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createBank)).Methods("POST")
	bankApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allBank)).Methods("GET")
	bankApi.HandleFunc("/sign", middlewares.SetMiddlewareJSON(server.signApplication)).Methods("POST")
	bankApi.HandleFunc("/product", middlewares.SetMiddlewareJSON(server.createProduct)).Methods("POST")

	workApi := server.Router.PathPrefix("/api/work").Subrouter()

	workApi.HandleFunc("/create-activity", middlewares.SetMiddlewareJSON(server.createWorkActivity)).Methods("POST")
	workApi.HandleFunc("/create-title", middlewares.SetMiddlewareJSON(server.createJobTitle)).Methods("POST")
	workApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allWorkActivity)).Methods("GET")

	cityApi := server.Router.PathPrefix("/api/city").Subrouter()

	cityApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allCity)).Methods("GET")
	cityApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.findCityById)).Methods("GET")

	templateApi := server.Router.PathPrefix("/api/template").Subrouter()
	templateApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.GenerateTemplate)).Methods("POST")

	carApi := server.Router.PathPrefix("/api/cars").Subrouter()

	carApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createCarBrand)).Methods("POST")
	carApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allCarBrands)).Methods("GET")

	insuranceApi := server.Router.PathPrefix("/api/insurance").Subrouter()

	insuranceApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allInsurance)).Methods("GET")
	insuranceApi.HandleFunc("/create-insurance", middlewares.SetMiddlewareJSON(server.createInsurance)).Methods("POST")
	insuranceApi.HandleFunc("/create-kasko", middlewares.SetMiddlewareJSON(server.createKasko)).Methods("POST")
	insuranceApi.HandleFunc("/create-road-help", middlewares.SetMiddlewareJSON(server.createRoadHelp)).Methods("POST")
	insuranceApi.HandleFunc("/create-life-insurance", middlewares.SetMiddlewareJSON(server.createLifeInsurance)).Methods("POST")

	countryApi := server.Router.PathPrefix("/api/country").Subrouter()

	countryApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allCountries)).Methods("GET")

	documentApi := server.Router.PathPrefix("/api/document").Subrouter()

	documentApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.issuingAuthorityAll)).Methods("GET")
	documentApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.uploadFile)).Methods("POST")
}
