package worker

import (
	"context"

	"github.com/frain-dev/convoy"
	"github.com/frain-dev/convoy/datastore"
	"github.com/frain-dev/convoy/worker/task"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/taskq/v3"
)

func RegisterNewGroupTask(applicationRepo datastore.ApplicationRepository, eventDeliveryRepo datastore.EventDeliveryRepository, groupRepo datastore.GroupRepository) chan error {
	errChan := make(chan error)

	go func() {
		for {
			filter := &datastore.GroupFilter{}
			groups, err := groupRepo.LoadGroups(context.Background(), filter)
			if err != nil {
				errChan <- err
				return
			}
			for _, g := range groups {
				name := convoy.TaskName(g.Name)
				if t := taskq.Tasks.Get(string(name)); t == nil {
					handler := task.ProcessEventDelivery(applicationRepo, eventDeliveryRepo, groupRepo)
					log.Infof("Registering task handler for %s", g.Name)
					task.CreateTask(name, *g, handler)
				}
			}
		}
	}()
	return errChan
}
