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

## v0.4.0
- Add updated `/orders` for interaction with payment processor

## v0.5.0
- Add support for non-token products to `/products`

## v0.6.0
- Minor breaking changes to SDK (mostly just field names)
  - added `rarity` to Occulta Novellia Character model

## v0.6.1
- Added connection to Postgres
  - Querying of live product data on `/products` endpoint

## v0.7.0
- Automatically sort returned products by `id`
- `/products` supports fetching list of product IDs and returning details for specific ones

## v0.8.0
- Support org and market descriptions
- Restructure SDK to match new API split introducing `order-fulfillment`
- Remove `/orders` endpoint (migrated to `order-fulfillment`)
- Improve `/status` endpoint

## v0.8.1
- Return time as `ISO-8601 format (2021-05-17T22:00:00-00:00)`

## v0.8.2
- Add health check metrics for Prometheus

## v0.8.3
- Use connection pool for concurrent queries to Postgres

## v0.8.4
- Fix attribution information missing (bug)

## v0.9.0
- `/products`
  - Request products by native token ID
  - `modified` field to for more easily overriding front-end cache
- `/wallet`
  - Support an array of payment and stake addresses
  - Return on-chain metadata
- monitoring
  - Namespace monitoring to isolate prod from demo
  - Add `Errors` list to Status endpoint result to help with diagnosing service specific problems
