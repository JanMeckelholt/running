package service

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/database-service/service/config"

	"github.com/jackc/pgx"
)

var DB *gorm.DB

type DBClient struct {
	gorm.Model
	ClientId     *string `gorm:"unique;not null"`
	ClientSecret *string
	Token        *string
	RefreshToken *string
	Activities   []DBActivity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AthletId     *uint64      `gorm:"unique;not null"`
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
	Id                 int64   `gorm:"unique;not null" json:"id,omitempty"`
	StartDate          string  `gorm:"unique;not null" json:"start_date,omitempty"`
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

func (s *Storer) UpsertClient(clientId, clientSecret, token, refreshToken string, athletId uint64) error {
	log.Infof("Creating: %s %s %s %s", clientId, clientSecret, token, refreshToken)
	result := DB.Create(&DBClient{ClientId: &clientId, ClientSecret: &clientSecret, Token: &token, RefreshToken: &refreshToken, AthletId: &athletId})
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
	log.Infof("Creating: %s %s %s", req.GetName(), req.GetAthlete().GetId(), req.GetStartDate())
	athletId := req.GetAthlete().GetId()
	var dbClient DBClient
	result := DB.Where(&DBClient{AthletId: &athletId}).First(&dbClient)
	if result.Error != nil {
		log.Errorf("DB error: could not get Client %s: %s", athletId, result.Error)
		return fmt.Errorf("DB error: could not get Client %d: %s", athletId, result.Error)
	}
	err := DB.Model(&dbClient).Association("Activities").Append(activityToDBActivity(req))
	if isDuplicateKeyError(err) {
		var duplicate DBActivity
		DB.Where(&DBActivity{StartDate: req.GetStartDate()}).First(&duplicate)
		err = DB.Model(&dbClient).Association("Activities").Delete(&duplicate)
		err = DB.Model(&dbClient).Association("Activities").Append(activityToDBActivity(req))
	}
	if err != nil {
		log.Errorf("DB error: could not add activity %s", err.Error())
		return fmt.Errorf("DB error: could not add activity %s", err.Error())
	}
	log.Info("Added activity %s to client %s", req.GetName(), athletId)
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

func activityToDBActivity(activity *grpcStrava.Activity) *DBActivity {
	return &DBActivity{
		ResourceState:      activity.GetResourceState(),
		Name:               activity.GetName(),
		Distance:           activity.GetDistance(),
		MovingTime:         activity.GetMovingTime(),
		ElapsedTime:        activity.GetElapsedTime(),
		TotalElevationGain: activity.GetTotalElevationGain(),
		Type:               activity.GetType(),
		SportType:          activity.GetSportType(),
		StartDate:          activity.GetStartDate(),
		StartDateLocale:    activity.GetStartDateLocale(),
		Timezone:           activity.GetTimezone(),
		UtcOffset:          activity.GetUtcOffset(),
		LocationCity:       activity.GetLocationCity(),
		LocationState:      activity.GetLocationState(),
		LocationCountry:    activity.GetLocationCountry(),
		AchievementCount:   activity.GetAchievementCount(),
		KudosCount:         activity.GetKudosCount(),
		CommentCount:       activity.GetCommentCount(),
		Manual:             activity.GetManual(),
		Visibility:         activity.GetVisibility(),
		AverageSpeed:       activity.GetAverageSpeed(),
		MaxSpeed:           activity.GetMaxSpeed(),
		AverageHeartrate:   activity.GetAverageHeartrate(),
		MaxHeartrate:       activity.GetMaxHeartrate(),
		ElevHigh:           activity.GetElevLow(),
	}
}

func isDuplicateKeyError(err error) bool {
	pgErr, ok := err.(pgx.PgError)
	if ok {
		// unique_violation = 23505
		return pgErr.Code == "23505"

	}
	return false
}
