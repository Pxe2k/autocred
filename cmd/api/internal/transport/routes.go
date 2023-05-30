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
	authApi.HandleFunc("/role", middlewares.SetMiddlewareJSON(server.getRoleID)).Methods("GET")

	clientApi := server.Router.PathPrefix("/api/client").Subrouter()

	clientApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createClient)).Methods("POST")
	clientApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allClients)).Methods("GET")
	clientApi.HandleFunc("/filter", middlewares.SetMiddlewareJSON(server.allClients)).Methods("GET")
	clientApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.getClient)).Methods("GET")
	clientApi.HandleFunc("/upload-avatar/{id}", middlewares.SetMiddlewareJSON(server.uploadAvatar)).Methods("PUT")
	clientApi.HandleFunc("/otp", middlewares.SetMiddlewareJSON(server.generateClientOTP)).Methods("POST")
	clientApi.HandleFunc("/submit", middlewares.SetMiddlewareJSON(server.submitClientOTP)).Methods("PATCH")

	pledgeApi := server.Router.PathPrefix("/api/pledge").Subrouter()

	pledgeApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createPledge)).Methods("POST")
	pledgeApi.HandleFunc("/all/{id}", middlewares.SetMiddlewareJSON(server.allPledges)).Methods("GET")
	pledgeApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.getPledge)).Methods("GET")

	applicationApi := server.Router.PathPrefix("/api/application").Subrouter()

	applicationApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createApplication)).Methods("POST")
	applicationApi.HandleFunc("/create-bcc", middlewares.SetMiddlewareJSON(server.createBCCApplication)).Methods("POST")
	applicationApi.HandleFunc("/create-eu", middlewares.SetMiddlewareJSON(server.createEUApplication)).Methods("POST")
	applicationApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allApplications)).Methods("GET")
	applicationApi.HandleFunc("/get/{id}", middlewares.SetMiddlewareJSON(server.getApplication)).Methods("GET")
	applicationApi.HandleFunc("/token", middlewares.SetMiddlewareJSON(server.getBankToken)).Methods("GET")
	applicationApi.HandleFunc("/response-bcc", middlewares.SetMiddlewareJSON(server.getBCCResponse)).Methods("POST")

	bankApi := server.Router.PathPrefix("/api/bank").Subrouter()

	bankApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createBank)).Methods("POST")
	bankApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allBank)).Methods("GET")
	bankApi.HandleFunc("/product/{id}", middlewares.SetMiddlewareJSON(server.getBankProduct)).Methods("GET")
	bankApi.HandleFunc("/product", middlewares.SetMiddlewareJSON(server.createProduct)).Methods("POST")
	bankApi.HandleFunc("/update/{id}", middlewares.SetMiddlewareJSON(server.updateBank)).Methods("PATCH")
	bankApi.HandleFunc("/update-product/{id}", middlewares.SetMiddlewareJSON(server.updateProduct)).Methods("PATCH")
	bankApi.HandleFunc("/delete/{id}", middlewares.SetMiddlewareJSON(server.deleteBank)).Methods("DELETE")
	bankApi.HandleFunc("/delete-product/{id}", middlewares.SetMiddlewareJSON(server.deleteBankProduct)).Methods("DELETE")

	workApi := server.Router.PathPrefix("/api/work").Subrouter()

	workApi.HandleFunc("/create-activity", middlewares.SetMiddlewareJSON(server.createWorkActivity)).Methods("POST")
	workApi.HandleFunc("/update-activity/{id}", middlewares.SetMiddlewareJSON(server.updateWorkActivity)).Methods("PATCH")
	workApi.HandleFunc("/delete-activity/{id}", middlewares.SetMiddlewareJSON(server.deleteWorkActivity)).Methods("DELETE")
	workApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allWorkActivity)).Methods("GET")

	cityApi := server.Router.PathPrefix("/api/city").Subrouter()

	cityApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allCity)).Methods("GET")
	cityApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.findCityById)).Methods("GET")

	autoDealerApi := server.Router.PathPrefix("/api/autodealer").Subrouter()

	autoDealerApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allAutoDealers)).Methods("GET")
	autoDealerApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.getAutoDealer)).Methods("GET")
	autoDealerApi.HandleFunc("/delete/{id}", middlewares.SetMiddlewareJSON(server.deleteAutoDealer)).Methods("DELETE")
	autoDealerApi.HandleFunc("/update/{id}", middlewares.SetMiddlewareJSON(server.updateAutoDealer)).Methods("PATCH")
	autoDealerApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createAutoDealer)).Methods("POST")

	templateApi := server.Router.PathPrefix("/api/template").Subrouter()
	templateApi.HandleFunc("/create/{id}", middlewares.SetMiddlewareJSON(server.generateTemplate)).Methods("POST")

	carApi := server.Router.PathPrefix("/api/cars").Subrouter()

	carApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createCarBrand)).Methods("POST")
	carApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allCarBrands)).Methods("GET")

	insuranceApi := server.Router.PathPrefix("/api/insurance").Subrouter()

	insuranceApi.HandleFunc("/create-kasko", middlewares.SetMiddlewareJSON(server.createKasko)).Methods("POST")
	insuranceApi.HandleFunc("/create-road-help", middlewares.SetMiddlewareJSON(server.createRoadHelp)).Methods("POST")
	insuranceApi.HandleFunc("/create-life-insurance", middlewares.SetMiddlewareJSON(server.createLifeInsurance)).Methods("POST")
	insuranceApi.HandleFunc("/update-kasko/{id}", middlewares.SetMiddlewareJSON(server.updateKasko)).Methods("PATCH")
	insuranceApi.HandleFunc("/update-road-help/{id}", middlewares.SetMiddlewareJSON(server.updateRoadHelp)).Methods("PATCH")
	insuranceApi.HandleFunc("/update-life-insurance/{id}", middlewares.SetMiddlewareJSON(server.updateLifeInsurance)).Methods("PATCH")
	insuranceApi.HandleFunc("/kasko/{id}", middlewares.SetMiddlewareJSON(server.getKasko)).Methods("GET")
	insuranceApi.HandleFunc("/road-help/{id}", middlewares.SetMiddlewareJSON(server.getRoadHelp)).Methods("GET")
	insuranceApi.HandleFunc("/life-insurance/{id}", middlewares.SetMiddlewareJSON(server.getLifeInsurance)).Methods("GET")
	insuranceApi.HandleFunc("/delete-kasko/{id}", middlewares.SetMiddlewareJSON(server.deleteKasko)).Methods("DELETE")
	insuranceApi.HandleFunc("/delete-road-help/{id}", middlewares.SetMiddlewareJSON(server.deleteRoadHelp)).Methods("DELETE")
	insuranceApi.HandleFunc("/delete-life-insurance/{id}", middlewares.SetMiddlewareJSON(server.softDeleteLifeInsurance)).Methods("DELETE")

	countryApi := server.Router.PathPrefix("/api/country").Subrouter()

	countryApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allCountries)).Methods("GET")

	documentApi := server.Router.PathPrefix("/api/document").Subrouter()

	documentApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.issuingAuthorityAll)).Methods("GET")
	documentApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.uploadFile)).Methods("POST")
	documentApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.deleteMedia)).Methods("DELETE")

	userApi := server.Router.PathPrefix("/api/user").Subrouter()

	userApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allUsers)).Methods("GET")
	userApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.getUser)).Methods("GET")
	userApi.HandleFunc("/update/{id}", middlewares.SetMiddlewareJSON(server.updateUser)).Methods("PATCH")
	userApi.HandleFunc("/deactivate/{id}", middlewares.SetMiddlewareJSON(server.deactivateUser)).Methods("DELETE")
	userApi.HandleFunc("/all-deleted/{id}", middlewares.SetMiddlewareJSON(server.allDeactivatedUsers)).Methods("GET")
	userApi.HandleFunc("/activate/{id}", middlewares.SetMiddlewareJSON(server.recoverUser)).Methods("PATCH")

}
