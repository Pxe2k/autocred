package transport

import (
	"autocredit/cmd/api/internal/storage"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	server.DB.Debug().AutoMigrate(
		storage.Role{},
		storage.User{},
		storage.IndividualClient{},
		storage.BusinessClient{},
		storage.Document{},
		storage.MaritalStatus{},
		storage.WorkPlaceInfo{},
		storage.CurrentLoans{},
		storage.ClientContact{},
		storage.BonusInfo{},
		storage.BeneficialOwnerIndividual{},
		storage.RegistrationAddressBusiness{},
		storage.BeneficialOwnerBusiness{},
		storage.MaritalStatusBusiness{},
		storage.WorkPlaceInfoBusiness{},
		storage.DocumentBusiness{},
		storage.BusinessContact{},
		storage.BonusInfoBusiness{},
		storage.CurrentLoanBusiness{},
		storage.RegistrationAddress{},
		storage.ResidentialAddress{},
		storage.BusinessClient{},
		storage.Media{},
		storage.Pledge{},
		storage.Application{},
		storage.BankApplication{},
		storage.Bank{},
		storage.WorkingActivity{},
		storage.City{},
		storage.BankProduct{},
		storage.Kasko{},
		storage.RoadHelp{},
		storage.LifeInsurance{},
		storage.CarBrand{},
		storage.CarModel{},
		storage.IssuingAuthority{}) //migrations

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port " + os.Getenv("APP_PORT"))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Location"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(addr, handlers.CORS(headersOk, methodsOk, originsOk)(server.Router)))
}
