# Identicon Server

This is a simple server written in Go that generates identicons for user avatars. Identicons are visual representations of a hash value, often used to represent users or entities in an application.

## Installation

To install the server, simply clone this repository to your local machine:

```
git clone https://github.com/kuma/identicon-server.git
```


## Usage

### Starting the Server

To start the server, run the following command:

```sh
go run main.go
```
This will start the server on port 8080. You can access the server by visiting http://localhost:8080 in your web browser.

### Generating Identicons

To generate an identicon, make a GET request to the / endpoint with the value you would like to create with png extension.
The server will generate an identicon based on the hash value of the name parameter and return the image data in PNG format.

Here's an example request using the curl command:

```
curl http://localhost:8080/kuma.png
```

This will generate an identicon for the name "kuma".

## Contributing
If you'd like to contribute to this project, feel free to submit a pull request. We welcome contributions of all kinds, including bug fixes, feature requests, and documentation improvements.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/kuma/identicon-server/blob/main/LICENSE) file for details.