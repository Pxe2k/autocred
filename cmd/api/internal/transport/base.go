package transport

import (
	"autocredit/cmd/api/internal/storage"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	dsn := "host=" + DbHost + " " + "user=" + DbUser + " " + "password=" + DbPassword + " " + "dbname=" + DbName + " " + "port=" + DbPort + " " + "sslmode=disable TimeZone=Asia/Shanghai"
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database \n", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(storage.User{},
		storage.Client{},
		storage.Document{},
		storage.WorkPlaceInfo{},
		storage.MaritalStatus{},
		storage.RelationWithBank{},
		storage.RegistrationAddress{},
		storage.ResidentialAddress{},
		storage.ClientContact{},
		storage.BeneficialOwner{},
		storage.ClientComment{},
		storage.BonusInfo{},
		storage.PersonalProperty{},
		storage.CurrentLoans{},
		storage.Pledge{},
		storage.Application{},
		storage.Bank{},
		storage.BankResponse{},
		storage.WorkingActivity{},
		storage.JobTitle{},
		storage.City{},
		storage.BankProduct{},
		storage.Media{},
		storage.CarBrand{},
		storage.CarModel{},
		storage.Insurance{},
		storage.Kasko{},
		storage.RoadHelp{},
		storage.LifeInsurance{},
		storage.BankApplication{}) //migrations

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port " + os.Getenv("APP_PORT"))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Location"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(addr, handlers.CORS(headersOk, methodsOk, originsOk)(server.Router)))
}
