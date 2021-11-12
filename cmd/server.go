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

	rec := &logger.LogRecord{}

	setupAppD(config.AppD)

	ruleSetRepo := createRuleSetRepository(config.Mongo, log)

	validateTransactionApp := createValidateTransactionApp(ruleSetRepo, log, rec)

	log.Output.Infof("Starting REST API server at port %d", config.HTTPPort)

	createRuleSetAppFactory := func() createRuleSet.CreateRuleSet {
		var rec *logger.LogRecord
		newRec := rec.NewRecord()
		return createRuleSet.NewCreateRuleSet(log, newRec, ruleSetRepo)
	}

	getRuleSetAppFactory := func() getRuleSet.GetRuleSet {
		var rec *logger.LogRecord
		newRec := rec.NewRecord()
		return getRuleSet.NewGetRuleSet(log, newRec, ruleSetRepo)
	}

	deleteRuleSetAppFactory := func() deleteRuleSet.DeleteRuleSet {
		var rec *logger.LogRecord
		newRec := rec.NewRecord()
		return deleteRuleSet.NewDeleteRuleSet(log, newRec, ruleSetRepo)
	}

	listRuleSetAppFactory := func() listRuleSet.ListRuleSet {
		var rec *logger.LogRecord
		newRec := rec.NewRecord()
		return listRuleSet.NewListRuleSet(log, newRec, ruleSetRepo)
	}

	listAncestorsRuleSetAppFactory := func() listAncestorsRuleSet.ListAncestorsRuleSet {
		var rec *logger.LogRecord
		newRec := rec.NewRecord()
		return listAncestorsRuleSet.NewListAncestorsRuleSet(log, newRec, ruleSetRepo)
	}

	listDescendantsRuleSetAppFactory := func() listDescendantsRuleSet.ListDescendantsRuleSet {
		var rec *logger.LogRecord
		newRec := rec.NewRecord()
		return listDescendantsRuleSet.NewListDescendantsRuleSet(log, newRec, ruleSetRepo)
	}

	updateRuleSetAppFactory := func() updateRuleSet.UpdateRuleSet {
		var rec *logger.LogRecord
		newRec := rec.NewRecord()
		return updateRuleSet.NewUpdateRuleSet(log, newRec, ruleSetRepo)
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
		logConfig.Format,
		logConfig.LevelValue(),
		logConfig.LogFile,
		logConfig.LogFileMaxMb,
		logConfig.LogFileRotationCount,
		logConfig.LogFileRotationDays,
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
	if proxyHost := appDConfig.Controller.ProxyHost; proxyHost != "" {
		cfg.Controller.ProxyHost = appDConfig.Controller.ProxyHost
	}
	if proxyPort := appDConfig.Controller.ProxyPort; proxyPort != "" {
		cfg.Controller.ProxyPort = appDConfig.Controller.ProxyPort
	}

	if err := appd.InitSDK(&cfg); err != nil {
		log.Panic("Error initializing the AppDynamics SDK\n", err)
	}

	backendProperties := map[string]string{
		"DATABASE": "mongodb",
	}

	if err := appd.AddBackend(string(appdBackend.MongoDB), "DB", backendProperties, true); err != nil {
		log.Printf("Failed to add AppD database backend: %s", err.Error())
	}
}

func createRuleSetRepository(mongoConfig config.Mongo, logger *logger.Logger) *ruleSet.MongoRuleSetRepository {
	mongoRetryDelay := time.Duration(mongoConfig.RetryMilliseconds) * time.Millisecond

	ruleSetRepository, err := ruleSet.NewMongoRepository(
		mongoConfig.URL,
		mongoConfig.DB,
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
	rec *logger.LogRecord,
) validateTransaction.ValidatorService {
	validator := validateTransaction.NewValidatorService(runtime.NumCPU(), r, l, rec)
	return &validator
}
