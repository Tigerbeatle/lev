// Copyright 2013 Ardan Studios. All rights reserved.
// Use of controller source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package controllers implements the controller layer for the egg API.
package controllers

import (
    "github.com/tigerbeatle/le/models/eggModels"
    bc "github.com/tigerbeatle/le/controllers/baseController"
    "github.com/tigerbeatle/le/services/eggService"
    log "github.com/tigerbeatle/tracelog"
    "encoding/json"
//"os/user"
// "debug/gosym"
    "strconv"
)

//** TYPES

// EggController manages the API for egg related functionality.
type EggController struct {
    bc.BaseController
}

//** WEB FUNCTIONS

// Index is the initial view for the egg system.
func (controller *EggController) Index() {
    region := "Gulf Of Mexico"
    log.Startedf(controller.UserID, "EggController.Index", "Region[%s]", region)

    eggStations, err := eggService.FindRegion(&controller.Service, region)
    if err != nil {
        log.CompletedErrorf(err, controller.UserID, "EggController.Index", "Region[%s]", region)
        controller.ServeError(err)
        return
    }

    controller.Data["Stations"] = eggStations
    controller.Layout = "shared/basic-layout.html"
    controller.TplNames = "egg/content.html"
    controller.LayoutSections = map[string]string{}
    controller.LayoutSections["PageHead"] = "egg/page-head.html"
    controller.LayoutSections["Header"] = "shared/header.html"
    controller.LayoutSections["Modal"] = "shared/modal.html"
}



// Purchase records a sale though the inworld system
func (controller *EggController) Purchase() {

    Timestamp, _        := strconv.Atoi(controller.Ctx.Input.Query("timestamp"))
    UnitsSold, _        := strconv.Atoi(controller.Ctx.Input.Query("unitsSold"))
    BonusPayment, _     := strconv.Atoi(controller.Ctx.Input.Query("bonusPayment"))
    Payment, _          := strconv.Atoi(controller.Ctx.Input.Query("payment"))
    VendorID, _         := strconv.Atoi(controller.Ctx.Input.Query("vendorID"))
    TransactionID, _    := strconv.Atoi(controller.Ctx.Input.Query("transactionID"))
    ContestID, _        := strconv.Atoi(controller.Ctx.Input.Query("contestID"))

    Owner := &eggModels.Person{
        UUID:   controller.Ctx.Input.Query("ownerUUID"),
        Name:   controller.Ctx.Input.Query("ownerName"),
    }
    Buyer := &eggModels.Person{
        UUID:   controller.Ctx.Input.Query("buyerUUID"),
        Name:   controller.Ctx.Input.Query("buyerName"),
    }

    t := &eggModels.Transaction{
        Timestamp:      Timestamp,
        Date:           controller.Ctx.Input.Query("date"),
        UnitsSold:      UnitsSold,
        BonusPayment:   BonusPayment,
        Payment:        Payment,
        Location:       controller.Ctx.Input.Query("location"),
        TransactionID:  TransactionID,
        VendorID:       VendorID,
        ContestID:      ContestID,
        Owner:          *Owner,
        Buyer:          *Buyer,
    }
    Transaction, _ := json.Marshal(t)

    log.Startedf(controller.UserID, "EggController.Purchase", "t=[%s]", Transaction)

    _, err := eggService.RecordPurchase(&controller.Service, Transaction)
    if err != nil {
        log.CompletedErrorf(err, controller.UserID, "EggController.Purchase", "t=[%s]", Transaction)
        controller.ServeError(err)
        return
    }

    controller.Layout = "shared/basic-layout.html"
    controller.TplNames = "egg/content.html"
    controller.LayoutSections = map[string]string{}
    controller.LayoutSections["PageHead"] = "egg/page-head.html"
    controller.LayoutSections["Header"] = "shared/header.html"
    controller.LayoutSections["Modal"] = "shared/modal.html"

}



// Purchase records a sale though the inworld system
func (controller *EggController) DNS() {

    SerialNumber, _     := strconv.Atoi(controller.Ctx.Input.Query("serialNumber"))
    Version, _          := strconv.Atoi(controller.Ctx.Input.Query("version"))
    AliveTestCount, _   := strconv.Atoi(controller.Ctx.Input.Query("aliveTestCount"))

    Owner := &eggModels.Person{
        UUID:   controller.Ctx.Input.Query("ownerUUID"),
        Name:   controller.Ctx.Input.Query("ownerName"),
    }

    DnsParcel := &eggModels.Parcel{
        Surl:   controller.Ctx.Input.Query("parcelSurl"),
        Url:    controller.Ctx.Input.Query("parcelUrl"),
        Name:   controller.Ctx.Input.Query("parcelName"),
    }

    t := &eggModels.EggDNS{
        SerialNumber:       SerialNumber,
        Language:           controller.Ctx.Input.Query("language"),
        Version:            Version,
        AliveTestCount:     AliveTestCount,
        RemoveTarget:       controller.Ctx.Input.Query("removeTarget"),
        AliveTestStatus:    controller.Ctx.Input.Query("aliveTestStatus"),
        Owner:              *Owner,
        Parcel:             *DnsParcel,
    }
    EggDNS, _ := json.Marshal(t)

    log.Startedf(controller.UserID, "EggController.DNS", "trans2[%s]", EggDNS)

    _, err := eggService.RecordDNS(&controller.Service, EggDNS)
    if err != nil {
        log.CompletedErrorf(err, controller.UserID, "EggController.DNS", "trans2[%s]", EggDNS)
        controller.ServeError(err)
        return
    }

    controller.Layout = "shared/basic-layout.html"
    controller.TplNames = "egg/content.html"
    controller.LayoutSections = map[string]string{}
    controller.LayoutSections["PageHead"] = "egg/page-head.html"
    controller.LayoutSections["Header"] = "shared/header.html"
    controller.LayoutSections["Modal"] = "shared/modal.html"

}









//** AJAX FUNCTIONS

// RetrieveStation handles the example 2 tab.
func (controller *EggController) RetrieveStation() {
    var params struct {
        StationID string `form:"stationID" valid:"Required; MinSize(4)" error:"invalid_station_id"`
    }

    if controller.ParseAndValidate(&params) == false {
        return
    }

    eggStation, err := eggService.FindStation(&controller.Service, params.StationID)
    if err != nil {
        log.CompletedErrorf(err, controller.UserID, "EggController.RetrieveStation", "StationID[%s]", params.StationID)
        controller.ServeError(err)
        return
    }

    controller.Data["Station"] = eggStation
    controller.Layout = ""
    controller.TplNames = "egg/modal/pv_station-detail.html"
    view, _ := controller.RenderString()

    controller.AjaxResponse(0, "SUCCESS", view)
}



// RecordTransactionJSON handles the example 2 tab.
func (controller *EggController) RecordTransactionJSON() {
    var params struct {
        PurchaseJson string `form:"json" valid:"Required; MinSize(4)" error:"invalid_json"`
    }

    if controller.ParseAndValidate(&params) == false {
        return
    }

    eggStation, err := eggService.FindStation(&controller.Service, params.PurchaseJson)
    if err != nil {
        log.CompletedErrorf(err, controller.UserID, "EggController.RecordTransactionJSON", "JSON[%s]", params.PurchaseJson)
        controller.ServeError(err)
        return
    }

    controller.Data["Json"] = eggStation
    controller.Layout = ""
    controller.TplNames = "egg/modal/pv_station-detail.html"
    view, _ := controller.RenderString()

    controller.AjaxResponse(0, "SUCCESS", view)
}










// RetrieveStationJSON handles the example 3 tab.
// http://localhost:9003/egg/station/42002
func (controller *EggController) RetrieveStationJSON() {
    // The call to ParseForm inside of ParseAndValidate is failing. This is a BAD FIX
    params := struct {
        StationID string `form:":stationId" valid:"Required; MinSize(4)" error:"invalid_station_id"`
    }{controller.GetString(":stationId")}

    if controller.ParseAndValidate(&params) == false {
        return
    }

    eggStation, err := eggService.FindStation(&controller.Service, params.StationID)
    if err != nil {
        log.CompletedErrorf(err, controller.UserID, "Station", "StationID[%s]", params.StationID)
        controller.ServeError(err)
        return
    }

    controller.Data["json"] = eggStation
    controller.ServeJson()
}


