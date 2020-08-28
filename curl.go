package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

// func maketrans() {
// 	dialer := &net.Dialer{
// 		Timeout:   30 * time.Second,
// 		KeepAlive: 30 * time.Second,
// 		DualStack: true,
// 	}
// 	// or create your own transport, there's an example on godoc.
// 	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
// 		if addr == tra {
// 			addr = selectips["f5"]
// 		}
// 		return dialer.DialContext(ctx, network, addr)
// 	}
// }

func main() {
	// logs.Debug()
	tra := "transact.williamhill-pp1.com:443"
	selectips := make(map[string]string)
	selectips["f5"] = "54.77.166.81:443"
	selectips["nlb"] = "100.73.177.242:443"
	selectips["cf"] = "13.32.169.89:443"

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	// or create your own transport, there's an example on godoc.
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		if addr == tra {
			addr = selectips["cf"]
		}
		return dialer.DialContext(ctx, network, addr)
	}
	// maketrans(tra)

	req, err := http.NewRequest("GET", "https://transact.williamhill-pp1.com/betslip/actuator", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Add Request Header
	req.Host = "transact.williamhill-pp1.com"

	// Create and Add cookie to request
	cookie := http.Cookie{
		Name:  "awsbts",
		Value: "yes",
	}
	req.AddCookie(&cookie)

	client := &http.Client{Timeout: time.Second * 10}

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	fmt.Printf("--resolve: %s:%s \n\n", tra, selectips["cf"])

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)

	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println("\n", string(body))
	fmt.Println("Response Headers")
	fmt.Println("Status:", string(resp.Status))
	for key, value := range resp.Header {
		fmt.Printf("%s: %s\n", key, value)
	}
	// fmt.Println(resp.TLS)
}
