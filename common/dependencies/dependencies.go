package dependencies

type Config struct {
	Port int
}

var Configs = map[string]Config{
	"strava-service":   {Port: 666},
	"database-service": {Port: 666},
	"http_gateway-API": {Port: 333},
	"http_gateway":     {Port: 80},
	"runner":           {Port: 666},
	"runner_frontend":  {Port: 443},
	"populate_db":      {Port: 666},
}
