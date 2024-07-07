package bookings

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func newRoomBookingsRequest(url string, start, end time.Time) *http.Request {
	formatedStart := start.UTC().Format("2006-01-02T15:04:05.000")
	formatedEnd := end.UTC().Format("2006-01-02T15:04:05.000")

	jsonData := []byte(fmt.Sprintf(`{
		"__type": "FindItemJsonRequest:#Exchange",
		"Header": {
			"__type": "JsonRequestHeaders:#Exchange",
			"RequestServerVersion": "Exchange2013",
			"TimeZoneContext": {
				"__type": "TimeZoneContext:#Exchange",
				"TimeZoneDefinition": {
					"__type": "TimeZoneDefinitionType:#Exchange",
					"Id": "Russian Standard Time"
				}
			}
		},
		"Body": {
			"__type": "FindItemRequest:#Exchange",
			"ItemShape": {
				"__type": "ItemResponseShape:#Exchange",
				"BaseShape": "IdOnly",
				"AdditionalProperties": [
					{"__type": "PropertyUri:#Exchange", "FieldURI": "ItemParentId"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "Sensitivity"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "AppointmentState"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "IsCancelled"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "HasAttachments"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "LegacyFreeBusyStatus"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "CalendarItemType"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "Start"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "End"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "IsAllDayEvent"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "Organizer"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "Subject"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "IsMeeting"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "UID"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "InstanceKey"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "ItemEffectiveRights"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "JoinOnlineMeetingUrl"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "ConversationId"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "CalendarIsResponseRequested"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "Categories"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "IsRecurring"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "IsOrganizer"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "EnhancedLocation"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "IsSeriesCancelled"},
					{"__type": "PropertyUri:#Exchange", "FieldURI": "Charm"}
				]
			},
			"ParentFolderIds": [
				{
					"__type": "FolderId:#Exchange",
					"Id": "AQMkAGRhNzJmODFkLWEyZDEtNGNkNC1hYQA3My1lOWMxNjNkYjgzOGEALgAAA7IGwULT4DxGjhqS6paOUGMBAHB6D7uh5FtPiONy7WVVDycAAAIBDQAAAA==",
					"ChangeKey": "AgAAAA=="
				}
			],
			"Traversal": "Shallow",
			"Paging": {
				"__type": "CalendarPageView:#Exchange",
				"StartDate": "%s",
				"EndDate": "%s"
			}
		}
	}`, formatedStart, formatedEnd))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Action", "FindItem")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return req
}
