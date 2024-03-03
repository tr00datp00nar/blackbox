Search a Google Calendar within a date range for dates with no events.

This go program will prompt the user for a start date and end date and uses Google OAuth2.0 to authenticate with google calendar API. It then searches through a calendar (proovided by calendar id) for all dates within the given range with no events and returns them as a list.

## Prerequisites

1. Create a [Google Cloud](https://console.cloud.google) account and project.
2. Generate OAuth2.0 credentials using the google cloud console:
   - Credentials should be of type 'desktop app'
   - Download the credentials into the same directory as this program. Save the file as `credentials.json`
3. Create a `.env` file in the same directory as this program.
   `CALENDAR_ID="<calendar-id>"`
