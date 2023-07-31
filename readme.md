# BNCAPIDemo

BNCAPIDemo is a demonstration application that interacts with the Brave New Coin API to retrieve asset data, access tokens, and asset ticker data.

## Prerequisites

Before running the BNCAPIDemo application, you need to have the following dependencies installed:

- Go programming language (version 1.16 or later)

## Building the Application

To build the BNCAPIDemo application, follow these steps:

1. Clone the repository to your local machine:
git clone https://github.com/your-username/bncapidemo.git
cd bncapidemo

2. Build the application:
go build -o bncapidemo
This will generate an executable file named `bncapidemo` on your current platform.

## Running the Application

To run the BNCAPIDemo application, use the following command:

This will generate an executable file named `bncapidemo` on your current platform.

## Running the Application

To run the BNCAPIDemo application, use the following command:
./bncapidemo

By default, the application will start the server on port 8080.

## API Endpoints

The BNCAPIDemo application provides the following API endpoints:

1. `/assets`: Retrieves a list of assets based on optional query parameters such as status, symbol, and type.

   Example Usage: `http://localhost:8080/assets?status=ACTIVE`

2. `/token`: Fetches the access token needed to make authenticated API calls.

   Example Usage: `http://localhost:8080/token`

3. `/assetTicker`: Retrieves asset ticker data based on optional query parameters such as symbol and type.

   Example Usage: `http://localhost:8080/assetTicker?symbol=BTC&type=CRYPTO`


## Contributors

- Your Name (Your Email)
