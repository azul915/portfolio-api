module portfolio-api

go 1.13

require (
	cloud.google.com/go/firestore v1.1.1
	firebase.google.com/go v3.12.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	google.golang.org/api v0.17.0
)

// https://github.com/oxequa/realize/issues/253#issuecomment-572871603
replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
