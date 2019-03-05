package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/yoktobit/stravagpxexport/strava"
	"github.com/yoktobit/stravagpxexport/util"
)

func main() {

	var dirName string
	flag.StringVar(&dirName, "out", ".", "output dir")
	flag.Parse()

	session := strava.NewSession()

	client := session.Login()

	athlete := new(strava.Athlete)
	util.GetAnswer("athlete", athlete, client)

	activities := make([]strava.Activity, 0)
	util.GetAnswer("athlete/activities", &activities, client)

	for _, activity := range activities {
		activityDetail := new(strava.Activity)
		util.GetAnswer(fmt.Sprintf("activities/%d", activity.ID), activityDetail, client)
		println("Exporting Activity ", activity.ID, ": ", activityDetail.Map.Polyline)

		targetName := fmt.Sprintf("%d.gpx", activity.ID)
		targetName = path.Join(dirName, targetName)
		println(dirName)

		util.ExportPolylineToGpxFile(activityDetail.Map.Polyline, targetName)
	}
}
