package bookings

type GetRoomBookingsResponse struct {
	Header Header `json:"Header"`
	Body   Body   `json:"Body"`
}

type Body struct {
	ResponseMessages ResponseMessages `json:"ResponseMessages"`
}

type ResponseMessages struct {
	Items []ResponseMessagesItem `json:"Items"`
}

type ResponseMessagesItem struct {
	Type           string      `json:"__type"`
	ResponseCode   string      `json:"ResponseCode"`
	ResponseClass  string      `json:"ResponseClass"`
	HighlightTerms interface{} `json:"HighlightTerms"`
	RootFolder     RootFolder  `json:"RootFolder"`
}

type RootFolder struct {
	IncludesLastItemInRange bool             `json:"IncludesLastItemInRange"`
	TotalItemsInView        int64            `json:"TotalItemsInView"`
	Groups                  interface{}      `json:"Groups"`
	Items                   []RootFolderItem `json:"Items"`
}

type RootFolderItem struct {
	Type             ItemType         `json:"__type"`
	ItemID           ID               `json:"ItemId"`
	ParentFolderID   ID               `json:"ParentFolderId"`
	Subject          string           `json:"Subject"`
	Sensitivity      Sensitivity      `json:"Sensitivity"`
	Start            string           `json:"Start"`
	End              string           `json:"End"`
	IsAllDayEvent    interface{}      `json:"IsAllDayEvent"`
	FreeBusyType     FreeBusyType     `json:"FreeBusyType"`
	CalendarItemType CalendarItemType `json:"CalendarItemType"`
	Location         Location         `json:"Location"`
}

type ID struct {
	ChangeKey string `json:"ChangeKey"`
	ID        string `json:"Id"`
}

type Location struct {
	DisplayName   string        `json:"DisplayName"`
	PostalAddress PostalAddress `json:"PostalAddress"`
	IDType        Type          `json:"IdType"`
	LocationType  Type          `json:"LocationType"`
}

type PostalAddress struct {
	Street         interface{}    `json:"Street"`
	City           interface{}    `json:"City"`
	State          interface{}    `json:"State"`
	Country        interface{}    `json:"Country"`
	PostalCode     interface{}    `json:"PostalCode"`
	PostOfficeBox  interface{}    `json:"PostOfficeBox"`
	Type           TypeEnum       `json:"Type"`
	LocationSource LocationSource `json:"LocationSource"`
}

type Header struct {
	ServerVersionInfo ServerVersionInfo `json:"ServerVersionInfo"`
}

type ServerVersionInfo struct {
	MajorVersion     int64  `json:"MajorVersion"`
	MinorVersion     int64  `json:"MinorVersion"`
	MajorBuildNumber int64  `json:"MajorBuildNumber"`
	MinorBuildNumber int64  `json:"MinorBuildNumber"`
	Version          string `json:"Version"`
}

type CalendarItemType string

const (
	Occurrence CalendarItemType = "Occurrence"
	Single     CalendarItemType = "Single"
)

type FreeBusyType string

const (
	Busy      FreeBusyType = "Busy"
	Tentative FreeBusyType = "Tentative"
)

type Type string

const (
	Unknown Type = "Unknown"
)

type LocationSource string

const (
	None LocationSource = "None"
)

type TypeEnum string

const (
	Home TypeEnum = "Home"
)

type Sensitivity string

const (
	Normal Sensitivity = "Normal"
)

type ItemType string

const (
	CalendarItemExchange ItemType = "CalendarItem:#Exchange"
)
