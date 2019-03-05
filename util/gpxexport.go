package util

import (
	"bufio"
	"os"

	gpx "github.com/twpayne/go-gpx"
	polyline "github.com/twpayne/go-polyline"
)

func ExportPolylineToGpxFile(polyLine string, fileName string) {
	coords, _, _ := polyline.DecodeCoords([]byte(polyLine))

	var g gpx.GPX
	trk := new(gpx.TrkType)
	trkSeg := new(gpx.TrkSegType)
	for _, coord := range coords {
		wpt := new(gpx.WptType)
		wpt.Lat = coord[0]
		wpt.Lon = coord[1]
		trkSeg.TrkPt = append(trkSeg.TrkPt, wpt)
	}
	trk.TrkSeg = append(trk.TrkSeg, trkSeg)
	g.Trk = append(g.Trk, trk)

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	w := bufio.NewWriter(f)
	if err := g.WriteIndent(w, "", "  "); err != nil {
		panic(err)
	}
}
