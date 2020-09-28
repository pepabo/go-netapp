package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/pepabo/go-netapp/netapp"
)

func main() {
	c, err := netapp.NewClient(
		"https://localhost:1443",
		"1.20",
		&netapp.ClientOptions{
			BasicAuthUser:     "oohori",
			BasicAuthPassword: "F:T3pxJSEbuMLuWh",
			SSLVerify:         false,
			Debug:             true,
		},
	)
	if err != nil {
		panic(err)
	}

	q := "f8e9666a-364a-4eb0-98b9-1dc78d5ba641"
	ql, _, _ := c.Quota.List(
		&netapp.QuotaOptions{
			Query: &netapp.QuotaQuery{
				QuotaEntry: &netapp.QuotaEntry{
					Volume:      "volume6",
					QuotaTarget: fmt.Sprintf("/vol/%s/%s", "volume6", q),
				}},
		})

	pp.Println(ql)
	/*
		qll := ql.Results.AttributesList.QuotaEntry[0]
		qll.FileLimit = strconv.Itoa(160000)
		qll.Vserver = ""
			qlu, _, err := c.Quota.Update(
				"cl1_nfs01",
				&qll,
			)
		qqq, _, err := c.Qtree.List(&netapp.QtreeOptions{
			MaxRecords: 1,
			Query: &netapp.QtreeQuery{
				QtreeInfo: &netapp.QtreeInfo{
					Vserver: "cl1_nfs01",
					Volume:  "volume6",
					Qtree:   q,
				},
			},
		})

		pp.Println(qqq)
		pp.Println(err)
	*/

}
