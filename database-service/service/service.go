package service

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpcDB "github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/database-service/service/config"
)

var DB *gorm.DB

type DBClient struct {
	gorm.Model
	ClientId     *string `gorm:"unique;not null"`
	ClientSecret *string
	Token        *string
	RefreshToken *string
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
