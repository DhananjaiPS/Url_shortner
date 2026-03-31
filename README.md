


# Go URL Shortener

A lightweight, fast, and simple URL shortener API written in Go. This project uses an in-memory data store to map long URLs to short, 8-character hashed IDs, allowing for quick redirection.

## Features
* In-Memory Database: Uses a Go map for lightning-fast URL storage and retrieval.
* MD5 Hashing: Generates unique short URLs by hashing the original URL using MD5 and extracting the first 8 characters.
* JSON API: Accepts and returns JSON data for seamless integration with frontend applications.
* Built-in Redirection: Automatically redirects users from the shortened URL to their original destination.

## Prerequisites
* [Go](https://go.dev/doc/install) (version 1.16 or higher recommended)

## How to Run

1. Clone the repository or save the code in a file named `main.go`.
2. Open your terminal and navigate to the directory containing `main.go`.
3. Run the following command to start the server:
   ```bash
   go run main.go
   ```
4. You should see the following output indicating the server is running:
   ```text
   Welcome to Url shortner
   Server is getting ready on port:3000
   Server is running on port:3000
   ```

## API Endpoints

The server runs locally on `http://localhost:3000`.

### 1. Root Endpoint
A simple endpoint to verify the server is running.
* Method: `GET`
* URL: `/`
* Response: `handler function is called on / route`

### 2. Create a Short URL
Submits a long URL to be shortened.
* Method: `POST`
* URL: `/shortner`
* Headers: `Content-Type: application/json`
* Body:
  ```json
  {
      "url": "https://www.google.com"
  }
  ```
* Success Response (200 OK):
  ```json
  {
      "shorturl": "8ffdefb1"
  }
  ```

cURL Example:
```bash
curl -X POST http://localhost:3000/shortner \
-H "Content-Type: application/json" \
-d '{"url": "https://www.google.com"}'
```

### 3. Redirect to Original URL
Takes the generated short URL and redirects the client to the original destination.
* Method: `GET`
* URL: `/redirect/{shorturl}` 
*(Example: `http://localhost:3000/redirect/8ffdefb1`)*
* Response (302 Found): Automatically redirects your browser to the original URL.

## How it Works (Under the Hood)
1. Hashing: When a URL is submitted, the `generateshortUrl` function creates an MD5 hash of the string, converts the byte slice to a hexadecimal string, and slices the first 8 characters to use as the ID.
2. Storage: The application uses a global map `UrlDb = map[string]Url{}` to store the mappings in memory. *(Note: Because this is an in-memory store, all generated URLs will be lost when the server restarts.)*
3. Handling Requests: The `http` package is used to set up multiplexers (`HandleFunc`) that route incoming HTTP requests to their respective Go functions.
