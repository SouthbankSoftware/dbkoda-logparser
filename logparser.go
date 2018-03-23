/**
 * @Author: Guan Gui <guiguan>
 * @Date:   2018-03-23T12:39:19+11:00
 * @Email:  root@guiguan.net
 * @Last modified by:   guiguan
 * @Last modified time: 2018-03-23T14:36:05+11:00
 */

package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/hpcloud/tail"
)

func main() {
	tailedLogs, err := tail.TailFile(os.Args[1], tail.Config{Follow: true})
	queryR, _ := regexp.Compile("(.+?) D QUERY\\s*\\[[\\w\\d]+\\]\\s*(.+?)\\s*query:\\s*({.*})\\s*sort:\\s*({.*})\\s*projection:\\s*({.*?})\\s*, planSummary:\\s*(.*)\\s*")

	if err != nil {
		fmt.Println(err)
	}

	for line := range tailedLogs.Lines {
		matches := queryR.FindStringSubmatch(line.Text)

		if matches != nil {
			fmt.Printf(`{
  "timestamp": %#v,
  "messages": %#v,
  "query": %#v,
  "sort": %#v,
  "projection": %#v,
  "planSummary": %#v
}
`, matches[1], matches[2], matches[3], matches[4], matches[5], matches[6])
		}
	}
}
