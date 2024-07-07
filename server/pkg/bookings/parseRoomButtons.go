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

	responseHTML, err := html.Parse(response.Body)
	if err != nil {
		log.Fatalf("failed to parse html: %v", err)
	}

	return findRoomButtons(responseHTML), nil
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
					roomButton.URL = transformURL(attr.Val)
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

func transformURL(oldURL string) string {
	u, err := url.Parse(oldURL)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}
	parts := strings.Split(u.Path, "/")
	if len(parts) < 4 {
		log.Fatalf("Unexpected URL format: %s", oldURL)
	}
	// New URL format
	newURL := fmt.Sprintf(
		"https://mail.innopolis.ru/owa/calendar/%s/%s/service.svc?action=FindItem&ID=-1&AC=1",
		parts[3],
		parts[4],
	)
	return newURL
}
