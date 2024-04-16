package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	TLD, IDN_TLD := GetTLD(website.Domain)

	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
		return
	}

	if !website.Valid {
		chErr <- errors.New("domain not valid")
		return
	}

	if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
		return
	}

	ch <- RowData{
		RankWebsite: website.RankWebsite,
		Domain:      website.Domain,
		TLD:         TLD,
		IDN_TLD:     IDN_TLD,
		Valid:       website.Valid,
		RefIPs:      website.RefIPs,
	}

}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)
	var filteredData []RowData

	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
		if website.Domain == "" {
			return nil, errors.New("domain name is empty")
		} else if !website.Valid {
			return nil, errors.New("domain not valid")
		} else if website.RefIPs == -1 {
			return nil, errors.New("domain RefIPs not valid")
		} else if strings.HasSuffix(website.Domain, TLD) {
			filteredData = append(filteredData, website)
		}
	}

	time.Sleep(100 * time.Millisecond)

	return filteredData, nil // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
