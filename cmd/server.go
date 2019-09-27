package cmd

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/http"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"github.com/go-chi/chi"
	"log"
	"os"
	"runtime"
)

// ServerCommand with command line flags and env
type ServerCommand struct {
	Mongo MongoGroup `group:"mongo" namespace:"mongo" env-namespace:"MONGO"`

	HTTPPort int `long:"httpPort" env:"HTTP_PORT" default:"8080" description:"HTTP port"`

	CommonOpts
}

// MongoGroup MongoDB configuration parameters
type MongoGroup struct {
	URL string `long:"url" env:"MONGO_URL" required:"MongoDB url required" description:"MongoDB url"`
	DB  string `long:"db" env:"MONGO_DB" default:"validationService" description:"MongoDB database name"`
}

// Execute is the entry point for "server" command
func (s *ServerCommand) Execute(args []string) error {
	log := s.setupLogger()

	ruleSetRepo := s.createRuleSetRepository(log)

	validateTransactionApp := s.createValidateTransactionApp(ruleSetRepo, log)

	log.Output.Infof("Starting REST API server at port %d", s.HTTPPort)

	createRuleSetAppFactory := func() createRuleSet.CreateRuleset {
		return createRuleSet.NewCreateRuleset(log, ruleSetRepo)
	}

	getRuleSetAppFactory := func() getRuleSet.GetRuleSet {
		return getRuleSet.NewGetRuleSet(log, ruleSetRepo)
	}

	err := http.NewServer(
		s.HTTPPort,
		chi.NewRouter(),
		log,
		ruleSetRepo,
		validateTransactionApp,
		createRuleSetAppFactory,
		getRuleSetAppFactory,
	).Start()

	if err != nil {
		log.Error.WithError(err).Error("Failed to start REST API server")
		os.Exit(1)
	}

	return nil
}

func (s *ServerCommand) setupLogger() *logger.Logger {
	l, err := logger.NewLogger(
		AppName,
		Version,
		s.Log.FormatValue(),
		s.Log.LevelValue(),
	)

	if err != nil {
		log.Panic("Failed to initialize logger")
	}
	return l
}

func (s *ServerCommand) createRuleSetRepository(logger *logger.Logger) *ruleSet.MongoRuleSetRepository {
	ruleSetRepository, err := ruleSet.NewMongoRepository(s.Mongo.URL, s.Mongo.DB)

	if err != nil {
		logger.Error.WithError(err).Error("Failed to initialize RuleSetRepository")
		os.Exit(1)
	}

	return ruleSetRepository
}

func (s *ServerCommand) createValidateTransactionApp(r *ruleSet.MongoRuleSetRepository, l *logger.Logger) *validateTransaction.ValidatorService {
	validator := validateTransaction.NewValidatorService(runtime.NumCPU(), r, l)
	return &validator
}
