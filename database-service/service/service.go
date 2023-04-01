package service

import (
	"errors"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/database-service/service/config"
)

var DB *gorm.DB

type DBClient struct {
	gorm.Model
	ClientId     *string `gorm:"unique;not null"`
	ClientSecret *string
	Token        *string
	RefreshToken *string
	Activities   []DBActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type DBActivity struct {
	gorm.Model
	DBClientID         uint
	Name               string  `json:"name,omitempty"`
	Distance           float64 `json:"distance,omitempty"`
	MovingTime         int64   `json:"moving_time,omitempty"`
	ElapsedTime        int64   `json:"elapsed_time,omitempty"`
	TotalElevationGain float64 `json:"total_elevation_gain,omitempty"`
	Type               string  `json:"type,omitempty"`
	SportType          string  `json:"sport_type,omitempty"`
	Id                 int64   `json:"id,omitempty"`
	StartDate          string  `json:"start_date,omitempty"`
	StartDateLocale    string  `json:"start_date_locale,omitempty"`
	Timezone           string  `json:"timezone,omitempty"`
	UtcOffset          float64 `json:"utc_offset,omitempty"`
	LocationCity       string  `json:"location_city,omitempty"`
	LocationState      string  `json:"location_state,omitempty"`
	LocationCountry    string  `json:"location_country,omitempty"`
	AchievementCount   int64   `json:"achievement_count,omitempty"`
	KudosCount         int64   `json:"kudos_count,omitempty"`
	CommentCount       int64   `json:"comment_count,omitempty"`
	Manual             bool    `json:"manual,omitempty"`
	Visibility         string  `json:"visibility,omitempty"`
	AverageSpeed       float64 `json:"average_speed,omitempty"`
	MaxSpeed           float64 `json:"max_speed,omitempty"`
	AverageHeartrate   float64 `json:"average_heartrate,omitempty"`
	MaxHeartrate       float64 `json:"max_heartrate,omitempty"`
	ElevHigh           float64 `json:"elev_high,omitempty"`
	ElevLow            float64 `json:"elev_low,omitempty"`
	ResourceState      int64   `json:"resource_state,omitempty"`
}

type Storer struct {
	StorerConfig config.StorerConfig
}

func NewStorer(storerConfig config.StorerConfig) *Storer {
	return &Storer{
		StorerConfig: storerConfig,
	}
}

func (s *Storer) InitStorage() error {
	log.Infof("Using ConnString for DB: host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", s.StorerConfig.HOST, s.StorerConfig.PORT, s.StorerConfig.PostgresDB, s.StorerConfig.PostgresUser, s.StorerConfig.PostgresPassword)
	var err error
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", s.StorerConfig.HOST, s.StorerConfig.PORT, s.StorerConfig.PostgresDB, s.StorerConfig.PostgresUser, s.StorerConfig.PostgresPassword)), &gorm.Config{})
	if err != nil {
		log.Error("DB error: could not open sql connection")
		panic("failed to connect to database")
	}
	DB = db
	log.Info("InitStorage finished")
	return nil
}

func (s *Storer) AutoMigrate(object interface{}) error {
	log.Info("Starting automatic migrations")

	var err error
	if err = DB.Debug().AutoMigrate(object); err != nil {
		log.Error("DB error: could not automigrate object")
		return err
	}
	log.Info("Automatic migrations finished")
	return nil
}

func (s *Storer) UpsertClient(clientId, clientSecret, token, refreshToken string) error {
	log.Infof("Creating: %s %s %s %s", clientId, clientSecret, token, refreshToken)
	result := DB.Create(&DBClient{ClientId: &clientId, ClientSecret: &clientSecret, Token: &token, RefreshToken: &refreshToken})
	if result.Error != nil {
		log.Error("DB error: could not create object")
		return result.Error
	}
	log.Info("Stored client: %s", result.RowsAffected)
	return nil
}

func (s *Storer) UpdateClient(clientId string, kvPairs []*grpcDB.KvPair) (*grpcDB.Client, error) {
	oldDBClient, err := s.GetDBClient(clientId)
	if err != nil {
		return nil, err
	}
	for _, kvPair := range kvPairs {
		log.Infof("Updating: %s %s %s ", clientId, kvPair.Key, kvPair.Value)
		result := DB.Model(oldDBClient).Update(kvPair.Key, &kvPair.Value)
		if result.Error != nil {
			log.Errorf("DB error: could not update client: %s", result.Error)
			return nil, result.Error
		}
	}
	log.Info("Stored client: %s", oldDBClient)
	return dbClientToClient(oldDBClient), nil
}
func (s *Storer) GetDBClient(clientId string) (*DBClient, error) {
	var (
		result   *gorm.DB
		dbClient DBClient
	)
	result = DB.Where(&DBClient{ClientId: &clientId}).First(&dbClient)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Infof("No record found with id %s", dbClient)
		return nil, nil
	}
	if result.Error != nil {
		log.Errorf("DB error: could not get object by id %s: %s", clientId, result.Error)
		return nil, result.Error
	}
	log.Infof("Retrieved client %s, %d, %s", *dbClient.ClientId, dbClient.ID, *dbClient.Token)
	return &dbClient, nil
}

func dbClientToClient(dbClient *DBClient) *grpcDB.Client {
	return &grpcDB.Client{
		ClientId:     *dbClient.ClientId,
		ClientSecret: *dbClient.ClientSecret,
		Token:        *dbClient.Token,
		RefreshToken: *dbClient.RefreshToken,
	}
}

func (s *Storer) UpsertActivity(req *grpcStrava.Activity) error {
	log.Infof("Creating: %s\n%s", req.GetName(), req.GetAthlete().String())
	clientId := strconv.FormatInt(req.GetAthlete().GetId(), 10)
	result := DB.Model(&DBClient{ClientId: &clientId}).Association("Activities").Append(DBActivity{
		Name:               req.GetName(),
		Distance:           req.GetDistance(),
		ElapsedTime:        req.GetElapsedTime(),
		AverageHeartrate:   req.GetAverageHeartrate(),
		AverageSpeed:       req.GetAverageSpeed(),
		TotalElevationGain: req.GetTotalElevationGain(),
	})
	if result != nil {
		log.Errorf("DB error: could not add activity %s", result.Error())
		return fmt.Errorf("DB error: could not add activity %s", result.Error())
	}
	log.Info("Added activity %s to client %s", req.GetName(), clientId)
	return nil
}

func dbActivityToActivity(dbActivity *DBActivity) *grpcStrava.Activity {
	return &grpcStrava.Activity{
		ResourceState:      dbActivity.ResourceState,
		Name:               dbActivity.Name,
		Distance:           dbActivity.Distance,
		MovingTime:         dbActivity.MovingTime,
		ElapsedTime:        dbActivity.ElapsedTime,
		TotalElevationGain: dbActivity.TotalElevationGain,
		Type:               dbActivity.Type,
		SportType:          dbActivity.SportType,
		StartDate:          dbActivity.StartDate,
		StartDateLocale:    dbActivity.StartDateLocale,
		Timezone:           dbActivity.Timezone,
		UtcOffset:          dbActivity.UtcOffset,
		LocationCity:       dbActivity.LocationCity,
		LocationState:      dbActivity.LocationState,
		LocationCountry:    dbActivity.LocationCountry,
		AchievementCount:   dbActivity.AchievementCount,
		KudosCount:         dbActivity.KudosCount,
		CommentCount:       dbActivity.CommentCount,
		Manual:             dbActivity.Manual,
		Visibility:         dbActivity.Visibility,
		AverageSpeed:       dbActivity.AverageSpeed,
		MaxSpeed:           dbActivity.MaxSpeed,
		AverageHeartrate:   dbActivity.AverageHeartrate,
		MaxHeartrate:       dbActivity.MaxHeartrate,
		ElevHigh:           dbActivity.ElevHigh,
		ElevLow:            dbActivity.ElevLow,
	}
}
