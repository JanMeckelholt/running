package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
	grpcStrava "github.com/JanMeckelholt/running/common/grpc/strava"
	"github.com/JanMeckelholt/running/database-service/service/config"
)

var DB *gorm.DB

type DBClient struct {
	ClientId     *string `gorm:"primaryKey;not null"`
	ClientSecret *string
	Token        *string
	RefreshToken *string
	Activities   []*DBActivity `gorm:"foreignKey:ClientId;references:ClientId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AthletId     *uint64       `gorm:"unique;not null"`
}

type DBActivity struct {
	Id                 *uint64  `gorm:"primaryKey" json:"id"`
	ClientId           *string  `gorm:"not null"`
	Name               *string  `json:"name,omitempty"`
	Distance           *float64 `json:"distance,omitempty"`
	MovingTime         *uint64  `json:"moving_time,omitempty"`
	ElapsedTime        *uint64  `json:"elapsed_time,omitempty"`
	TotalElevationGain *float64 `json:"total_elevation_gain,omitempty"`
	Type               *string  `json:"type,omitempty"`
	SportType          *string  `json:"sport_type,omitempty"`
	StartDate          *string  `json:"start_date,omitempty"`
	StartDateLocale    *string  `json:"start_date_locale,omitempty"`
	StartDateUnix      *uint64  `gorm:"unique;not null" json:"start_date_unix,omitempty"`
	Timezone           *string  `json:"timezone,omitempty"`
	UtcOffset          *float64 `json:"utc_offset,omitempty"`
	LocationCity       *string  `json:"location_city,omitempty"`
	LocationState      *string  `json:"location_state,omitempty"`
	LocationCountry    *string  `json:"location_country,omitempty"`
	AchievementCount   *uint64  `json:"achievement_count,omitempty"`
	KudosCount         *uint64  `json:"kudos_count,omitempty"`
	CommentCount       *uint64  `json:"comment_count,omitempty"`
	Manual             *bool    `json:"manual,omitempty"`
	Visibility         *string  `json:"visibility,omitempty"`
	AverageSpeed       *float64 `json:"average_speed,omitempty"`
	MaxSpeed           *float64 `json:"max_speed,omitempty"`
	AverageHeartrate   *float64 `json:"average_heartrate,omitempty"`
	MaxHeartrate       *float64 `json:"max_heartrate,omitempty"`
	ElevHigh           *float64 `json:"elev_high,omitempty"`
	ElevLow            *float64 `json:"elev_low,omitempty"`
	ResourceState      *uint64  `json:"resource_state,omitempty"`
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
	log.Infof("Using ConnString for DB: host=%s port=%d dbname=%s user=%s password=<XXX> sslmode=%s sslrootcert=volumes-data/certs/ca-cert.pem sslkey=secret/certs/postgres-key.pem sslcert=volumes-data/certs/postgres-cert.pem", s.StorerConfig.PostgresHost, s.StorerConfig.PostgresPort, s.StorerConfig.PostgresDB, s.StorerConfig.PostgresUser, s.StorerConfig.PostgresSSLMode)
	var err error
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s sslrootcert=volumes-data/certs/ca-cert.pem sslkey=secret/certs/postgres-key.pem sslcert=volumes-data/certs/postgres-cert.pem", s.StorerConfig.PostgresHost, s.StorerConfig.PostgresPort, s.StorerConfig.PostgresDB, s.StorerConfig.PostgresUser, s.StorerConfig.PostgresPassword, s.StorerConfig.PostgresSSLMode)), &gorm.Config{})
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
	log.Infof("Stored client: %s", result.RowsAffected)
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
	log.Infof("Stored client: %s", oldDBClient)
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
	log.Infof("Retrieved client %s, %d, %s", *dbClient.ClientId, dbClient.AthletId, *dbClient.Token)
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

func (s *Storer) UpsertActivity(req *grpcStrava.Activity, fromCSV bool) error {
	log.Infof("Creating: %s %d %s", req.GetName(), req.GetAthlete().GetId(), req.GetStartDate())
	athletId := req.GetAthlete().GetId()
	var dbClient DBClient
	result := DB.Where(&DBClient{AthletId: &athletId}).First(&dbClient)
	if result.Error != nil {
		log.Errorf("DB error: could not get Client %s: %s", athletId, result.Error)
		return fmt.Errorf("DB error: could not get Client %d: %s", athletId, result.Error)
	}
	dbActivity := activityToDBActivity(req, *dbClient.ClientId, fromCSV)
	err := DB.Model(&dbClient).Association("Activities").Append(dbActivity)
	if isDuplicateKeyError(err) {
		var duplicates []DBActivity
		startDateUnix := dbActivity.StartDateUnix
		_ = DB.Where(&DBActivity{ClientId: dbClient.ClientId, StartDateUnix: startDateUnix}).Find(&duplicates)
		log.Infof("found %d duplicates. Deleting..", len(duplicates))
		result = DB.Unscoped().Delete(&duplicates)
		if result.Error != nil {
			log.Errorf("Could not delete duplicates: %s", result.Error)
			return err
		}
		log.Infof("Deleted %d", result.RowsAffected)
		result = DB.Create(&dbActivity)
		err = result.Error
	}
	if err != nil {
		log.Infof("Error: %s", err)
		log.Errorf("DB error: could not add activity: %s", err.Error())
		return fmt.Errorf("DB error: could not add activity %s", err.Error())
	}
	log.Infof("Added activity %s to athlet %d", req.GetName(), athletId)
	return nil
}

func (s *Storer) GetActivities(req *grpcDB.ActivitiesRequest) (*grpcDB.ActivitiesResponse, error) {

	var (
		dbActivities []*DBActivity
		activities   []*grpcStrava.Activity
	)
	until := req.GetUntil()
	if until == 0 {
		until = uint64(time.Now().Unix())
	}
	result := DB.Where("client_id = ? AND start_date_unix > ?  AND start_date_unix < ?", req.GetClientId(), req.GetSince(), until).Find(&dbActivities)
	if result.Error != nil {
		log.Errorf("DB error: could not get activities %s_%d: %s", req.GetClientId(), req.GetSince(), result.Error)
		return nil, fmt.Errorf("DB error: could not get activities %s_%d: %s", req.GetClientId(), req.GetSince(), result.Error)
	}
	for _, a := range dbActivities {
		sA := dbActivityToActivity(a)
		activities = append(activities, sA)
	}

	log.Infof("Got %d activities for %s since %d", len(dbActivities), req.GetClientId(), req.GetSince())
	return &grpcDB.ActivitiesResponse{Activities: activities}, nil
}

func dbActivityToActivity(dbActivity *DBActivity) (res *grpcStrava.Activity) {
	b, err := json.Marshal(dbActivity)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal(b, &res)
	return res
}

func ptr[T any](x T) *T {
	return &x
}

func activityToDBActivity(activity *grpcStrava.Activity, clientId string, fromCSV bool) *DBActivity {
	var (
		startDateUnix uint64
		layout        string
	)
	layoutCSV := "02.01.2006, 15:04:05"
	layoutStravaAPI := "2006-01-02T15:04:05Z"

	if fromCSV {
		layout = layoutCSV
	} else {
		layout = layoutStravaAPI
	}
	t, err := time.Parse(layout, activity.GetStartDate())
	if err != nil {
		log.Errorf("could nor parse %s", activity.GetStartDate())
	}
	startDateUnix = uint64(t.Unix())

	return &DBActivity{
		ClientId:           &clientId,
		ResourceState:      ptr(activity.GetResourceState()),
		Name:               ptr(activity.GetName()),
		Distance:           ptr(activity.GetDistance()),
		MovingTime:         ptr(activity.GetMovingTime()),
		ElapsedTime:        ptr(activity.GetElapsedTime()),
		TotalElevationGain: ptr(activity.GetTotalElevationGain()),
		Type:               ptr(activity.GetType()),
		SportType:          ptr(activity.GetSportType()),
		StartDate:          ptr(activity.GetStartDate()),
		StartDateUnix:      &startDateUnix,
		StartDateLocale:    ptr(activity.GetStartDateLocale()),
		Timezone:           ptr(activity.GetTimezone()),
		UtcOffset:          ptr(activity.GetUtcOffset()),
		LocationCity:       ptr(activity.GetLocationCity()),
		LocationState:      ptr(activity.GetLocationState()),
		LocationCountry:    ptr(activity.GetLocationCountry()),
		AchievementCount:   ptr(activity.GetAchievementCount()),
		KudosCount:         ptr(activity.GetKudosCount()),
		CommentCount:       ptr(activity.GetCommentCount()),
		Manual:             ptr(activity.GetManual()),
		Visibility:         ptr(activity.GetVisibility()),
		AverageSpeed:       ptr(activity.GetAverageSpeed()),
		MaxSpeed:           ptr(activity.GetMaxSpeed()),
		AverageHeartrate:   ptr(activity.GetAverageHeartrate()),
		MaxHeartrate:       ptr(activity.GetMaxHeartrate()),
		ElevHigh:           ptr(activity.GetElevLow()),
	}
}

func isDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	if err.Error() == "duplicated key not allowed" {
		return true
	}
	return false
}
