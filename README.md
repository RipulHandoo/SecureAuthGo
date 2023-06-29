# SecureAuthGo

SecureAuthGo is a powerful and secure authentication system built with GoLang. It provides a robust foundation for implementing user authentication and access control in your Go applications.

## Features

- User registration and login functionality
- Password hashing and salting for enhanced security
- Session management and authentication middleware
- Access control and role-based permissions
- Account recovery and password reset
- Integration with popular databases

## Getting Started

To get started with SecureAuthGo, follow these steps:

1. Clone the repository: `git clone https://github.com/RipulHandoo/SecureAuthGo.git`
2. Install the necessary dependencies: `go get -d ./...`
3. Configure the application settings in `config.yml`
4. Set up the database:
   - Create a PostgreSQL database
   - Run the database migration using [Goose](https://github.com/pressly/goose): `goose up`
   - Alternatively, you can use an ORM tool like [sqlc](https://github.com/kyleconroy/sqlc) to generate Go code from your SQL schema.
5. Build and run the application: `go run main.go`
6. Access the application in your web browser: `http://localhost:8080`

For detailed installation instructions and usage guidelines, refer to the [Documentation](docs/README.md).

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request. Make sure to follow our [Contribution Guidelines](CONTRIBUTING.md).

## License

SecureAuthGo is released under the [MIT License](LICENSE).

## Acknowledgements

We would like to acknowledge the following open-source projects that have inspired and contributed to SecureAuthGo:

- [Goose](https://github.com/pressly/goose) - Database migration tool for Go
- [sqlc](https://github.com/kyleconroy/sqlc) - Generate Go code from SQL

## Contact

For any inquiries or further information, please reach out to the project maintainer:

Ripul Handoo - ripulhandoo1234@gmail.com

