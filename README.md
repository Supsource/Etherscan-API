# Etherscan API Wrapper in Go
This project is an API in Go that allows users to query and filter data from the [Etherscan](https://etherscan.io) API. The API retrieves ERC20 token transfer data by address and stores it in-memory for efficient filtering and retrieval.
## Requirements
* [Go](https://go.dev/dl/) 1.16 or higher
* [Gorilla Mux](github.com/gorilla/mux) package
* Etherscan API key, [get yours](https://etherscan.io/apis)

## Installation
1. Clone the repository: `git clone https://github.com/Supsource/Etherscan-API-Wrapper`
2. Install dependencies: `go mod download`
3. Set your Etherscan API key as an environment variable: `export ETHERSCAN_API_KEY=your-api-key`

## Usage
To start the API server, run the following command in your terminal:
```
go run main.go
```
This will start the server on `localhost:8001`. You can now send requests to the API.

## Endpoints
* `GET /transactions`: Returns all transactions stored in the in-memory dataset.
* `GET /transactions?limit=10&offset=20`: Returns up to 10 transactions starting from the 20th transaction in the dataset.
* `GET /transactions?from=0x123&to=0x456`: Returns all transactions where the from address matches `0x123` and the to address matches `0x456`.
* `GET /transactions?value=1000000000000000000`: Returns all transactions where the value is greater than or equal to 1 ether.

## Response format
The API returns JSON-formatted data in the following format:
```json
{
  "status": "success",
  "data": [
    {
      "hash": "0x123...",
      "blockNumber": "123456",
      "timeStamp": "1620164051",
      "from": "0xabc...",
      "to": "0xdef...",
      "value": "1000000000000000000"
    },
    {
      "hash": "0x456...",
      "blockNumber": "123457",
      "timeStamp": "1620164052",
      "from": "0xghi...",
      "to": "0xjkl...",
      "value": "2000000000000000000"
    }
  ]
}
```

## Error handling
If the API encounters an error, it will return an error response in the following format:
```json
{
  "status": "error",
  "message": "Error message goes here"
}
```
## Contributing
If you find a bug or have a feature request, please open an issue on the GitHub repository. Pull requests are also welcome!
