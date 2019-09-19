package app

import (
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/transaction"
)

type task struct {
	transaction transaction.Transaction
	ruleSets    []ruleSet.RuleSet
	response    chan report.Report
}

func newTask(t transaction.Transaction, ruleSets []ruleSet.RuleSet, response chan report.Report) task {
	return task{
		transaction: t,
		ruleSets:    ruleSets,
		response:    response,
	}
}

func (task *task) run() {
	r := report.New()

	for _, rs := range task.ruleSets {
		switch rs.EvaluateTransaction(task.transaction) {
		case ruleSet.Block:
			r.AppendBlockRuleSet(rs)
		case ruleSet.Tag:
			r.AppendTagRuleSet(rs)
		}
	}

	task.response <- r
	close(task.response)
}

type ValidatorService struct {
	ruleSetRepository ruleSet.Repository
	queue             chan task
	numOfWorkers      int
	workers           []chan bool
}

func NewValidatorService(numOfWorkers int, ruleSetRepository ruleSet.Repository) *ValidatorService {
	v := &ValidatorService{
		ruleSetRepository: ruleSetRepository,
		queue:             make(chan task),
		numOfWorkers:      numOfWorkers,
		workers:           []chan bool{},
	}

	for workerId := 0; workerId < v.numOfWorkers; workerId++ {
		v.workers = append(v.workers, newWorker(v.queue, workerId))
	}

	return v
}

func (v *ValidatorService) Enqueue(trx transaction.Transaction) chan report.Report {
	ruleSets := v.ruleSetRepository.ListForOrganization(trx.Organization)
	response := make(chan report.Report, 1)

	v.queue <- newTask(trx, ruleSets, response)

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
