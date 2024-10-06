# bb2gh

Bitbucket to Github is a simple git migration tool written in Go.

## Getting Started

You will need to have Bitbucket and Github accounts/organizations to use this tool. You will need to have ssh access to both accounts. Finally, the github cli tool will need to be installed on your machine.

### Prerequisites

Requirements for the software and other tools to build, test and push 
- Bitbucket
- Github
- Github CLI
- SSH Access to both Bitbucket and Github

### Installing from source

You can use `go build` to build the binary from source.

```bash
go build -o bb2gh
```

## Usage

### Flags

- `-f` or `--file` - The file path to the `config.yaml` file.
- `gc` or `--generate-config` - Generate a `config.yaml` file with sample data.
- `h` or `--help` - Display the help menu.

### Example

```bash
./bb2gh -f config.yaml
```
### General Troubleshooting

If you are having pulling or pushing repositories, ensure you have the correct permissions in both Github and Bitbucket. Ensure you are properly authenticated with the Github CLI tool. 

