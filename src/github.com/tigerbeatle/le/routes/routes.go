// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package routes initializes the routes for the web service.
package routes

import (
	"github.com/astaxie/beego"
	"github.com/tigerbeatle/le/controllers"
)

func init() {
	beego.Router("/", new(controllers.EggController), "get:Index")
	beego.Router("/egg/retrievestation", new(controllers.EggController), "post:RetrieveEggStats")
	beego.Router("/egg/stats/:stationId", new(controllers.EggController), "get,post:RetrieveEggStatsJSON")
}
