# Problem Statement: Weather API Data Retrieval and Processing

 - Create a program that interacts with a weather API to fetch and process weather information based on user input. The API used for this task is https://jsonmock.hackerrank.com/api/weather/search.


# Code Intuition:

## Structs:

 - WeatherInfo struct represents individual weather information with fields like Name, Weather, and Status.
 - WeatherResponse struct represents the overall structure of the API response, containing metadata (PageNo, TotalPage) and an array of WeatherInfo.

## Global Variables:

 - ansWeathers is a global slice used to store processed weather information.

## Function: 
 ### PopulateAnsStruct:

    - Takes an array of WeatherInfo and converts relevant data into integers.
    Populates ansWeathers with arrays containing Name, Weather, Wind, and Humidity.
  ### fetchApi:

    - Takes an API URL as input, fetches data from the URL, and populates ansWeathers using PopulateAnsStruct.
    Handles HTTP errors and reads the response body.

 ### Function: main:

    - Asks the user for a string (presumably a city name) to search for weather information.
    Constructs an API URL based on the input.
    Retrieves initial weather data and then iterates through additional pages if there are more results.

## Output:
 - Prints the length of the resultant slice containing weather info
 - Prints the contents of the  finalWeather as each element is an independent unit