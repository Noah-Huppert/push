# Design
Server design.

# Table Of Contents
- [Models](#models)
	- [User Model](#user-model)
    - [Device Model](#device-model)
	- [API Token Model](#api-token-model)
	- [Push Model](#push-model)
- [gRPC Interface](#grpc-interface)
	- [Users Service](#users-service)
	- [Devices Service](#devices-service)
	- [Authentication Service](#authentication-service)
	- [Push Service](#push-service)
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
- `data` (`text`)

# gRPC Interface
gRPC services.  

Methods require call credentials with an API token provided in the 
`authorization` field unless noted otherwise.

## Users Service
### GetSelf
Returns information about authenticated user.

## Devices Service
### CreateDevice
Create a new device for user.

### GetDevices
Get all of user's devices.

### SetDevice
Edit properties of a device.

### DeleteDevice
Delete device.

## Authentication Service
### GetTokens
Get details about all of user's API tokens.

### CreateToken
Exchange user's password for API token.  

Call credentials are not required for this method.

### DeleteToken
Delete an API token.

## Push Service
### GetPushes
Get user's pushes. Paginated.

### CreatePush
Create push.
