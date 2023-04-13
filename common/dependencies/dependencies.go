package dependencies

type Config struct {
	Port int
}

var Configs = map[string]Config{
	"strava-service":   {Port: 666},
	"database-service": {Port: 666},
	"httpGateway":      {Port: 80},
	"runner":           {Port: 666},
	"populate-db":      {Port: 666},
}
