package validateTransaction

import (
	"context"
	"validation-service/logger"
	"validation-service/report"
	"validation-service/ruleSet"
	"validation-service/transaction"
)

type ValidatorService interface {
	Enqueue(ctx context.Context, trx transaction.Transaction) (chan report.Report, chan AppError)
	ResizeWorkers(numOfWorkers int)
}

const appName = "validateTransaction"

type task struct {
	ctx               context.Context
	transaction       transaction.Transaction
	ruleSetRepository ruleSet.Repository
	response          chan report.Report
	error             chan AppError
	instrumentation   *instrumentation
}

func newTask(
	ctx context.Context,
	trx transaction.Transaction,
	ruleSetRepository ruleSet.Repository,
	logger *logger.Logger,
	record *logger.LogRecord,
) task {
	instrumentation := newInstrumentation(logger, record)
	instrumentation.setContext(ctx)
	instrumentation.setMetadata(metadata{
		"amount": trx.Amount,
		"entity": trx.EntityId,
	})

	return task{
		ctx:               ctx,
		transaction:       trx,
		ruleSetRepository: ruleSetRepository,
		response:          make(chan report.Report),
		error:             make(chan AppError),
		instrumentation:   instrumentation,
	}
}

func (task *task) run() {
	task.instrumentation.startTransactionValidation()

	defer close(task.response)
	defer close(task.error)

	entityIds := []string{task.transaction.EntityId} // TODO: replace with list of entity ids coming in the request

	ruleSets, err := task.ruleSetRepository.ListByEntityIds(task.ctx, entityIds...)

	if err != nil {
		task.error <- NewError(UnexpectedErr, err)
		return
	}

	r := report.New()

	for _, rs := range ruleSets {
		action, err := rs.Matches(task.transaction)
		if err != nil {
			task.error <- NewError(UnexpectedErr, err)
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

type App struct {
	ruleSetRepository ruleSet.Repository
	queue             chan task
	numOfWorkers      int
	workers           []chan struct{}
	logger            *logger.Logger
	record            *logger.LogRecord
}

func NewValidatorService(
	numOfWorkers int,
	ruleSetRepository ruleSet.Repository,
	logger *logger.Logger,
	record *logger.LogRecord,
) App {
	v := App{
		ruleSetRepository: ruleSetRepository,
		queue:             make(chan task),
		numOfWorkers:      numOfWorkers,
		workers:           []chan struct{}{},
		logger:            logger,
		record:            record.Scoped(appName),
	}

	for workerId := 0; workerId < v.numOfWorkers; workerId++ {
		v.workers = append(v.workers, newWorker(v.queue, workerId))
	}

	return v
}

func (v *App) Enqueue(ctx context.Context, trx transaction.Transaction) (chan report.Report, chan AppError) {
	task := newTask(ctx, trx, v.ruleSetRepository, v.logger, v.record)
	v.queue <- task
	return task.response, task.error
}

func (v *App) ResizeWorkers(numOfWorkers int) {
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
