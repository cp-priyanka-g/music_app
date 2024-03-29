module music_app

go 1.17

replace db => ./db

replace artist => ./artist

replace register => ./register

replace album => ./album

replace track => ./track

replace playlist => ./playlist

replace favourite => ./favourite

require (
	album v0.0.0-00010101000000-000000000000
	artist v0.0.0-00010101000000-000000000000
	db v0.0.0-00010101000000-000000000000
	favourite v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.4
	playlist v0.0.0-00010101000000-000000000000
	register v0.0.0-00010101000000-000000000000
	track v0.0.0-00010101000000-000000000000
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect

)
