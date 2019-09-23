package validateTransaction

import (
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/transaction"
	"context"
)

type task struct {
	ctx               context.Context
	transaction       transaction.Transaction
	ruleSetRepository ruleSet.Repository
	response          chan report.Report
	error             chan error
	instrumentation *instrumentation
}

func newTask(
	ctx context.Context,
	trx transaction.Transaction,
	ruleSetRepository ruleSet.Repository,
	logger *logger.Logger,
) task {
	instrumentation := newInstrumentation(logger)
	instrumentation.setContext(ctx)
	instrumentation.setMetadata(metadata{
		"amount": t.Amount,
		"entity": t.Entity,
	})

	return task{
		ctx:               ctx,
		transaction:       trx,
		ruleSetRepository: ruleSetRepository,
		response:          make(chan report.Report),
		error:             make(chan error),
		instrumentation: instrumentation,
	}
}

func (task *task) run() {
	task.instrumentation.startTransactionValidation()

	defer close(task.response)
	defer close(task.error)

	ruleSets, err := task.ruleSetRepository.ListByEntityId(task.ctx, task.transaction.EntityId)

	if err != nil {
		task.error <- err
		return
	}

	r := report.New()

	for _, rs := range ruleSets {
		action, err := rs.Matches(task.transaction)
		if err != nil {
			task.error <- err
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
	ruleSetRepository ruleSet.Repository
	queue             chan task
	numOfWorkers      int
	workers           []chan struct{}
	logger            *logger.Logger
}

func NewValidatorService(numOfWorkers int, ruleSetRepository ruleSet.Repository, logger *logger.Logger) ValidatorService {
	v := ValidatorService{
		ruleSetRepository: ruleSetRepository,
		queue:             make(chan task),
		numOfWorkers:      numOfWorkers,
		workers:           []chan struct{}{},
		logger:            logger,
	}

	for workerId := 0; workerId < v.numOfWorkers; workerId++ {
		v.workers = append(v.workers, newWorker(v.queue, workerId))
	}

	return v
}

func (v *ValidatorService) Enqueue(ctx context.Context, trx transaction.Transaction) (chan report.Report, chan error) {
	task := newTask(ctx, trx, v.ruleSetRepository, v.logger)
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
