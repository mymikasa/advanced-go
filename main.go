//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"strings"
//	"sync"
//)
//
//func Req(wg *sync.WaitGroup) {
//
//	url := "https://cn.airbusan.com/web/bookingApi/flightsAvail"
//	method := "POST"
//
//	payload := strings.NewReader(`bookingCategory=Individual&focYn=N&tripType=RT&depCity1=TAO&depCity2=PUS&depCity3=&depCity4=&arrCity1=PUS&arrCity2=TAO&arrCity3=&arrCity4=&depDate1=2022-12-20&depDate2=2022-12-23&depDate3=&depDate4=&itinNo=&openReturnYN=&paxCountCorp=0&paxCountAd=1&paxCountCh=0&paxCountIn=0`)
//
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, payload)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0")
//	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
//	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
//	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
//	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
//	req.Header.Add("X-Requested-With", "XMLHttpRequest")
//	req.Header.Add("Origin", "https://cn.airbusan.com")
//	req.Header.Add("Connection", "keep-alive")
//	req.Header.Add("Referer", "https://cn.airbusan.com/web/individual/booking/flightsAvail")
//	req.Header.Add("Cookie", "cf_chl_2=e732132f89468df; cf_clearance=MEuP7SyAxXevo8ppyWGN9zIiYQUyZlP7ulx0EMbu6IY-1671502578-0-160; PLAY_LANG=zh-CN; _ga=GA1.2.1101151323.1671501709; _gid=GA1.2.1159645409.1671501709; _gat=1; _gat_gtag_UA_61679819_1=1; NetFunnel_ID=5002%3A200%3Akey%3D790D00AC447B9884AC648F554BE17FE72C82362C22AB7C9855B3EEE49568661242D27492A2CB8A07AA050F21B454FF16CBF30FDECDD4238808FF23B15161478A7C7EC2AE827D6D2C4CC6E76ED33779B1CFEA7F080CAACB616E70137C5E228AF6FEDEA30D151FC8445ACF2BAEA89FD2677661696C2C302C322C312C30%26nwait%3D0%26nnext%3D0%26tps%3D0%26ttl%3D0%26ip%3Dpromo.airbusan.com%26port%3D443")
//	req.Header.Add("Sec-Fetch-Dest", "empty")
//	req.Header.Add("Sec-Fetch-Mode", "cors")
//	req.Header.Add("Sec-Fetch-Site", "same-origin")
//
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//
//	wg.Done()
//}
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(10)
//	for i := 0; i < 10; i++ {
//		go Req(&wg)
//	}
//	wg.Wait()
//}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://us.trip.com/flights/graphql/intlFlightListSearchAll"
	method := "POST"

	payload := strings.NewReader(`{
    "operationName": "flightListSearch",
    "variables": {
        "request": {
            "Head": {
                "Currency": "USD",
                "ExtendFields": {
                    "SpecialSupply": "false",
                    "TraceId": ""
                }
            },
            "fullData": false,
            "criteriaToken": "",
            "shoppingId": null,
            "segmentNo": 1,
            "searchCriteria": {
                "cabinClass": "YSGroup",
                "searchSegmentList": [
                    {
                        "departDate": "2022-12-23",
                        "departCity": "SZX",
                        "arriveCity": "TYO"
                    },
                    {
                        "departDate": "2022-12-26",
                        "departCity": "TYO",
                        "arriveCity": "SZX"
                    }
                ],
                "passengerCount": {
                    "adult": 1,
                    "child": 0,
                    "infant": 0
                }
            },
            "mode": 1,
            "sort": {
                "topAirline": true,
                "asc": true,
                "topList": [],
                "type": "Score"
            }
        }
    },
    "extensions": {
        "persistedQuery": {
            "version": 1,
            "sha256Hash": "ee76b682ced93645d97a486e4a610c1abd2974dac51aae50e23b5f28fdba9564"
        }
    }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "us.trip.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "_bfa=1.1671601525169.1luhnv.1.1671601525169.1671601525169.1.3.1; _bfs=1.6; ")
	req.Header.Add("origin", "https://us.trip.com")
	req.Header.Add("referer", "https://us.trip.com/flights/shenzhen-to-tokyo/tickets-szx-tyo?dcity=szx&acity=tyo&ddate=2022-12-23&rdate=2022-12-26&flighttype=rt&class=y&quantity=1&searchboxarg=t")
	req.Header.Add("sec-ch-ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/532.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := string(body)
	fmt.Println(strings.Contains(data, "AntiCrawler"))
	//fmt.Println()
}
