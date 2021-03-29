module github.com/darinmilner/productiveapp

// +heroku goVersion go1.16
go 1.16

// +herolu install ./cmd/...
require (
	github.com/alexedwards/scs/v2 v2.4.0
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/hablullah/go-hijri v1.0.2
	github.com/heroku/x v0.0.28
	github.com/joho/godotenv v1.3.0
	github.com/justinas/nosurf v1.1.1
	go.mongodb.org/mongo-driver v1.5.0
)
