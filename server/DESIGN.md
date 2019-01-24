# Design
Server design.

# Table Of Contents
- [Models](#models)
- [Endpoints](#endpoints)

# Models
SQL tables.

All `*id` fields are of type `SERIAL`

## User Model
- `id`
- `username` (`text`)
- `password_hash` (`text`)

## Device Model
- `id`
- `user_id`
- `platform` (`enum`)
    - Values: `mobile.android`, `web.firefox`
- `name` (`text`)

## API Token Model
- `id`
- `user_id`
- `device_id`
- `token_hash` (`text`)

## Push Model
- `id`
- `user_id`
- `from_device_id`
- `text` (`text`)

# Endpoints
API endpoints.  

All data should be JSON encoded.  

Unless specified an API Token should be provided in the 
`Authorization` header.

## User Endpoints
### Create User
#### Request
POST `/api/v0/users`

Authentication not required.

Body:

- `username` (`string`)
- `password` (`string`)

#### Response
Body:

- `user` ([`User`](#user-model))

## Device Endpoints
### Create Device
#### Request
POST `/api/v0/devices`

Body:

- `platform` (`string`)
    - Must be one of the allowed [Device model's](#device-model) `platform` 
	enum values
- `name` (`string`)

#### Response
Body:

- `device` ([`device`](#device-model))
- `api_token` (`string`)

### Get Devices
#### Request
GET `/api/v0/devices`

Retrieves all a user's devices.

#### Response

- `devices` ([`device`](#device-model)[])

## Push Endpoints
### Create Push
#### Request
POST `/api/v0/push`

Body:

- `text` (`string`)

#### Response
Body:

- `push` ([`push`](#push-model))
