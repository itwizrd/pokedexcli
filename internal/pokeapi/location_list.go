package pokeapi

import(
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	//build the GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}
	//execute the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return Locations{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}
	if err != nil {
		return Locations{}, err
	}

	loc := Locations{}
	err = json.Unmarshal(body, &loc)
	if err != nil {
		return Locations{}, err
	}
	return loc, nil
}
