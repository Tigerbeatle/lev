// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package main provides sample web application for beego and mgo.
package main

import (
	"github.com/astaxie/beego"
	"github.com/tigerbeatle/le/localize"
	_ "github.com/tigerbeatle/le/routes"
	"github.com/tigerbeatle/le/utilities/helper"
	"github.com/tigerbeatle/le/utilities/mongo"
	"github.com/tigerbeatle/tracelog"
	"os"
)

func main() {
	tracelog.Start(tracelog.LevelTrace)

	// Init mongo
	tracelog.Started("main", "Initializing Mongo")
	err := mongo.Startup(helper.MainGoRoutine)
	if err != nil {
		tracelog.CompletedError(err, helper.MainGoRoutine, "initApp")
		os.Exit(1)
	}

	// Load message strings
	localize.Init("en-US")

	beego.Run()

	tracelog.Completed(helper.MainGoRoutine, "Website Shutdown")
	tracelog.Stop()
}
