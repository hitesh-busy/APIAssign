package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

var ansWeathers []interface{}

type WeatherInfo struct {
	Name    string   `json:"name"`
	Weather string   `json:"weather"`
	Status  []string `json:"status"`
}

type WeatherResponse struct {
	PageNo    int           `json:"page"`
	TotalPage int           `json:"total_pages"`
	Weathers  []WeatherInfo `json:"data"`
}

func PopulateAnsStruct(weather []WeatherInfo) error {

	for _, v := range weather {

		//make an array for rach
		var temp []interface{}
		//\d: Matches any digit (0-9).
		//+: Matches one or more occurrences of the preceding element (in this case, \d).
		pattern := regexp.MustCompile(`\d+`) //fix the pattern for int values

		weather, _ := strconv.Atoi(pattern.FindString(v.Weather))
		wind, _ := strconv.Atoi(pattern.FindString(v.Status[0]))
		humidity, _ := strconv.Atoi(pattern.FindString(v.Status[1]))

		temp = append(temp, v.Name)
		temp = append(temp, weather)
		temp = append(temp, wind)
		temp = append(temp, humidity)

		ansWeathers = append(ansWeathers, temp)
	}
	return nil
}
func fetchApi(apiUrl string) error {
	response, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New("Status not 200")
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	//unMarshalling to extract the json data
	var data WeatherResponse
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return err
	}

	//populating all the weathers in the "data " field of response body
	PopulateAnsStruct(data.Weathers)
	return nil
}

func main() {
	fmt.Println("Weather API. Please enter your desired string to be searched ")
	var name string
	fmt.Scanln(&name)

	apiUrl := fmt.Sprintf("https://jsonmock.hackerrank.com/api/weather/search?name=%s", name)

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("error fetching intially ", err.Error())
	}
	if response.StatusCode != 200 {
		fmt.Println("Status not 200")
		return
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var data WeatherResponse
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		fmt.Println("error Unmarshalling ", err.Error())
	}

	for i := 1; i <= data.TotalPage; i++ {
		newUrl := fmt.Sprintf("%s&page=%d", apiUrl, i)
		err := fetchApi(newUrl)
		if err != nil {
			fmt.Println("error fetching  API", err.Error())
		}
	}

	fmt.Println("Length of the final ans Slice is ", len(ansWeathers))
	fmt.Println("final Ans is \n", ansWeathers)

}
