package main

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/JanMeckelholt/running/common/commonconf"
	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/common/utils"
	"github.com/JanMeckelholt/running/populate_db/service"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	commonConf := commonconf.GPGConf{}
	err := env.Parse(&commonConf)
	if err != nil {
		return
	}
	err = utils.DecryptPGP("./volumes-data/env/.env.docker.secret.asc", "./secret/env/.env.docker.secret", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt env: %s", err.Error())
		return
	}
	err = godotenv.Load("./volumes-data/env/.env.docker", "./secret/env/.env.docker.secret")
	if err != nil {
		log.Errorf("Could not load env: %s", err.Error())
		return
	}
	srv := &service.Service{}
	err = env.Parse(&srv.ServiceConfig)
	if err != nil || !srv.ServiceConfig.Enabled {
		return
	}
	err = srv.Clients.Dial(srv.ServiceConfig)
	if err != nil {
		log.Errorf("could not Dial Clients! %s", err.Error())
	}
	srv.Clients.DatabaseClient.UpsertClient(context.Background(), &database.Client{
		ClientId:     srv.ServiceConfig.ClientsConfig.ClientId,
		ClientSecret: srv.ServiceConfig.ClientsConfig.ClientSecret,
	})

	srv.Clients.DatabaseClient.UpsertAthlete(context.Background(), &database.Athlete{
		Token:        srv.ServiceConfig.ClientsConfig.Token,
		RefreshToken: srv.ServiceConfig.ClientsConfig.RefreshToken,
		AthleteId:    srv.ServiceConfig.ClientsConfig.AthletId,
		ClientId:     srv.ServiceConfig.ClientsConfig.ClientId,
	})

	records := readCsvFile("./volumes-data/data/activities.csv")
	for i := 1; i < len(records); i++ {
		distance, _ := strconv.ParseFloat(records[i][17], 64)
		totalElevationGain, _ := strconv.ParseFloat(records[i][20], 64)
		averageSpeed, _ := strconv.ParseFloat(records[i][19], 64)
		maxSpeed, _ := strconv.ParseFloat(records[i][18], 64)
		averageHR, _ := strconv.ParseFloat(records[i][31], 64)
		maxHR, _ := strconv.ParseFloat(records[i][30], 64)
		elevHigh, _ := strconv.ParseFloat(records[i][23], 64)
		elevLow, _ := strconv.ParseFloat(records[i][22], 64)

		movingTime, _ := strconv.ParseUint(records[i][16], 10, 64)
		elapsedTime, _ := strconv.ParseUint(records[i][5], 10, 64)
		id, _ := strconv.ParseUint(records[i][0], 10, 64)

		activity := strava.Activity{
			Athlete: &strava.AthleteType{
				Id: 7845894,
			},
			Name:               records[i][2],
			Distance:           distance,
			MovingTime:         movingTime,
			ElapsedTime:        elapsedTime,
			TotalElevationGain: totalElevationGain,
			Type:               records[i][3],
			Id:                 id,
			StartDate:          records[i][1],
			AverageSpeed:       averageSpeed,
			MaxSpeed:           maxSpeed,
			AverageHeartrate:   averageHR,
			MaxHeartrate:       maxHR,
			ElevHigh:           elevHigh,
			ElevLow:            elevLow,
		}

		srv.Clients.DatabaseClient.UpsertActivityFromCSV(context.Background(), &activity)

	}
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
