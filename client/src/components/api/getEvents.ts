const url =
	"https://mail.innopolis.ru/owa/calendar/1fbf1ef9b7fb43d1bb69c8c20e41a9c1@innopolis.ru/589470ef18664498a886b3c2d8a1eb6c10478442459916160547/service.svc?action=FindItem&ID=-1&AC=1";

const headers = new Headers({
	Accept: "*/*",
	"Accept-Encoding": "gzip, deflate, br, zstd",
	"Accept-Language": "en-US,en;q=0.9",
	Action: "FindItem",
	"Client-Request-Id": "86B61BDA20D0408FB8B848B92B68102A_172016655142104",
	"Content-Type": "application/json; charset=UTF-8",
	Cookie:
		"X-BackEndCookie=18abd935-8fde-474f-a3aa-1f93f5a87152=u56Lnp2ejJqBz8jPms3HyJ7SzM2Zy9LLns+b0p6enJrSzsnGzsnGz8mZxp7LgYHNz83L0s/H0s/Lq8/IxcrPxc7N&1fbf1ef9-b7fb-43d1-bb69-c8c20e41a9c1=u56Lnp2ejJqBz8jPms3HyJ7SzM2Zy9LLns+b0p6enJrSzsnGzsnGz8mZxp7LgYHNz83L0s/H0s/Lq8/Hxc/NxczP&a4787616-623e-4c25-839a-c7e4566768f3=u56Lnp2ejJqBz8jPms3HyJ7SzM2Zy9LLns+b0p6enJrSzsnGzsnGz8mZxp7LgYHNz83L0s/H0s/Lq8/IxcrNxcrM&0f4f1dd4-0168-402e-9058-bf2993ef1777=u56Lnp2ejJqBz8jPms3HyJ7SzM2Zy9LLns+b0p6enJrSzsnGzsnGz8mZxp7LgYHNz83L0s/H0s/Lq8/IxcrLxc/K&ff320c0c-21bc-4d6d-9320-dcfc7bc7c400=u56Lnp2ejJqBz8jPms3HyJ7SzM2Zy9LLns+b0p6enJrSzsnGzsnGz8mZxp7LgYHNz83L0s/H0s/Lq8/IxcrLxc7H; X-BackEndCookie=S-1-5-21-1518683800-3227197516-3513364682-10911=u56Lnp2ejJqBxs3MzcvOzM7Sx8/MntLLy5uZ0sabypzSz8/Jm8nMns3PmsnKgYHNz83L0s/H0s/Nq8/Jxc3NxczH; ClientId=86B61BDA20D0408FB8B848B92B68102A; X-OWA-JS-PSD=1; PrivateComputer=true; OutlookSession=9d8eab4c303441309f6e216a0fac06b6",
	Origin: "https://mail.innopolis.ru",
	"Sec-Ch-Ua": '"Not/A)Brand";v="8", "Chromium";v="126", "Brave";v="126"',
	"Sec-Ch-Ua-Mobile": "?0",
	"Sec-Ch-Ua-Platform": '"Linux"',
	"Sec-Fetch-Dest": "empty",
	"Sec-Fetch-Mode": "cors",
	"Sec-Fetch-Site": "same-origin",
	"Sec-Gpc": "1",
	"User-Agent":
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	"X-Anchormailbox": "1fbf1ef9b7fb43d1bb69c8c20e41a9c1@innopolis.ru",
	"X-Owa-Actionid": "-1",
	"X-Owa-Actionname": "GetCalendarItemsAction_Month",
	"X-Owa-Attempt": "1",
	"X-Owa-Canary": "X-OWA-CANARY_cookie_is_null_or_empty",
	"X-Owa-Clientbegin": "2024-07-05T08:02:31.421",
	"X-Owa-Clientbuildversion": "15.2.1258.28",
	"X-Owa-Correlationid": "86B61BDA20D0408FB8B848B92B68102A_172016655142104",
});

const body = {
	__type: "FindItemJsonRequest:#Exchange",
	Header: {
		__type: "JsonRequestHeaders:#Exchange",
		RequestServerVersion: "Exchange2013",
		TimeZoneContext: {
			__type: "TimeZoneContext:#Exchange",
			TimeZoneDefinition: {
				__type: "TimeZoneDefinitionType:#Exchange",
				Id: "Russian Standard Time",
			},
		},
	},
	Body: {
		__type: "FindItemRequest:#Exchange",
		ItemShape: {
			__type: "ItemResponseShape:#Exchange",
			BaseShape: "IdOnly",
			AdditionalProperties: [
				{ __type: "PropertyUri:#Exchange", FieldURI: "ItemParentId" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "Sensitivity" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "AppointmentState" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "IsCancelled" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "HasAttachments" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "LegacyFreeBusyStatus" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "CalendarItemType" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "Start" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "End" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "IsAllDayEvent" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "Organizer" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "Subject" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "IsMeeting" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "UID" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "InstanceKey" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "ItemEffectiveRights" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "JoinOnlineMeetingUrl" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "ConversationId" },
				{
					__type: "PropertyUri:#Exchange",
					FieldURI: "CalendarIsResponseRequested",
				},
				{ __type: "PropertyUri:#Exchange", FieldURI: "Categories" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "IsRecurring" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "IsOrganizer" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "EnhancedLocation" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "IsSeriesCancelled" },
				{ __type: "PropertyUri:#Exchange", FieldURI: "Charm" },
			],
		},
		ParentFolderIds: [
			{
				__type: "FolderId:#Exchange",
				Id: "AQMkAGRhNzJmODFkLWEyZDEtNGNkNC1hYQA3My1lOWMxNjNkYjgzOGEALgAAA7IGwULT4DxGjhqS6paOUGMBAHB6D7uh5FtPiONy7WVVDycAAAIBDQAAAA==",
				ChangeKey: "AgAAAA==",
			},
		],
		Traversal: "Shallow",
		Paging: {
			__type: "CalendarPageView:#Exchange",
			StartDate: "2024-06-30T00:00:00.001",
			EndDate: "2024-08-04T00:00:00.000",
		},
	},
};

export default function getUniversityBookingsJSON() {
	fetch(url, {
		method: "POST",
		headers: headers,
		body: JSON.stringify(body),
	})
		.then((response) => response.json())
		.then((data) => console.log(data))
		.catch((error) => console.error("Error:", error));
}
