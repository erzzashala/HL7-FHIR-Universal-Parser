# HL7/FHIR Parser

This project provides a parser to convert and process HL7 and FHIR patient data into a unified format.

## Overview

The parser can handle messages in both HL7 and FHIR formats, parsing patient information such as:
- ID
- First name
- Last name
- Date of birth
- Gender
- Address

The project aims to provide a seamless solution to convert HL7 messages into FHIR format and vice versa, focusing on patient information exchange.

## Features

- **HL7 Parsing**: Extracts relevant data from HL7 messages and transforms it into a structured format.
- **FHIR Parsing**: Processes FHIR JSON messages and maps them into a standard patient model.
- **Unified Patient Model**: Uses a common `Patient` model to represent patient data across both formats.
- **Error Handling**: Implements robust error handling with custom error responses to ensure secure and valid data exchange.
- **Unit Tests**: Includes unit tests for some cases for HL7 and FHIR parsing

## Installation

To install and run the project, follow the steps below:

### Prerequisites

- **Go**: Ensure that Go (1.18 or later) is installed on your system.

### Install Dependencies

```bash
go mod tidy
```

### Running the Server

Start the server to handle parsing requests.

```bash
go run cmd/server.go
```

This will start a local server (default port 8080) to listen for HL7 and FHIR parsing requests.

### Testing

Run unit tests for HL7 and FHIR parsing:

```bash
go test ./...
```
