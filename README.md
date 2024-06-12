
# Options Contracts Risk & Reward Analysis Service

This project is a backend service for analyzing the risk and reward of options contracts. The service accepts up to four options contracts as input and returns the data necessary to generate a risk and reward graph. It also provides the maximum possible profit, the maximum possible loss, and all break-even points.

## Features

- Analyze up to four options contracts.
- Generate risk and reward graph data.
- Calculate maximum profit, maximum loss, and break-even points.
- Validate input data to ensure correctness.

## Requirements

- Go 1.16 or later
- Gin Gonic Web Framework
- testify for unit testing

## Project Structure

```
.
├── controllers
│   └── analysisController.go
├── model
│   └── optionsContract.go
├── routes
│   └── routes.go
├── tests
│   └── analysis_test.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/yourproject.git
   cd yourproject
   ```

2. Initialize the Go module:
   ```sh
   go mod init github.com/yourusername/yourproject
   ```

3. Tidy up the Go module to download dependencies:
   ```sh
   go mod tidy
   ```

4. Run the server:
   ```sh
   go run main.go
   ```

## Usage

The server runs on `http://localhost:8080`.

### Analyze Endpoint

- **URL**: `/analyze`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Body**:
  ```json
  [
    {
      "type": "Call",
      "strike_price": 100,
      "bid": 10.05,
      "ask": 12.04,
      "long_short": "long",
      "expiration_date": "2025-12-17T00:00:00Z"
    }
  ]
  ```

- **Response**:
  ```json
  {
    "xy_values": [
      {"x": 50, "y": -12.04},
      {"x": 51, "y": -12.04},
      ...
    ],
    "max_profit": 37.96,
    "max_loss": -12.04,
    "break_even_points": [113]
  }
  ```

## Testing

Run the tests using the following command:
```sh
go test ./tests/analysis_test.go
```

This will run all the unit and integration tests to ensure the service is working correctly.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
