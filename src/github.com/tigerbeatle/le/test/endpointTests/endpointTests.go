// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package endpointTests implements boilerplate code for all testing.
package endpointTests

import (
	"github.com/tigerbeatle/le/localize"
	_ "github.com/tigerbeatle/le/routes" // Initalize routes
	"github.com/tigerbeatle/le/utilities/helper"
	"github.com/tigerbeatle/le/utilities/mongo"
	log "github.com/tigerbeatle/tracelog"
)

//** CONSTANTS

const (
	// SessionID is just mocking the id for testing.
	SessionID = "testing"
)

//** INIT

// init initializes all required packages and systems
func init() {
	log.Start(log.LevelTrace)

	// Init mongo
	log.Started("main", "Initializing Mongo")
	err := mongo.Startup(helper.MainGoRoutine)
	if err != nil {
		log.CompletedError(err, helper.MainGoRoutine, "initTesting")
		return
	}

	// Load message strings
	localize.Init("en-US")
}
