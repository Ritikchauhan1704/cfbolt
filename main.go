package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	targetURL := "https://codeforces.com/enter"

	fmt.Printf("📡 Making GET request to %s...\n", targetURL)

	resp, err := http.Get(targetURL)
	if err != nil {
		log.Fatalf("❌ Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	cookies := resp.Cookies()
	var cookieStrings []string
	fmt.Println("\n🍪 Cookies received:")
	for _, cookie := range cookies {
		fmt.Printf("  - %s = %s\n", cookie.Name, cookie.Value)
		cookieStrings = append(cookieStrings, cookie.String())
	}
	
	fullCookieHeader := strings.Join(cookieStrings, "; ")
	fmt.Printf("\nFull Cookie Header for next request:\n%s\n", fullCookieHeader)


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("❌ Failed to read response body: %v", err)
	}

	re := regexp.MustCompile(`name='csrf_token' value='([a-f0-9]{32})'`)
	matches := re.FindStringSubmatch(string(body))

	if len(matches) < 2 {
		log.Fatal("❌ CSRF token not found. The page structure might have changed.")
	}

	csrfToken := matches[1]
	fmt.Println("\n🔑 CSRF Token found:")
	fmt.Println(csrfToken)
}