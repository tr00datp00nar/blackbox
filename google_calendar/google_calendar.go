package google_calendar

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/tr00datp00nar/blackbox/utils"
)

const calendarIDEnvKey = "CALENDAR_ID"

// func createNewEvent() {
// 	// Load Environment Variables
// 	utils.LoadEnvVars()
// 	// Read credentials from file
// 	ctx := context.Background()
// 	b, err := os.ReadFile(filepath.Join(utils.UserConfigDir, "credentials.json"))
// 	if err != nil {
// 		log.Fatalf("Unable to read client secret file: %v", err)
// 	}

// 	// Generate new client
// 	// If modifying these scopes, delete your previously saved token.json.
// 	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
// 	if err != nil {
// 		log.Fatalf("Unable to parse client secret file to config: %v", err)
// 	}
// 	client := getClient(config)

// 	// Create new service
// 	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Calendar client: %v", err)
// 	}

// 	// Get all events from calendar
// 	events, err := srv.Events.List(calendarIDEnvKey).Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve calendar events: %v", err)
// 	}

// 	for _, item := range events.Items {
// 		start := item.Start.DateTime
// 		if item.Start.TimeZone != "" {
// 			start = item.Start.TimeZone + ":" + start
// 		}
// 		end := item.End.DateTime
// 		if item.End.TimeZone != "" {
// 			end = item.End.TimeZone + ":" + end
// 		}
// 		fmt.Printf("%s (%s)\n", start, end)
// 	}

// 	// Check if event id matches in notion database
// 	// Get the Notion client
// 	notionClient := update.ConnectNotion()

// 	// Get all pages in the database
// 	dbPages := update.GetScheduledEvents(context.Background(), notionClient)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Notion database pages: %v", err)
// 	}

// 	// Loop through the events and check if the event ID exists in the Notion database
// 	for _, event := range events.Items {
// 		eventID := event.Id
// 		for _, dbPage := range dbPages.Object {
// 			for _, prop := range dbPage {
// 				if prop.ID == "event_id" && prop.Type == "title" {
// 					for _, title := range prop.Title {
// 						if title.Text.Content == eventID {
// 							fmt.Printf("Event ID %s found in Notion database\n", eventID)
// 							break
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func updateEvent()    {}

type DateRange struct {
	Start time.Time
	End   time.Time
}

func GetAvailability() {
	// Load Environment Variables
	UserConfigDir, err := utils.LoadEnvVars()
	if err != nil {
		log.Printf("Unable to load environment variables: %v", err)
	}
	// Read credentials from file
	ctx := context.Background()
	b, err := os.ReadFile(filepath.Join(UserConfigDir, "credentials.json"))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Generate new client
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	// Create new service
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	// Get date range from user
	var start, end time.Time
	for {
		var startStr, endStr string
		fmt.Print("Enter start date (YYYY-MM-DD): ")
		fmt.Scan(&startStr)
		fmt.Print("Enter end date (YYYY-MM-DD): ")
		fmt.Scan(&endStr)

		start, err = time.Parse("2006-01-02", startStr)
		if err != nil {
			log.Println("Invalid start date. Please try again.")
			continue
		}
		end, err = time.Parse("2006-01-02", endStr)
		if err != nil {
			log.Println("Invalid end date. Please try again.")
			continue
		}

		// Check if start is before end
		if start.After(end) {
			log.Println("Start date must be before end date. Please try again.")
			continue
		}

		break
	}

	// Get list of events in date range
	events, err := getEventsInDateRange(srv, start, end)
	if err != nil {
		log.Fatalf("Unable to get events in date range: %v", err)
	}

	// Get list of dates in range with no events
	noEventDates := getDatesWithNoEvents(start, end, events)

	// // Print list of dates with no events
	fileLocation := os.Getenv("RESULTS_LOC")

	if fileLocation == "" {
		fmt.Println("No results location defined in .env")
		os.Exit(1)
	}

	file, err := os.Create(fileLocation)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}

	initDateRange := DateRange{
		Start: start,
		End:   end,
	}
	t := template.Must(template.New("header").Parse(HeaderTmpl))

	err = t.Execute(file, initDateRange)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	defer file.Close()

	t = template.Must(template.New("dates").Parse(BodyTmpl))

	var dateRanges []DateRange
	start = noEventDates[0]
	end = noEventDates[0]

	for i := 1; i < len(noEventDates); i++ {
		current := noEventDates[i]
		previous := noEventDates[i-1]
		diff := current.Sub(previous)

		if diff.Hours()/24 == 1 {
			end = current
		} else {
			dateRanges = append(dateRanges, DateRange{Start: start, End: end})
			start = current
			end = current
		}
	}

	dateRanges = append(dateRanges, DateRange{Start: start, End: end})

	err = t.Execute(file, dateRanges)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Dates written to", fileLocation)
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	configDir, err := utils.GetUserConfig()
	if err != nil {
		log.Fatal(err)
	}
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := filepath.Join(configDir, "token.json")
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// getEventsInDateRange returns a list of events in the given date range
func getEventsInDateRange(srv *calendar.Service, start, end time.Time) ([]*calendar.Event, error) {
	events := []*calendar.Event{}

	// Call the Calendar API
	call := srv.Events.List(os.Getenv(calendarIDEnvKey))
	call = call.TimeMin(start.Format(time.RFC3339))
	call = call.TimeMax(end.Format(time.RFC3339))
	call = call.SingleEvents(true)
	call = call.OrderBy("startTime")

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	events = append(events, response.Items...)

	return events, nil
}

// getDatesWithNoEvents returns a list of dates in the given range with no events
func getDatesWithNoEvents(start, end time.Time, events []*calendar.Event) []time.Time {
	noEventDates := []time.Time{}

	date := start
	for date.Before(end) || date.Equal(end) {
		// Check if there's any event on this date
		eventOnDate := false
		for _, event := range events {
			if event.Start.Date == date.Format("2006-01-02") {
				eventOnDate = true
				break
			}
		}

		if !eventOnDate {
			noEventDates = append(noEventDates, date)
		}

		date = date.AddDate(0, 0, 1)
	}

	return noEventDates
}
