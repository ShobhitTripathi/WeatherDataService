### Installation
Make sure you have Go installed on your machine. 
If not, you can download and install it from the official [Go website](official Go website).

Install Redis on your machine. You can download Redis from the official Redis website or use a package manager like Homebrew (for macOS) or apt (for Ubuntu).

Clone this repository to your local machine:
````
git clone https://github.com/your-username/go-redis-cache-example.git
````
Change into the project directory:
````
cd WeatherDataService
````
Install dependencies:
````
go mod tidy
````

## Usage
Start the Redis server on your local machine.

Run the Go application:
````
go run main.go
````

Open your web browser and navigate to

``````
To get Current Weather Data
http://localhost:8081/api/weather/your-zipcode 
``````

``````
To get Current Weather Forecast
http://localhost:8081/api/weather/your-zipcode/days
``````
to see the weather forecast for the specified zipcode. The application will fetch the weather data from the Redis cache if available; otherwise, it will fetch it from the weather API and store it in the cache for future use.


## Limitations
- This project is a basic example and may not cover all edge cases or production-level scenarios.
- The caching mechanism implemented here is simple and may not be suitable for high-traffic or mission-critical applications.
- The API supports only US zip e.g.: q=10001
- Days can only be between 1- 14
- API Key is valid till 2nd of April 2024

## Configuration
You can configure the Redis connection settings in the config/config.go file. Update the RedisHost and RedisPort constants to match your Redis server configuration.
````
const (
    RedisHost = "localhost"
    RedisPort = 6379
)
````
