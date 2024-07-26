# weather_app

<!-- GETTING STARTED -->
## Getting Started
This is an example of how you may give instructions on setting up to run project locally.

### Prerequisites
With in the config folder we have config.go and we need place our open weather API Key in that place to connect to open weather API Client.
weather_app/config/config.go
   ```sh
    "apiKey": "add_your_openweather_api_key_here"
   ```

### Command to run your go application
Open command prompt and navigate to the root path of your project which is weather_app and then run below command to run your application.
   ```sh
    go run main.go
   ```

### Testing your go application
Use any API testing tools like postman or insomnia to test the endpoint. Below is the sample endpoint you can use to test your API.
   ```sh
    curl --location 'http://localhost:8080/weather'
    curl --location 'http://localhost:8080/weather?lat=40&lon=74&units=metric'
   ```