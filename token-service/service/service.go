package service

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpcToken "github.com/JanMeckelholt/running/common/grpc/token"
	"github.com/JanMeckelholt/running/token-service/service/config"
)

var DB *gorm.DB

type DBClient struct {
	gorm.Model
	ClientId     string
	ClientSecret string
	Token        string
	RefreshToken string
}

type TokenStorer struct {
	StorerConfig config.StorerConfig
}

func NewStorer(storerConfig config.StorerConfig) *TokenStorer {
	return &TokenStorer{
		StorerConfig: storerConfig,
	}
}

func (s *TokenStorer) InitStorage() error {
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

func (s *TokenStorer) AutoMigrate(object interface{}) error {
	log.Info("Starting automatic migrations")

	var err error
	if err = DB.Debug().AutoMigrate(object); err != nil {
		log.Error("DB error: could not automigrate object")
		return err
	}
	log.Info("Automatic migrations finished")
	return nil
}

func (s *TokenStorer) UpsertClient(clientId, clientSecret, token, refreshToken string) error {
	log.Infof("Creating: %s %s %s %s", clientSecret, token, refreshToken, clientId)
	result := DB.Create(&DBClient{ClientId: clientId, ClientSecret: clientSecret, Token: token, RefreshToken: refreshToken})
	if result.Error != nil {
		log.Error("DB error: could not create object")
		return result.Error
	}
	log.Tracef("Stored object")
	return nil
}

func (s *TokenStorer) UpdateClient(clientId string, kvPairs []*grpcToken.KvPair) (*grpcToken.Client, error) {
	oldDBClient, err := s.GetDBClient(clientId)
	if err != nil {
		return nil, err
	}
	for _, kvPair := range kvPairs {
		log.Infof("Updating: %s %s %s ", clientId, kvPair.Key, kvPair.Value)
		result := DB.Model(oldDBClient).Update(kvPair.Key, kvPair.Value)
		if result.Error != nil {
			log.Errorf("DB error: could not update client: %s", result.Error)
			return nil, result.Error
		}
	}
	log.Tracef("Stored client: %s", oldDBClient)
	return dbClientToClient(oldDBClient), nil
}
func (s *TokenStorer) GetDBClient(clientId string) (*DBClient, error) {
	var (
		result   *gorm.DB
		dbClient *DBClient
	)
	result = DB.First(dbClient, clientId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Infof("No record found with id %s", dbClient)
		return nil, nil
	}
	if result.Error != nil {
		log.Errorf("DB error: could not get object by id %s: %s", clientId, result.Error)
		return nil, result.Error
	}
	log.Tracef("Retrieved object by id")
	return dbClient, nil
}

func dbClientToClient(athlete *DBClient) *grpcToken.Client {
	return &grpcToken.Client{
		ClientId:     athlete.ClientId,
		ClientSecret: athlete.ClientSecret,
		Token:        athlete.Token,
		RefreshToken: athlete.RefreshToken,
	}
}
