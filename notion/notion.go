package notion

import (
	"context"
	"log"
	"os"

	"github.com/jomei/notionapi"

	"github.com/tr00datp00nar/blackbox/utils"
)

func getToken() notionapi.Token {
	utils.LoadEnvVars()
	userToken := notionapi.Token(os.Getenv("NOTION_API_TOKEN"))
	return userToken
}

func ConnectNotion() *notionapi.Client {
	client := notionapi.NewClient(getToken())
	return client
}

type ScheduledEventNotion struct {
	Title         *notionapi.TitleProperty
	ScheduleID    *notionapi.UniqueID
	SpaceRented   *notionapi.MultiSelectProperty
	StartEndDates *notionapi.DateObject
}

type GoogleCalendarEvent struct {
}

// Update Google Calendar Events

// Get a list of Notion Events
func GetScheduledEvents(ctx context.Context, client *notionapi.Client) *notionapi.Database {
	db, err := client.Database.Get(ctx, notionapi.DatabaseID(os.Getenv("NOTION_SCHEDULE_DB_ID")))
	if err != nil {
		log.Println("Error getting database:", err)
		return nil
	}
	// TODO: Unmarshal the response into a ScheduledEvent object and return it
	return db
}

// Set Google Calendar ID based on ScheduledEvent.Location
// If Location == Theater >> Compare against Theater Google Calendar
// If Location  == Rehearsal Room >> Compare Against Rehearsal Room Google Calendar
// If ScheduledEvent.GCalID already exists in Google Calendar, update the calendar event
// If ScheduledEvent.GCalID is not in Google Calendar, add it
// For ScheduledEvent.GCalID, if ScheduledEvent.Cancelled is true, remove from Google Calendar
