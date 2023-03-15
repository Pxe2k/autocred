package transport

import "autocredit/cmd/api/middlewares"

func (server *Server) InitializeRoutes() {
	authApi := server.Router.PathPrefix("/api/auth").Subrouter()

	authApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createUser)).Methods("POST")
	authApi.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.login)).Methods("POST")
	authApi.HandleFunc("/submit", middlewares.SetMiddlewareJSON(server.submit)).Methods("POST")

	clientApi := server.Router.PathPrefix("/api/client").Subrouter()

	clientApi.HandleFunc("/create", middlewares.SetMiddlewareJSON(server.createClient)).Methods("POST")
	clientApi.HandleFunc("/all", middlewares.SetMiddlewareJSON(server.allClients)).Methods("GET")
	clientApi.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(server.getClient)).Methods("GET")

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
}
