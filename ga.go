package main

import (
	"log"
	"strconv"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/analytics/v3"
)

var jwtc = jwt.Config{
	Email:      gaServiceAcctEmail,
	PrivateKey: []byte(goPrivateKey),
	Scopes:     []string{analytics.AnalyticsReadonlyScope},
	TokenURL:   gaTokenurl,
}
var clt = jwtc.Client(oauth2.NoContext)

var maxResults = 21

type query struct {
	metric     string
	dimensions string
	maxResults int64
	sort       string
	returnRows bool
}

func (p *query) getData() *analytics.RealtimeData {
	as, err := analytics.New(clt)
	if err != nil {
		log.Fatal("Error creating Analytics Service at analytics.New() -", err)
	}

	rt := analytics.NewDataRealtimeService(as)
	rtSetup := rt.Get(gaTableID, p.metric)
	rtSetup.Dimensions(p.dimensions)

	if p.sort != "" {
		rtSetup.Sort(p.sort)
	}

	if p.maxResults != 0 {
		rtSetup.MaxResults(p.maxResults)
	}

	gtadata, err := rtSetup.Do()

	if err != nil {
		log.Fatal("Could not load data from GA service:", err)

	}

	return gtadata

}

func getRtMostViewedPages() []string {
	q := query{
		metric:     "rt:activeUsers",
		dimensions: "rt:pageTitle,rt:pagePath",
		maxResults: int64(maxResults),
		sort:       "-rt:activeUsers",
		returnRows: true,
	}

	result := q.getData()

	rtMostViewed := make([]string, maxResults)

	for i, data := range result.Rows {
		if data[1] == "/" {
			data[0] = "Home"
		}

		rtMostViewed[i] = "[" + data[2] + "](fg-cyan) " + data[0]

	}
	return rtMostViewed

}

func updateRtMostViewedByTrafficType(n *int, t *map[string]int, bcD *[]int, bcL *[]string) {

	trafficType := map[string]int{
		"CUSTOM":   0,
		"DIRECT":   0,
		"ORGANIC":  0,
		"REFERRAL": 0,
		"SOCIAL":   0,
	}

	q := query{
		metric:     "rt:activeUsers",
		dimensions: "rt:trafficType",
	}

	result := q.getData()

	nusers, _ := strconv.Atoi(result.TotalsForAllResults["rt:activeUsers"])

	for _, result := range result.Rows {
		percentage, _ := strconv.Atoi(result[1])
		trafficType[result[0]] = percentage * 100 / nusers
	}

	*bcL = addBcLabel(*bcL)
	*bcD = addBcData(*bcD, nusers)

	*t = trafficType
	*n = nusers

}

func addBcData(a []int, value int) []int {

	a = append([]int{value}, a...)
	a = a[:len(a)-1]

	return a

}

func addBcLabel(d []string) []string {
	t := time.Now().Format("03:04")
	if t != d[0] {
		d = append([]string{t}, d...)
		d = d[:len(d)-1]
	}

	return d
}
