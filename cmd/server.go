package cmd

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/app/listAncestorsRuleSet"
	"bitbucket.verifone.com/validation-service/app/listDescendantsRuleSet"
	"bitbucket.verifone.com/validation-service/app/listRuleSet"
	"bitbucket.verifone.com/validation-service/app/updateRuleSet"
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/entityService"
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
	Mongo         MongoGroup         `group:"mongo" namespace:"mongo"`
	EntityService EntityServiceGroup `group:"entityService" namespace:"entityService"`

	HTTPPort int `long:"httpPort" env:"HTTP_PORT" default:"8080" description:"HTTP port"`

	CommonOpts
}

// MongoGroup MongoDB configuration parameters
type MongoGroup struct {
	URL string `long:"url" env:"MONGO_URL" required:"MongoDB url required" description:"MongoDB url"`
	DB  string `long:"db" env:"MONGO_DB" default:"validationService" description:"MongoDB database name"`
}

// EntityServiceGroup Entity Service configuration parameters
type EntityServiceGroup struct {
	URL                  string `long:"url" env:"ENTITY_SERVICE_URL" required:"Entity service url required" description:"Entity service URL (without trailing slash)"`
	JWTToken             string `long:"jwtToken" env:"ENTITY_SERVICE_JWT_TOKEN" required:"Entity service JWT auth token required" description:"Entity service JWT auth token"`
	SkipCertVerification bool   `long:"skipCertVerification" env:"ENTITY_SERVICE_SKIP_CERT_VERIFICATION" description:"Skip HTTPS certificate verification or not"`
}

// Execute is the entry point for "server" command
func (s *ServerCommand) Execute(args []string) error {
	log := s.setupLogger()

	ruleSetRepo := s.createRuleSetRepository(log)

	entityServiceClient := entityService.NewClient(
		log, s.EntityService.URL,
		s.EntityService.JWTToken,
		s.EntityService.SkipCertVerification,
	)

	validateTransactionApp := s.createValidateTransactionApp(entityServiceClient, ruleSetRepo, log)

	log.Output.Infof("Starting REST API server at port %d", s.HTTPPort)

	createRuleSetAppFactory := func() createRuleSet.CreateRuleSet {
		return createRuleSet.NewCreateRuleSet(log, ruleSetRepo)
	}

	getRuleSetAppFactory := func() getRuleSet.GetRuleSet {
		return getRuleSet.NewGetRuleSet(log, ruleSetRepo)
	}

	deleteRuleSetAppFactory := func() deleteRuleSet.DeleteRuleSet {
		return deleteRuleSet.NewDeleteRuleSet(log, ruleSetRepo)
	}

	listRuleSetAppFactory := func() listRuleSet.ListRuleSet {
		return listRuleSet.NewListRuleSet(log, ruleSetRepo)
	}

	listAncestorsRuleSetAppFactory := func() listAncestorsRuleSet.ListAncestorsRuleSet {
		return listAncestorsRuleSet.NewListAncestorsRuleSet(log, ruleSetRepo, entityServiceClient)
	}

	listDescendantsRuleSetAppFactory := func() listDescendantsRuleSet.ListDescendantsRuleSet {
		return listDescendantsRuleSet.NewListDescendantsRuleSet(log, ruleSetRepo, entityServiceClient)
	}

	updateRuleSetAppFactory := func() updateRuleSet.UpdateRuleSet {
		return updateRuleSet.NewUpdateRuleSet(log, ruleSetRepo)
	}

	err := http.NewServer(
		s.HTTPPort,
		chi.NewRouter(),
		log,
		ruleSetRepo,
		validateTransactionApp,
		createRuleSetAppFactory,
		listRuleSetAppFactory,
		listAncestorsRuleSetAppFactory,
		listDescendantsRuleSetAppFactory,
		getRuleSetAppFactory,
		deleteRuleSetAppFactory,
		entityServiceClient,
		updateRuleSetAppFactory,
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
	ruleSetRepository, err := ruleSet.NewMongoRepository(s.Mongo.URL, s.Mongo.DB, logger)

	if err != nil {
		logger.Error.WithError(err).Error("Failed to initialize RuleSetRepository")
		os.Exit(1)
	}

	return ruleSetRepository
}

func (s *ServerCommand) createValidateTransactionApp(
	e entityService.EntityService,
	r *ruleSet.MongoRuleSetRepository,
	l *logger.Logger,
) validateTransaction.ValidatorService {
	validator := validateTransaction.NewValidatorService(runtime.NumCPU(), e, r, l)
	return &validator
}
