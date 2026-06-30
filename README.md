
# MCP Mood Quote Server

This is a simple Model Context Protocol (MCP) server that returns an inspirational quote based on the user’s mood. It can be used with large language models (LLMs) such as Claude, or tested locally using tools like `curl`.

## What It Does

Send a POST request to the server with a mood (e.g., "happy", "sad", "tired") and it returns a quote associated with that emotion.

## Installation

1. Clone this repository:
   git clone https://github.com/computerscienceiscool/mcp-mood-quotes.git

2. Navigate to the project directory:
   cd mcp-mood-quotes

3. Run the server:
   go run main.go

## Usage

Test locally using curl:

curl -X POST -H "Content-Type: application/json" \
  -d '{"tool":"mcp-mood-quotes","input":{"mood":"happy"}}' \
  http://localhost:8080/messages/

Expected output:

data: {"output":{"quote":"Happiness is not something ready made. It comes from your own actions. — Dalai Lama"}}

## Supported Moods

- happy
- sad
- tired
- excited
- angry

Other moods will return a default message indicating no quote is available.

## Integration with Nanda and Claude

This server is compatible with the Model Context Protocol and may be registered with the Nanda Registry for broader usage.

## Project Structure

mcp-mood-quotes/
├── main.go
└── README.md

## Author

Created by @computerscienceiscool
