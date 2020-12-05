package routing

import (
	. "github.com/alameddinc/terminal-wars/controllers"
	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func routes() {
	homeRoutes()
	userRoutes()
	appRoutes()
}

func homeRoutes() {
	r.HandleFunc("/", GetHome).Methods("GET")
}

func appRoutes() {
	ar := createSubRouter("app")
	ar.HandleFunc("/mining/get", GetMining).Methods("GET")     // <- 3xHEX
	ar.HandleFunc("/mining/set", PostMining).Methods("POST")    // -> 3xHex + String
	ar.HandleFunc("/fishing", PostFishing).Methods("POST")       // -> x, y  <- Point
	ar.HandleFunc("/repair/get", GetRepair).Methods("GET") // <- md5
	ar.HandleFunc("/repair/set", PostRepair).Methods("POST") // -> string(9 haneli sayı dönecek)
	ar.HandleFunc("/farming/get", GetFarming).Methods("GET") // <- file.txt yollayacak (2000/100.000.000)
	ar.HandleFunc("/farming/set", PostFarming).Methods("POST") // <- 1000 tane kordinat (x/50)
	// Korsanlık
	ar.HandleFunc("/watch/{location}", GetWatch).Methods("GET") // Wanted, Farming, Fishing, repair ve Mining alanları izlenecek // Burada 24 saat içinde işlem yapılanlar görülecek
	ar.HandleFunc("/attack/", PostAttack).Methods("POST") // -> array[2000], user_address(h://123.43.22) <-
	// <- eğer başarılıysa Saldırana madurun altınların %50'sini vereceğiz ve başarılı mesajı döneceğiz
	// <- eğer başarısızsa Saldırana fail mesajı döneceğiz
	// <- madurun log kayıtlarına saldırıya uğradığını söyleyeceğiz durum bilgisi vereceğiz.
	// Saldıran 1 saatte sadece 50 istek atabilir.


}

func userRoutes() {
	ur := createSubRouter("user")
	ur.HandleFunc("/register", PostRegister).Methods("POST")
	ur.HandleFunc("/logs", GetLogs).Methods("GET") // dakikada 10 saniyede yenilecek ve 15 dakikada en fazla 1000 istek
	ur.HandleFunc("/login", PostLogin).Methods("POST")
	ur.HandleFunc("/change/password", PostChangePassword).Methods("POST")
	ur.HandleFunc("/change/hash", PostChangeHash).Methods("POST")
	ur.HandleFunc("/{username}", GetUser).Methods("GET")
}
