package validateTransaction

import (
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/transaction"
	"context"
)

type task struct {
	ctx 			context.Context
	transaction 	transaction.Transaction
	ruleSets    	[]ruleSet.RuleSet
	response    	chan report.Report
	instrumentation *instrumentation
}

func newTask(
	t transaction.Transaction,
	ruleSets []ruleSet.RuleSet,
	response chan report.Report,
	logger *logger.Logger,
	ctx context.Context,
) task {
	instrumentation := newInstrumentation(logger)
	instrumentation.setContext(ctx)
	instrumentation.setMetadata(metadata{
		"amount": t.Amount,
		"organization": t.Organization,
	})

	return task{
		ctx: 		 	 ctx,
		transaction: 	 t,
		ruleSets:    	 ruleSets,
		response:    	 response,
		instrumentation: instrumentation,
	}
}

func (task *task) run() {
	task.instrumentation.startTransactionValidation()

	r := report.New()

	for _, rs := range task.ruleSets {
		switch rs.EvaluateTransaction(task.transaction) {
		case ruleSet.Block:
			r.AppendBlockRuleSet(rs)
		case ruleSet.Tag:
			r.AppendTagRuleSet(rs)
		}
	}

	task.instrumentation.endTransactionValidation()

	task.response <- r

	close(task.response)
}

type ValidatorService struct {
	ruleSetRepository ruleSet.Repository
	queue             chan task
	numOfWorkers      int
	workers           []chan bool
	logger			  *logger.Logger
}

func NewValidatorService(numOfWorkers int, ruleSetRepository ruleSet.Repository, logger *logger.Logger) ValidatorService {
	v := ValidatorService{
		ruleSetRepository: ruleSetRepository,
		queue:             make(chan task),
		numOfWorkers:      numOfWorkers,
		workers:           []chan bool{},
		logger:			   logger,
	}

	for workerId := 0; workerId < v.numOfWorkers; workerId++ {
		v.workers = append(v.workers, newWorker(v.queue, workerId))
	}

	return v
}

func (v *ValidatorService) Enqueue(trx transaction.Transaction, ctx context.Context) chan report.Report {
	ruleSets := v.ruleSetRepository.ListForOrganization(trx.Organization)
	response := make(chan report.Report, 1)

	v.queue <- newTask(trx, ruleSets, response, v.logger, ctx)

	return response
}

func (v *ValidatorService) ResizeWorkers(numOfWorkers int) {
	delta := v.numOfWorkers - numOfWorkers

	if delta > 0 {
		for _, worker := range v.workers[numOfWorkers:] {
			worker <- true
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

func newWorker(queue chan task, id int) chan bool {
	done := make(chan bool)

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
