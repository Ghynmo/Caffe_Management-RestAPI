package thirdparty

import (
	"encoding/json"
	"io/ioutil"
	"miniProject/business/menus"
	"net/http"
)

func GetMenusAPI(key string) (menus.Domain, error) {
	link := "https://masak-apa.tomorisakura.vercel.app/api/recipe/" + key
	response, err_api := http.Get(link)

	if err_api != nil {
		return menus.Domain{}, err_api
	}

	responseData, err_read := ioutil.ReadAll(response.Body)

	if err_read != nil {
		return menus.Domain{}, err_read
	}

	defer response.Body.Close()
	var menu MasakApi
	json.Unmarshal(responseData, &menu)
	return menu.ToDomain(), nil
}
