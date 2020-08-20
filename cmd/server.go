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
	appd "bitbucket.verifone.com/validation-service/appdynamics"
	"bitbucket.verifone.com/validation-service/config"
	"bitbucket.verifone.com/validation-service/http"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"github.com/go-chi/chi"
	"log"
	"os"
	"runtime"
)

func StartServer(config config.Server) error {
	log := setupLogger(config.Log)

	setupAppD(config.AppD)

	ruleSetRepo := createRuleSetRepository(config.Mongo, log)

	validateTransactionApp := createValidateTransactionApp(ruleSetRepo, log)

	log.Output.Infof("Starting REST API server at port %d", config.HTTPPort)

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
		return listAncestorsRuleSet.NewListAncestorsRuleSet(log, ruleSetRepo)
	}

	listDescendantsRuleSetAppFactory := func() listDescendantsRuleSet.ListDescendantsRuleSet {
		return listDescendantsRuleSet.NewListDescendantsRuleSet(log, ruleSetRepo)
	}

	updateRuleSetAppFactory := func() updateRuleSet.UpdateRuleSet {
		return updateRuleSet.NewUpdateRuleSet(log, ruleSetRepo)
	}

	err := http.NewServer(
		config.HTTPPort,
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
		updateRuleSetAppFactory,
	).Start()

	if err != nil {
		log.Error.WithError(err).Error("Failed to start REST API server")
		os.Exit(1)
	}

	return nil
}

func setupLogger(logConfig config.Log) *logger.Logger {
	l, err := logger.NewLogger(
		config.AppName,
		config.Version,
		logConfig.FormatValue(),
		logConfig.LevelValue(),
	)

	if err != nil {
		log.Panic("Failed to initialize logger")
	}

	return l
}

func setupAppD(appDConfig config.AppD) {
	cfg := appd.Config{}

	cfg.AppName = appDConfig.AppName
	cfg.TierName = appDConfig.TierName
	cfg.NodeName = appDConfig.NodeName
	cfg.InitTimeoutMs = appDConfig.InitTimeout
	cfg.Controller.Host = appDConfig.Controller.Host
	cfg.Controller.Port = appDConfig.Controller.Port
	cfg.Controller.UseSSL = appDConfig.Controller.UseSSL
	cfg.Controller.Account = appDConfig.Controller.Account
	cfg.Controller.AccessKey = appDConfig.Controller.AccessKey

	if err := appd.InitSDK(&cfg); err != nil {
		log.Panic("Error initializing the AppDynamics SDK\n")
	}
}

func createRuleSetRepository(mongoConfig config.Mongo, logger *logger.Logger) *ruleSet.MongoRuleSetRepository {
	ruleSetRepository, err := ruleSet.NewMongoRepository(mongoConfig.URL, mongoConfig.DB, logger)

	if err != nil {
		logger.Error.WithError(err).Error("Failed to initialize RuleSetRepository")
		os.Exit(1)
	}

	return ruleSetRepository
}

func createValidateTransactionApp(
	r *ruleSet.MongoRuleSetRepository,
	l *logger.Logger,
) validateTransaction.ValidatorService {
	validator := validateTransaction.NewValidatorService(runtime.NumCPU(), r, l)
	return &validator
}
