package myuniversity

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/zap"
)

func cancelRedirectHandler(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func getRetriableClient(jar http.CookieJar) *http.Client {
	client := retryablehttp.NewClient()

	return client.StandardClient()
}

func GetMyUniversityCookies(email, password string, log *zap.Logger) ([]*http.Cookie, error) {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: nil})

	// client := getRetriableClient(jar)
	client := &http.Client{}
	client.Jar = jar
	client.CheckRedirect = cancelRedirectHandler

	/////////////////////////////////////////////////////////////////////////

	const myUniURL = "https://my.university.innopolis.ru/"
	resp, err := client.Get(myUniURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request %q: %w", myUniURL, err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Sugar().Errorw("failed to request", "url", myUniURL, "response", resp.Status)

		return nil, fmt.Errorf("failed to request %q: %w", myUniURL, ErrMyUniversityAPIChanged)
	}
	defer resp.Body.Close()

	/////////////////////////////////////////////////////////////////////////

	const myUniAuthURL = "https://my.university.innopolis.ru/site/auth?authclient=adfs"
	respWithRedirect, err := client.Get(myUniAuthURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request %q: %w", myUniAuthURL, err)
	}
	if respWithRedirect.StatusCode != http.StatusFound {
		log.Sugar().
			Errorw("status must be StatusFound", "url", myUniAuthURL, "status", respWithRedirect.Status)

		return nil, fmt.Errorf("%w", ErrMyUniversityAPIChanged)
	}
	defer respWithRedirect.Body.Close()

	/////////////////////////////////////////////////////////////////////////

	ssoURL, err := respWithRedirect.Location()
	if err != nil {
		log.Sugar().Errorw("failed to get redirect url to sso")

		return nil, fmt.Errorf("%w", ErrMyUniversityAPIChanged)
	}

	ssoURLstr := ssoURL.String()

	/////////////////////////////////////////////////////////////////////////

	ssoGetResp, err := client.Get(ssoURLstr)
	if err != nil {
		return nil, fmt.Errorf("failed to request sso: %w", err)
	}
	if ssoGetResp.StatusCode != http.StatusOK {
		log.Sugar().
			Errorw("status must be StatusOK", "url", ssoURLstr, "status", ssoGetResp.Status)

		return nil, fmt.Errorf("%w", ErrMyUniversityAPIChanged)
	}
	defer ssoGetResp.Body.Close()

	/////////////////////////////////////////////////////////////////////////

	body, err := io.ReadAll(ssoGetResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body with button: %w", err)
	}

	re := regexp.MustCompile(`client-request-id=([a-f0-9-]+)`)

	matches := re.FindStringSubmatch(string(body))
	if matches == nil || len(matches) < 1 {
		log.Sugar().Errorw("no client-request-id was found", "url", ssoURLstr)

		return nil, fmt.Errorf("%w", ErrMyUniversityAPIChanged)
	}

	clientRequestID := matches[1]

	queryParams := ssoURL.Query()
	queryParams.Add("client-request-id", clientRequestID)

	ssoURL.RawQuery = queryParams.Encode()

	credentials := fmt.Sprintf(
		"UserName=%s&Password=%s&AuthMethod=FormsAuthentication",
		email,
		password,
	)

	/////////////////////////////////////////////////////////////////////////

	ssoPostResp, err := client.Post(
		ssoURLstr,
		"application/x-www-form-urlencoded",
		strings.NewReader(credentials),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request %q: %w", ssoURLstr, err)
	}
	if ssoPostResp.StatusCode == http.StatusOK {
		return nil, fmt.Errorf("%w", ErrInvalidCredentials)
	}
	if ssoPostResp.StatusCode != http.StatusFound {
		log.Sugar().
			Errorw("status must be StatusFound", "url", myUniAuthURL, "status", ssoPostResp.Status)

		return nil, fmt.Errorf("%w", ErrMyUniversityAPIChanged)
	}

	/////////////////////////////////////////////////////////////////////////

	ssoGetResp, err = client.Get(ssoURLstr)
	if err != nil {
		return nil, fmt.Errorf("failed to request %q: %w", ssoURLstr, err)
	}
	if ssoGetResp.StatusCode != http.StatusFound {
		log.Sugar().
			Errorw("status must be StatusFound", "url", myUniAuthURL, "status", ssoGetResp.Status)

		return nil, fmt.Errorf("%w", ErrMyUniversityAPIChanged)
	}
	defer ssoGetResp.Body.Close()

	/////////////////////////////////////////////////////////////////////////

	myUniRedirectURL, err := ssoGetResp.Location()
	if err != nil {
		log.Sugar().Debugw("failed to get redirect url to sso", "url", myUniAuthURL)

		return nil, fmt.Errorf("%w", ErrMyUniversityAPIChanged)
	}

	/////////////////////////////////////////////////////////////////////////

	myUniRedirectResp, err := client.Get(myUniRedirectURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make GET reqest to myUni after sso: %w", err)
	}
	if myUniRedirectResp.StatusCode != http.StatusOK {
		log.Sugar().
			Errorw("status must be StatusOK", "url", myUniRedirectURL.String(), "status", myUniRedirectResp.Status)

		return nil, fmt.Errorf("%w", ErrMyUniversityAPIChanged)
	}

	/////////////////////////////////////////////////////////////////////////

	myUniveristyUrl, _ := url.ParseRequestURI("https://my.university.innopolis.ru")
	myUniveristyCookies := jar.Cookies(myUniveristyUrl)
	return myUniveristyCookies, nil
}
