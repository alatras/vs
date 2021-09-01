package cmd

import (
	"log"
	"os"
	"runtime"
	"time"

	"validation-service/app/createRuleSet"
	"validation-service/app/deleteRuleSet"
	"validation-service/app/getRuleSet"
	"validation-service/app/listAncestorsRuleSet"
	"validation-service/app/listDescendantsRuleSet"
	"validation-service/app/listRuleSet"
	"validation-service/app/updateRuleSet"
	"validation-service/app/validateTransaction"
	appd "validation-service/appdynamics"
	"validation-service/config"
	"validation-service/enums/appdBackend"
	"validation-service/http"
	"validation-service/logger"
	"validation-service/ruleSet"

	"github.com/go-chi/chi"
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
		logConfig.LogFileValue(),
		logConfig.LogFileMaxMbValue(),
		logConfig.LogFileRotationCountValue(),
		logConfig.LogFileRotationDaysValue(),
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
	if proxyHost := appDConfig.GetConfig("APP_DYNAMICS_PROXY_HOST"); proxyHost != "" {
		cfg.Controller.ProxyHost = appDConfig.Controller.ProxyHost
	}
	if proxyPort := appDConfig.GetConfig("APP_DYNAMICS_PROXY_HOST"); proxyPort != "" {
		cfg.Controller.ProxyPort = appDConfig.Controller.ProxyPort
	}

	if err := appd.InitSDK(&cfg); err != nil {
		log.Panic("Error initializing the AppDynamics SDK\n")
	}

	backendProperties := map[string]string{
		"DATABASE": "mongodb",
	}

	if err := appd.AddBackend(string(appdBackend.MongoDB), "DB", backendProperties, true); err != nil {
		log.Printf("Failed to add AppD database backend: %s", err.Error())
	}
}

func createRuleSetRepository(mongoConfig config.Mongo, logger *logger.Logger) *ruleSet.MongoRuleSetRepository {
	var mongoRetryDelay time.Duration

	if mongoConfig.RetryMilliseconds != 0 {
		mongoRetryDelay = time.Duration(mongoConfig.RetryMilliseconds) * time.Millisecond
	} else {
		mongoRetryDelay = time.Duration(config.DefaultMongoRetryMilliseconds) * time.Millisecond
	}

	ruleSetRepository, err := ruleSet.NewMongoRepository(
		mongoConfig.GetConfig("MONGO_URL"),
		mongoConfig.GetConfig("MONGO_DB"),
		mongoRetryDelay,
		logger,
	)

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
