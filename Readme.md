# Weather Application

This is a simple Go application that fetches weather data from an API and displays relevant weather information. The application demonstrates how to make HTTP requests, parse JSON responses, and present weather details.

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [Conclusion](#conclusion)

## Overview

The application follows these main steps:
1. Makes an HTTP GET request to a weather API.
2. Reads the JSON response and extracts weather information.
3. Prints current weather details and future forecast hours.

## Getting Started

## Getting Started

1. Clone the repository:

   ```sh
   git clone https://github.com/onoja123/Go-cli
   cd weather-app

## Usage

To run this application, execute:

```bash
go run main.go
```


The application fetches weather data from a specified API and displays the current weather information along with future forecast hours.

## Dependencies

The application uses the following dependencies:

- `"encoding/json": For JSON encoding and decoding.
- `"fmt": For formatting and printing.
- `"io": For input/output operations.
- `"net/http": For making HTTP requests.
- `"time": For working with time-related operations.
- `"github.com/fatih/color": For colored output (install using go get github.com/fatih/color).


You might also need to install additional packages (like the color package for colored output) if you decide to use them

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open a pull request.

## Conclusion 
The Weather Application project serves as an example that covers several important concepts in Go programming. Throughout this project, you've gained experience and understanding in the following areas: 

1.  **HTTP Requests.**
2. **JSON Parsing**
3. **Structs and Nested Structs**
4. **Error Handling**
5. **Time Handling**

