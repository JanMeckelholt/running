package dependencies

import "fmt"

type Config struct {
	Address string
}

var Configs = map[string]Config{
	"strava-service": {Address: fmt.Sprintf("strava-service:%d", 666)},
}
