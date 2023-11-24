# Healthchecker

Healthchecker is a tiny tool written in Go that checks whether a website is running or is down.

## Usage

To use the tool, follow the steps below:

1. Clone the repository: `git clone <repository-url>`
2. Build the Go project: `go build`
3. Run the tool with the following command: `./healthchecker --domain <domain-name> [--port <port-number>]`

### Command-line Arguments

- `--domain, -d` (required): Specifies the domain name to check.
- `--port, -p` (optional, default: 80): Specifies the port number to use.

## Example

To check if example.com is up or down on port 8080, run the command:

```shell
./healthchecker --domain example.com --port 8080

```

This will output the status of the website.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details
