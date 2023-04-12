package dependencies

type Config struct {
	Port int
}

var Configs = map[string]Config{
	"strava-service":   {Port: 666},
	"database-service": {Port: 666},
	"runner-http":      {Port: 80},
	"runner-grpc":      {Port: 666},
	"populate-db":      {Port: 666},
}
