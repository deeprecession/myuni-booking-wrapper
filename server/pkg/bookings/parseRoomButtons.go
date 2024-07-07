package bookings

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type RoomButton struct {
	ID   int
	Name string
	URL  string
}

func getRoomButtons(client *http.Client) ([]RoomButton, error) {
	response, err := client.Get("https://my.university.innopolis.ru/profile/room-booking")
	if err != nil {
		log.Fatalf("failed to get bookings links: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("staus code have to be StatusOK: %w", ErrMyUniversityApiChanged)
	}

	responseHTML, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse html: %w", ErrMyUniversityApiChanged)
	}

	roomButtons := findRoomButtons(responseHTML)

	if len(roomButtons) == 0 {
		return nil, fmt.Errorf("%w", ErrBadCookies)
	}

	roomButtons, err = transformButtonURLs(roomButtons)
	if err != nil {
		return nil, fmt.Errorf("failed to transform butotn URLs: %w", ErrMyUniversityApiChanged)
	}

	return roomButtons, nil
}

func findRoomButtons(n *html.Node) []RoomButton {
	var buttons []RoomButton

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "button" {
			var roomButton RoomButton

			for _, attr := range n.Attr {
				if attr.Key == "data-id" {
					_, _ = fmt.Sscanf(attr.Val, "%d", &roomButton.ID)
					break
				}
			}

			var extractName func(*html.Node)
			extractName = func(n *html.Node) {
				if n.Type == html.TextNode {
					roomButton.Name += strings.TrimSpace(n.Data)
				}
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					extractName(c)
				}
			}
			extractName(n)

			for _, attr := range n.Attr {
				if attr.Key == "href" {
					roomButton.URL = attr.Val
					break
				}
			}

			if roomButton.ID != 0 {
				buttons = append(buttons, roomButton)
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)

	return buttons
}

func transformButtonURLs(roomButtons []RoomButton) ([]RoomButton, error) {
	newRoomButtons := make([]RoomButton, 0, len(roomButtons))

	for _, room := range roomButtons {

		newURL, err := transformURL(room.URL)
		if err != nil {
		}

		newRoom := RoomButton{
			ID:   room.ID,
			Name: room.Name,
			URL:  newURL,
		}

		newRoomButtons = append(newRoomButtons, newRoom)
	}

	return newRoomButtons, nil
}

func transformURL(oldURL string) (string, error) {
	u, err := url.Parse(oldURL)
	if err != nil {
		return "", fmt.Errorf("bad room button URL: %w", ErrMyUniversityApiChanged)
	}

	parts := strings.Split(u.Path, "/")
	if len(parts) < 4 {
		return "", fmt.Errorf("bad room button URL: %w", ErrMyUniversityApiChanged)
	}

	newURL := fmt.Sprintf(
		"https://mail.innopolis.ru/owa/calendar/%s/%s/service.svc?action=FindItem&ID=-1&AC=1",
		parts[3],
		parts[4],
	)
	return newURL, nil
}
