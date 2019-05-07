# Google Search Results in Golang

This Golang package allows you to scrape and parse Google Search Results using [SerpWow](https://serpwow.com). In addition to [Search](https://serpwow.com/docs/search-api/overview) you can also use this package to access the SerpWow [Locations API](https://serpwow.com/docs/locations-api/overview), [Batches API](https://serpwow.com/docs/batches-api/overview) and [Account API](https://serpwow.com/docs/account-api).

The package requires Golang 1.8 or above, the package has no dependancies.

## Documentation
We have included examples here but full SerpWow API documentation is available at the [API Docs](https://serpwow.com/docs):
- [Search API Docs](https://serpwow.com/docs/search-api/overview) 
- [Locations API Docs](https://serpwow.com/docs/locations-api/overview) 
- [Account API Docs](https://serpwow.com/docs/account-api)
- [Batches API Docs](https://serpwow.com/docs/batches-api)

You can also use the [API Playground](https://app.serpwow.com/playground) to visually build Google search requests using SerpWow.

## Examples
- [Simple Example](#simple-example) 
- [Example Response](#example-response) 
- [Getting an API Key](#getting-an-api-key)
- [Searching with a Location](#searching-with-a-location)
- [Searching Google Places, Google Videos, Google Images, Google Shopping and Google News](#searching-google-places-google-videos-google-images-google-shopping-and-google-news)
- [Returning results as JSON, HTML and CSV](#returning-results-as-json-html-and-csv)
- [Requesting mobile and tablet results](#requesting-mobile-and-tablet-results)
- [Parsing Results](#parsing-results)
- [Paginating results, returning up to 100 results per page](#paginating-results-returning-up-to-100-results-per-page)
- [Search example with all parameters](#search-example-with-all-parameters)
- [Locations API Example](#locations-api-example)
- [Account API Example](#account-api-example)
- [Batches API](#batches-api)

## Simple Example
Simplest example for a standard query "pizza", returning the Google SERP (Search Engine Results Page) data as JSON.
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // set up our parameters
	parameters := map[string]string{
		"q": "pizza",
  }
  
  // run the search, requesting results in JSON format
  response, error := serpwow.GetJSON(parameters, apiKey)

	// print the response, or error, if one occurred
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
	}
}
```

## Example Response
A snapshot (shortened for brevity) of the JSON response returned is shown below. For details of all of the fields from the Google search results page that are parsed please see the [docs](https://serpwow.com/docs/search-api/results/overview).
```json
{
  "request_info": {
    "success": true
  },
  "search_metadata": {
    "id": "20c8e44e9cacedabbdff2d9b7e854436056d4f33",
    "google_url": "https://www.google.com/search?q=pizza&oq=pizza&sourceid=chrome&ie=UTF-8",
    "total_time_taken": 0.14
  },
  "search_parameters": {
    "q": "pizza"
  },
  "search_information": {
    "total_results": 1480000000,
    "time_taken_displayed": 0.45,
    "query_displayed": "pizza",
    "detected_location": "Ireland"
  },
  "local_map": {
    "link": "https://www.google.com/search?q=pizza&npsic=0&rflfq=1&rldoc=1&rlha=0&rllag=53350059,-6249133,1754&tbm=lcl&sa=X&ved=2ahUKEwiC3cLZ0JLgAhXHUxUIHQrsBC4QtgN6BAgBEAQ",
    "gps_coordinates": {
      "latitude": 53.350059,
      "longitude": -6.249133,
      "altitude": 1754
    },
    "local_results": [{
        "position": 1,
        "title": "Apache Pizza Temple Bar",
        "extensions": [
          "American-style pizza-delivery chain"
        ],
        "rating": 3.6,
        "reviews": 382,
        "type": "Pizza",
        "block_position": 1
      }
    ]
  },
  "knowledge_graph": {
    "title": "Pizza",
    "type": "Dish",
    "description": "Pizza is a savory dish of Italian origin, consisting of a usually round, flattened base of leavened wheat-based dough topped with tomatoes, cheese, and various other ingredients baked at a high temperature, traditionally in a wood-fired oven.",
    "source": {
      "name": "Wikipedia",
      "link": "https://en.wikipedia.org/wiki/Pizza"
    },
    "nutrition_facts": {
      "total_fat": [
        "10 g",
        "15%"
      ],
      "sugar": [
        "3.6 g"
      ]
    }
  },
  "related_searches": [{
      "query": "apache pizza",
      "link": "https://www.google.com/search?q=apache+pizza&sa=X&ved=2ahUKEwiC3cLZ0JLgAhXHUxUIHQrsBC4Q1QIoAHoECAUQAQ"
    }
  ],
  "organic_results": [{
      "position": 1,
      "title": "10 Best Pizzas in Dublin - A slice of the city for every price point ...",
      "link": "https://www.independent.ie/life/travel/ireland/10-best-pizzas-in-dublin-a-slice-of-the-city-for-every-price-point-37248689.html",
      "domain": "www.independent.ie",
      "displayed_link": "https://www.independent.ie/.../10-best-pizzas-in-dublin-a-slice-of-the-city-for-every-p...",
      "snippet": "Oct 20, 2018 - Looking for the best pizza in Dublin? Pól Ó Conghaile scours the city for top-notch pie... whatever your budget.",
      "cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:wezzRov42dkJ:https://www.independent.ie/life/travel/ireland/10-best-pizzas-in-dublin-a-slice-of-the-city-for-every-price-point-37248689.html+&cd=4&hl=en&ct=clnk&gl=ie",
      "block_position": 2
    }
  ],
  "related_places": [{
    "theme": "Best dinners with kids",
    "places": "Pinocchio Italian Restaurant - Temple Bar, Cafe Topolisand more",
    "images": [
      "https://lh5.googleusercontent.com/p/AF1QipNhGt40OpSS408waVJUHeItGrrGqImmEKzuVbBv=w152-h152-n-k-no"
    ]
  }],
  "pagination": {
    "current": "1",
    "next": "https://www.google.com/search?q=pizza&ei=fRZQXMKqL8en1fAPitiT8AI&start=10&sa=N&ved=0ahUKEwiC3cLZ0JLgAhXHUxUIHQrsBC4Q8NMDCOkB"
  }
}
```

## Getting an API Key
To get a free API Key head over to [app.serpwow.com/signup](https://app.serpwow.com/signup).

## Searching with a location
Example of a Google query geo-locating the query as if the user were located in New York. 
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // set up our parameters
	parameters := map[string]string{
    "q": "pizza",
    "location": "New York,New York,United States",
  }
  
  // run the search, requesting results in JSON format
  response, error := serpwow.GetJSON(parameters, apiKey)

	// print the response, or error, if one occurred
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
	}
}
```

## Searching Google Places, Google Videos, Google Images, Google Shopping and Google News
Use the ``search_type`` param to search Google Places, Videos, Images and News. See the [Search API Parameters Docs](https://serpwow.com/docs/search-api/searches/parameters) for full details of the additional params available for each search type.
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // perform a search on Google News, just looking at blogs, ordered by date, in the last year, filtering out duplicates
	parameters := map[string]string{
    "q": "football news",
    "search_type": "news",
    "news_type": "blogs",
    "sort_by": "date",
    "time_period": "last_year",
    "show_duplicates": "false",
  }
  response, error := serpwow.GetJSON(parameters, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
  }

  // perform a search on Google Places for "plumber" in London
  parameters = map[string]string{
    "search_type": "places",
    "q": "plumber",
    "location": "London,England,United Kingdom",
  }
  response, error := serpwow.GetJSON(parameters, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
  }

  // perform an image search on Google Images for "red flowers"
  parameters = map[string]string{
    "q" : "red flowers",
    "search_type" : "images",
  }
  response, error := serpwow.GetJSON(parameters, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
  }
}
```

## Returning results as JSON, HTML and CSV
SerpWow can return data in JSON, HTML and CSV formats using the ``get_json``, ``get_html`` and ``get_csv`` methods. For CSV results use the ``csv_fields`` param ([docs](https://serpwow.com/docs/search-api/searches/csv-fields)) to request specific result fields.
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // set up our parameters
	parameters := map[string]string{
    "q": "pizza",
    "location": "New York,New York,United States",
  }
  
  // retrieve the Google search results as JSON
  responseJson, error := serpwow.GetJSON(parameters, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(responseJson)
  }
  
  // retrieve the Google search results as HTML
  responseHtml, error := serpwow.GetHTML(parameters, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(responseHtml)
  }
  
  // retrieve the Google search results as CSV
  responseCsv, error := serpwow.GetCSV(parameters, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(responseCsv)
	}
}
```

## Requesting mobile and tablet results
To request that SerpWow renders the Google search results via a mobile or tablet browser use the ``device`` param:
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // set up the mobile params
	parametersMobile := map[string]string{
		"q" : "pizza",
    "device" : "mobile",
  }

  // set up the tablet params
	parametersTablet := map[string]string{
		"q" : "pizza",
    "device" : "tablet",
  }

  // set up the desktop params (note we omit the "device" param)
	parametersDesktop := map[string]string{
		"q" : "pizza",
  }

  // retrieve the mobile search results
  responseMobile, error := serpwow.GetJSON(parametersMobile, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(responseMobile)
  }

  // retrieve the tablet search results
  responseTablet, error := serpwow.GetJSON(parametersTablet, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(responseTablet)
  }

  // retrieve the desktop search results
  responseDesktop, error := serpwow.GetJSON(parametersDesktop, apiKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(responseDesktop)
  }
}
```

## Parsing results
When making a request via the ``get_json`` method a standard Python ``dict`` is returned. You can inspect this dict to iterate, parse and store the results in your app.
```go
package main

import "fmt"
import "strings"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // make a simple query, returning JSON
	parameters := map[string]string{
		"q": "pizza",
  }
	response, error := serpwow.GetJSON(parameters, apiKey)
	
	// determine if the request was successful
	if error != nil {
		fmt.Println(error)
	} else {

		// extract the time taken and number of organic results
		timeTaken := fmt.Sprintf("%f", response["search_metadata"].(map[string]interface{})["total_time_taken"])
		organicResultCount := fmt.Sprintf("%v", len(response["organic_results"].([]interface{})))

		// print results
		s := []string{organicResultCount, "organic results returned in", timeTaken, "\n"};
    fmt.Printf(strings.Join(s, " "));
	}
}
```

## Paginating results, returning up to 100 results per page
Use the ``page`` and ``num`` parameters to paginate through Google search results. The maximum number of results returned per page (controlled by the ``num`` param) is 100 (a Google-imposed limitation) for all ``search_type``'s apart from Google Places, where the maximum is 20. Here's an example.
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // request the first 100 results
	parameters := map[string]string{
    "q" : "pizza",
    "page" : 1,
    "num": 100
  }
  responsePage1, errorPage1 := serpwow.GetJSON(parameters, apiKey)
  if errorPage1 != nil {
		fmt.Println(errorPage1)
	} else {
		fmt.Println(responsePage1)
	}

  // request the next 100 results
  parameters["page"] = 2
  responsePage2, errorPage2 := serpwow.GetJSON(parameters, apiKey)
  if errorPage2 != nil {
		fmt.Println(errorPage2)
	} else {
		fmt.Println(responsePage2)
	}
}
```

## Search example with all parameters
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // set up query parameters, retrieving results as CSV (note the csv_fields param)
	parameters := map[string]string{
    "q" : "pizza",
    "gl" : "us",
    "hl" : "en",
    "location" : "New York,New York,United States",
    "google_domain" : "google.com",
    "time_period" : "custom",
    "sort_by" : "date",
    "time_period_min" : "02/01/2018",
    "time_period_max" : "02/08/2019",
    "device" : "mobile",
    "csv_fields" : "search.q,organic_results.position,organic_results.domain",
    "page" : "1",
    "num" : "100",
  }
  
  // retrieve the search results as CSV
  response, error := serpwow.GetCSV(parameters, apiKey)

	// print the response, or error, if one occurred
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
	}
}
```

## Locations API Example
The [Locations API](https://serpwow.com/docs/locations-api/overview) allows you to search for SerpWow supported Google search locations. You can supply the ``full_name`` returned by the Locations API as the ``location`` parameter in a Search API query (see [Searching with a location](https://github.com/serpwow/google-search-results-python#searching-with-a-location) example above) to retrieve search results geo-located to that location.
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // # set up the parameters for the location query parameters
	parameters := map[string]string{
    "q" : "mumbai",
  }

  // retrieve locations matching the query parameters as JSON
  response, error := serpwow.GetLocations(parameters, apiKey)

	// print the response, or error, if one occurred
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
	}
}
```

## Account API Example
The [Account API](https://serpwow.com/docs/account-api) allows you to check your current SerpWow usage and billing information. 
```go
package main

import "fmt"
import "serpwow"

func main() {
  // set our API key
  apiKey := "demo"

  // get our account info
  response, error := serpwow.GetAccount(apiKey)

	// print the response, or error, if one occurred
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
	}
}
```

## Batches API
The [Batches API](https://serpwow.com/docs/batches-api) allows you to create, update and delete Batches on your SerpWow account (Batches allow you to save up to 15,000 Searches and have SerpWow run them on a schedule).

For more information and extensive code samples please see the [Batches API Docs](https://serpwow.com/docs/batches-api).