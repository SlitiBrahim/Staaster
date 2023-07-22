# Sasteer

Sasteer is a simple web application for bootstrapping a Minimum Viable Product (MVP) Software as a Service (SaaS) project. The application includes features like user registration, authentication, admin privileges, Stripe monthly payments, and the ability to add custom Go code. It also supports sending emails through an SMTP server setup.

## 🎁 Features

- 🔒 User Registration and Authentication
- 👑 Admin Privileges
- 💳 Stripe Monthly Payments
- 🛠️ Custom Go Code Integration
- 📧 SMTP Email Notifications

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

- Go 1.19 or later
- An SMTP server for sending emails

### Installation

1. Clone the repo:

   ```sh
   git clone https://github.com/slitibrahim/sasteer.git
   ```

2. Navigate to the project directory::

   ```sh
   cd sasteer
   ```

3. Download the required Go packages:

   ```sh
   go mod download
   ```

### Usage

To start the server:

   ```sh
   go run .
   ```

To build the project:

   ```sh
   go build
   ```

And that's it! A web server will be started with the following routes:

```ssh
2023/07/23 01:04:42 Server started at http://127.0.0.1:8090
 ➜ REST API: http://127.0.0.1:8090/api/
 ➜ Admin UI: http://127.0.0.1:8090/_/
```