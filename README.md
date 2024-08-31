# Doom 🐙

## Overview

This is an educational project designed to simulate a Distributed Denial of Service (DDoS) attack for learning purposes. It is intended for use in controlled environments only and should never be used for malicious activities. 

## Purpose

The primary goal of this project is to help users understand the mechanics of DDoS attacks, including how they work and their impact on target systems. It is meant as a learning tool to demonstrate concepts related to network stress testing and concurrency in Go.

## Features

- **Simulate DDoS Attacks**: Allows you to simulate a DDoS attack on a given URL with a specified number of workers.
- **Control and Monitor**: Start and stop attacks, and monitor success and total request counts.
- **Command-Line Interface (CLI)**: Provides a simple CLI for configuring the URL, number of workers, and attack duration.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/ginozza/doom.git
   cd doom
   ```

2. **Install Dependencies:** Make sure you have Go installed. Install necessary dependencies using:
    ```bash
    cd mod tidy
    ```

3. **Build the Project:** Build the executable with:
    ```bash
    cd ddos_cli
    go build -o doom
    ```

## Usage

To run the simulation, use the CLI with the following options:

```bash
./doom -url <URL> -workers <NUMBER_OF_WORKERS> -duration <DURATION_IN_SECONDS>
```

Example:

```bash
./doom -url http://127.0.0.1:8080 -workers 100 -duration 30
```

This command will start a DDoS simulation on http://127.0.0.1:8080 with 100 workers for 30 seconds.

## Testing

To run the unit tests, use:

```bash
go test -v ./...
```

Tests include verifying the creation of DDoS structures, handling of invalid inputs, and basic functionality checks.

## Important Notes

- **Controlled Environment:** This project is intended for educational purposes and should be run only in a controlled, non-production environment.

- **Ethical Use:** The knowledge gained from using this tool should be used ethically and responsibly. Misuse of DDoS simulations for unauthorized access or attacks is illegal and unethical.

## Contributing

I'm a newbie and student, If you have suggestions or improvements, please submit a pull request or open an issue, no rules for that.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/ginozza/doom/blob/main/LICENSE) file for details.

## Acknowledgments

Inspired by various educational resources on network security and Go programming.

## Disclaimer

**Disclaimer:** This tool is designed for educational purposes only. The author and contributors are not responsible for any misuse of the information or tools provided.