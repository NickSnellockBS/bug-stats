package retrievebugtickets

import (
	decodeticket "debug-stats/decode-ticket"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type TicketDetail struct {
	ID int
	CreatedAt time.Time
	CompletedAt interface{}
}

func RetrieveTickets(url string, year int, firstHalf bool) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	client := &http.Client{}

	queryString := fmt.Sprintf(`type:bug created:%d-07-01..%d-12-31`, year, year)

	if firstHalf {
		queryString = fmt.Sprintf(`type:bug created:%d-01-01..%d-06-30`, year, year)
	}

	requestBody := strings.NewReader(fmt.Sprintf(`{"detail": "slim", "page_size":25, "query": "%s"}`, queryString))
	fullUrl := fmt.Sprintf("https://api.app.shortcut.com%s", url)
	req, err := http.NewRequest("GET", fullUrl, requestBody)

	if (err != nil) {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Shortcut-Token", os.Getenv("SHORTCUT_TOIKEN"))

	resp, err := client.Do(req)

	if (err != nil) {
		fmt.Println(err)
		return ""
	}

	body, err := io.ReadAll(resp.Body)

	if (err != nil) {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

func GetTickets() (int, []TicketDetail) {
	shortcutTickets := decodeticket.Ticket{}
	tickets := []TicketDetail{}

	totalTickets := 0

	for year := 2015; year <= time.Now().Year(); year++ {
		for half := 0; half <= 1; half++ {
			url := "/api/v3/search/stories"
			shortcutTickets = GetTicket(url, year, half == 0)

			totalTickets += shortcutTickets.Total
			totalTicketsThisYear := shortcutTickets.Total

			if totalTicketsThisYear > 1000 {
				halfString := "first half of"
				if half == 1 {
					halfString = "second half of"
				}
				fmt.Printf("More than 1000 tickets (%d) for %s %d\n", totalTicketsThisYear, halfString, year)
			}

			for i := 0; i < len(shortcutTickets.Data); i++ {
				ticketDetail := TicketDetail{
					ID: shortcutTickets.Data[i].ID,
					CreatedAt: shortcutTickets.Data[i].CreatedAt,
					CompletedAt: shortcutTickets.Data[i].CompletedAt}

				tickets = append(tickets, ticketDetail)
			}

			ticketsRemaining := totalTicketsThisYear - len(shortcutTickets.Data)

			url = shortcutTickets.Next

			for i := ticketsRemaining; i > 0; i++ {
				shortcutTickets = GetTicket(url, year, half == 0)
				for j := 0; j < len(shortcutTickets.Data); j++ {
					ticketDetail := TicketDetail{
						ID: shortcutTickets.Data[j].ID,
						CreatedAt: shortcutTickets.Data[j].CreatedAt,
						CompletedAt: shortcutTickets.Data[j].CompletedAt}
		
					tickets = append(tickets, ticketDetail)
				}
				i = len(shortcutTickets.Data) - 1
				url = shortcutTickets.Next
			}
		}
	}

	return totalTickets, tickets
}

func GetTicket(url string, year int, firstHalf bool) decodeticket.Ticket {
	ticketJson := RetrieveTickets(url, year, firstHalf)
	ticket := decodeticket.DecodeTicket(ticketJson)

	return ticket
}