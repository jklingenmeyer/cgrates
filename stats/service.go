/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/
package stats

import (
	"errors"
	"fmt"
	"sync"

	"github.com/cgrates/cgrates/engine"
	"github.com/cgrates/cgrates/utils"
)

// NewStatService initializes a StatService
func NewStatService(dataDB engine.DataDB) (ss *StatService, err error) {
	ss = &StatService{dataDB: dataDB}
	return
}

// StatService builds stats for events
type StatService struct {
	sync.RWMutex
	dataDB   engine.DataDB
	stQueues StatsInstances   // ordered list of StatsQueues
	evCache  *StatsEventCache // so we can pass it to queues
}

// Called to start the service
func (ss *StatService) ListenAndServe() error {
	return nil
}

// Called to shutdown the service
func (ss *StatService) ServiceShutdown() error {
	return nil
}

// processEvent processes a StatsEvent through the queues and caches it when needed
func (ss *StatService) processEvent(ev engine.StatsEvent) (err error) {
	evStatsID := ev.ID()
	if evStatsID == "" { // ID is mandatory
		return errors.New("missing ID field")
	}
	for _, stQ := range ss.stQueues {
		if err := stQ.ProcessEvent(ev); err != nil {
			utils.Logger.Warning(fmt.Sprintf("<StatService> QueueID: %s, ignoring event with ID: %s, error: %s",
				stQ.cfg.ID, evStatsID, err.Error()))
		}
	}
	return
}

// store stores the necessary data to DB
func (ss *StatService) store() (err error) {
	return
}