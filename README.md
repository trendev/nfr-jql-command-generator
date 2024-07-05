# nfr-jql-command-generator

A tool to generate JQL commands for managing Non-Functional Requirements (NFR) in Jira. This script reads NFR labels from a JSON file and creates a CSV file with corresponding JQL commands, enabling efficient NFR tracking and management in Jira.

## Features

- Reads NFR labels from a JSON file.
- Sorts the NFR labels alphabetically.
- Generates JQL commands for each NFR label.
- Exports the JQL commands to a CSV file.

## Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

## Getting Started

### 1. Clone the Repository

```sh
git clone https://github.com/trendev/nfr-jql-command-generator.git
cd nfr-jql-command-generator
```

### 2. Run the Go Program

Make sure you have the `labels.json` file in the root directory of the project, then run the following command:

```sh
go run main.go
```

### 3. Check the Output

After running the program, a file named `NFR_JQL_Scripts.csv` will be created in the root directory containing the JQL commands for each NFR label.
