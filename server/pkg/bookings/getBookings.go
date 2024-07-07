package bookings

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type Booking struct {
	Name  string
	Start time.Time
	End   time.Time
}

type Room struct {
	Name     string
	Bookings []Booking
	ID       int
}

func GetRooms(myUniCookies []*http.Cookie, start, end time.Time) ([]Room, error) {
	client := getClientWithCookies(myUniCookies)

	roomButtons, err := getRoomButtons(client)
	if err != nil {
		return nil, err
	}

	rooms := getAllRoomsBookings(client, roomButtons, start, end)

	return rooms, nil
}

func getClientWithCookies(myUniCookies []*http.Cookie) *http.Client {
	jar, _ := cookiejar.New(nil)

	myUniversityURL, _ := url.Parse("https://my.university.innopolis.ru")
	jar.SetCookies(myUniversityURL, myUniCookies)

	client := retryablehttp.NewClient()
	client.HTTPClient.Jar = jar

	return client.StandardClient()
}

func getAllRoomsBookings(
	client *http.Client,
	urls []RoomButton,
	start, end time.Time,
) []Room {
	rooms := make([]Room, 0)

	for _, roomButtons := range urls {
		req := newRoomBookingsRequest(roomButtons.URL, start, end)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		bookings := parseRoomBookings(resp)

		room := Room{ID: roomButtons.ID, Name: roomButtons.Name, Bookings: bookings}

		rooms = append(rooms, room)
	}

	return rooms
}

func parseRoomBookings(resp *http.Response) []Booking {
	jsonResponse := parseBookingsResponse(resp)

	bookings := make([]Booking, 0)
	for _, item := range jsonResponse {

		start, _ := time.Parse(time.RFC3339, item.Start)
		end, _ := time.Parse(time.RFC3339, item.End)

		booking := Booking{
			Name:  item.Subject,
			Start: start,
			End:   end,
		}

		bookings = append(bookings, booking)
	}

	return bookings
}

func parseBookingsResponse(resp *http.Response) []RootFolderItem {
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			log.Fatalf("Error creating gzip reader: %v", err)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var jsonResponse GetRoomBookingsResponse
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}

	return jsonResponse.Body.ResponseMessages.Items[0].RootFolder.Items
}
