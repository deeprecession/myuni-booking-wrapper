package bookings

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"go.uber.org/zap"
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

func GetRooms(log *zap.Logger, myUniCookies []*http.Cookie, start, end time.Time) ([]Room, error) {
	client := getClientWithCookies(myUniCookies)

	roomButtons, err := getRoomButtons(client)
	if err != nil {
		return nil, err
	}

	rooms, err := getAllRoomsBookings(client, roomButtons, start, end)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func getClientWithCookies(myUniCookies []*http.Cookie) *http.Client {
	jar, _ := cookiejar.New(nil)

	myUniversityURL, _ := url.Parse("https://my.university.innopolis.ru")
	jar.SetCookies(myUniversityURL, myUniCookies)

	client := &http.Client{}
	client.Jar = jar

	return client
}

func getAllRoomsBookings(
	client *http.Client,
	roomButtons []RoomButton,
	start, end time.Time,
) ([]Room, error) {
	rooms := make([]Room, 0)

	for _, room := range roomButtons {
		req := newRoomBookingsRequest(room.URL, start, end)

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to make a request to get bookings for room %v : %w",
				room.Name,
				err,
			)
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf(
				"failed to make a request to get bookings for room %v : %w",
				room.Name, ErrMyUniversityApiChanged,
			)
		}
		defer resp.Body.Close()

		bookings, err := parseRoomBookings(resp)
		if err != nil {
			return nil, err
		}

		room := Room{ID: room.ID, Name: room.Name, Bookings: bookings}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

func parseRoomBookings(resp *http.Response) ([]Booking, error) {
	jsonResponse, err := parseBookingsResponse(resp)
	if err != nil {
		return nil, err
	}

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

	return bookings, nil
}

func parseBookingsResponse(resp *http.Response) ([]RootFolderItem, error) {
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Error creating gzip reader: %w", ErrMyUniversityApiChanged)
		}

		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("Error creating gzip reader: %w", err)
	}

	var jsonResponse GetRoomBookingsResponse
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling response body: %w", ErrMyUniversityApiChanged)
	}

	if len(jsonResponse.Body.ResponseMessages.Items) == 0 {
		return nil, fmt.Errorf(
			"ResponseMessages Items have to be not empty: %w",
			ErrMyUniversityApiChanged,
		)
	}

	return jsonResponse.Body.ResponseMessages.Items[0].RootFolder.Items, nil
}
