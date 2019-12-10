package validateTransaction

import (
	"bitbucket.verifone.com/validation-service/entityService"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/transaction"
	"context"
)

const appName = "validateTransaction"

type task struct {
	ctx                 context.Context
	transaction         transaction.Transaction
	entityServiceClient entityService.EntityService
	ruleSetRepository   ruleSet.Repository
	response            chan report.Report
	error               chan ValidationError
	instrumentation     *instrumentation
}

func newTask(
	ctx context.Context,
	trx transaction.Transaction,
	entityServiceClient entityService.EntityService,
	ruleSetRepository ruleSet.Repository,
	logger *logger.Logger,
) task {
	instrumentation := newInstrumentation(logger)
	instrumentation.setContext(ctx)
	instrumentation.setMetadata(metadata{
		"amount": trx.Amount,
		"entity": trx.EntityId,
	})

	return task{
		ctx:                 ctx,
		transaction:         trx,
		entityServiceClient: entityServiceClient,
		ruleSetRepository:   ruleSetRepository,
		response:            make(chan report.Report),
		error:               make(chan ValidationError),
		instrumentation:     instrumentation,
	}
}

func (task *task) run() {
	task.instrumentation.startTransactionValidation()

	defer close(task.response)
	defer close(task.error)

	entityIds, err := task.entityServiceClient.GetAncestorsOf(task.transaction.EntityId)

	if err != nil {
		var validationError ValidationError

		if err == entityService.EntityNotFound {
			validationError = newError(EntityIdNotFoundErr, err)
		} else {
			validationError = newError(UnexpectedErr, err)
		}

		task.error <- validationError
		return
	}

	ruleSets, err := task.ruleSetRepository.ListByEntityIds(task.ctx, entityIds...)

	if err != nil {
		task.error <- newError(UnexpectedErr, err)
		return
	}

	r := report.New()

	for _, rs := range ruleSets {
		action, err := rs.Matches(task.transaction)
		if err != nil {
			task.error <- newError(UnexpectedErr, err)
			return
		}

		switch action {
		case ruleSet.Block:
			r.AppendBlockRuleSet(rs)
		case ruleSet.Tag:
			r.AppendTagRuleSet(rs)
		}
	}

	task.instrumentation.endTransactionValidation()

	task.response <- r
}

type ValidatorService struct {
	entityServiceClient entityService.EntityService
	ruleSetRepository   ruleSet.Repository
	queue               chan task
	numOfWorkers        int
	workers             []chan struct{}
	logger              *logger.Logger
}

func NewValidatorService(
	numOfWorkers int,
	entityServiceClient entityService.EntityService,
	ruleSetRepository ruleSet.Repository,
	logger *logger.Logger,
) ValidatorService {
	v := ValidatorService{
		entityServiceClient: entityServiceClient,
		ruleSetRepository:   ruleSetRepository,
		queue:               make(chan task),
		numOfWorkers:        numOfWorkers,
		workers:             []chan struct{}{},
		logger:              logger.Scoped(appName),
	}

	for workerId := 0; workerId < v.numOfWorkers; workerId++ {
		v.workers = append(v.workers, newWorker(v.queue, workerId))
	}

	return v
}

func (v *ValidatorService) Enqueue(ctx context.Context, trx transaction.Transaction) (chan report.Report, chan ValidationError) {
	task := newTask(ctx, trx, v.entityServiceClient, v.ruleSetRepository, v.logger)
	v.queue <- task
	return task.response, task.error
}

func (v *ValidatorService) ResizeWorkers(numOfWorkers int) {
	delta := v.numOfWorkers - numOfWorkers

	if delta > 0 {
		for _, worker := range v.workers[numOfWorkers:] {
			worker <- struct{}{}
		}
	} else if delta < 0 {
		workerId := v.numOfWorkers

		for delta != 0 {
			v.workers = append(v.workers, newWorker(v.queue, workerId))
			workerId++
			delta++
		}
	}

	v.numOfWorkers = numOfWorkers
}

func newWorker(queue chan task, id int) chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case task := <-queue:
				task.run()
			case <-done:
				return
			}
		}
	}()

	return done
}
