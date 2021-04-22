# remember to change the version in 'server/version.go'

## v0.1.0
Initial deployment

## v0.2.0
- implement minimal service layer
  - `/cardano/status` returns GraphQL init status and sync percentage
  - `/cardano/tip` returns latest block and epoch
  - `/wallet/<address>` returns tokens held at address using the latest block as a reference point

- Add `/cardano/status` endpoint to return `cardano-graphql` status information

## v0.3.0
- Add `/products`
  - mocked with example response
- Add `/orders`
  - stubbed
- Rename `/cardano/status` to `/status`
  - add maintenance switch (stubbed)
  - generalize for multiple service statuses
- Add `mocked_api` for debugging without service dependencies
