package fakes

import (
	"github.com/cloudfoundry-incubator/notifications/db"
	"github.com/cloudfoundry-incubator/notifications/gobble"
)

type StrategyDeterminer struct {
	DetermineCall struct {
		Receives struct {
			Connection db.ConnectionInterface
			UAAHost    string
			Job        gobble.Job
		}
		WasCalled bool
	}
}

func NewStrategyDeterminer() *StrategyDeterminer {
	return &StrategyDeterminer{}
}

func (s *StrategyDeterminer) Determine(conn db.ConnectionInterface, uaaHost string, job gobble.Job) {
	s.DetermineCall.Receives.Connection = conn
	s.DetermineCall.Receives.UAAHost = uaaHost
	s.DetermineCall.Receives.Job = job
	s.DetermineCall.WasCalled = true
}
