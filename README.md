# Genesis Software Engineering School 2024 API

## Description

This repository contains the code for a simple API developed for the Genesis Software Engineering School 2024.

## Installation

1. Clone the repository: `git clone https://github.com/AndrisSmirnov/exchange-rate-2024.git`
2. Navigate to the project directory: `cd exchange_rate`

## Usage

### Run in Docker

1. Run `docker compose up --build` to build.

### List of endpoints:

- `/api/rate` (GET): get current bitcoin rate in UAH
- `/api/subscribe` (POST): subscribe to mailing list

### Features list:

- Getting exchange rates directly from the [BankGovIA](https://bank.gov.ua/admin_uploads/article/Instr_API_KURS_VAL_data.pdf)
- Added the possibility to expand the service to receive the exchange rate of other currency
- Building the e-mail using HTML templates for better viewing
- The e-mail contains an "Unsubscribe" button.
- The e-mail contains links to all social networks
