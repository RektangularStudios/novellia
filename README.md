# Novellia

The Novellia microservice exposes REST APIs for interacting with the Novellia Platform. This is the backend for the [Novellia SDK](https://github.com/RektangularStudios/novellia-sdk).

Initially, we neglect to create a CLI tool for interaction on the microservice's host. You can import the [Novellia API](https://github.com/RektangularStudios/novellia-api) OpenAPI specification into [Insomnia](https://insomnia.rest/) for manual testing.

## Running the server

### Pre-requisites

1. a locally running [cardano-node](https://github.com/input-output-hk/cardano-node) which should not be a block producer, or even better, should not be part of a stake pool.
2. a locally running [cardano-db-sync](https://github.com/input-output-hk/cardano-db-sync).
3. a locally running [cardano-graphql](https://github.com/input-output-hk/cardano-graphql).

This stack is required for Novellia to connect with and retrieve data from the Cardano blockchain.

### Execution

Install Go: https://golang.org/doc/install

Start the server without building
- `go run ./server/main.go`

Or compile a binary for deployment
- `go build -o novellia-server ./server/main.go`

Then just execute binary to start the server
- `./novellia-server`

## What features are supported?

For a comprehensive list, refer to the [Novellia SDK Documentation on our Wiki](https://rektangularstudios.com/wiki) (TODO)

Basically, the Novellia microservice is just a conventional server that abstracts away interaction with Cardano. It submits transactions and calls smart contracts so that you don't have to. If you know why light wallets exist (e.g. Yoroi, AdaLite), this exists for the same reason.

Mainstream game developers just don't care about blockchain. The APIs exposed by Novellia aren't meant to mirror the technicalities of Cardano, they're meant to abstract it away.

We want developers to call a function like `PostLimitOrderForNFT()` not `CheckIfLiquidityPoolExistsAndWhatAboutTheMarketMakerOhNoAndThenPostLimitOrder()`.

Otherwise mainstream game developers won't adopt the technology.

## I don't get it. I thought Novellia being decentralized meant everything ran on Cardano?

While any trustless components must be implemented on smart contracts, we still require a conventional backend.

Anyone building on Novellia will want to host their own microservice to ensure constant uptime for their business needs. There's no limitations on this since the code is open-source.

A centralization problem isn't expected to occur. In fact, it makes sense for a business to firewall their own microservice so they're only paying for the server loads relevant to their products.

### On-chain limitations

**We just can't run everything on Cardano, it'd be too expensive even with Hydra.**

The **Novellia Dashboard** needs to be able to read the Cardano blockchain to see transactions relevant to the Novellia Platform. This kind of intensive processing is not suitable for a smart contract. Moreover, it doesn't require trust: the trusty aspect is that the transactions are already correct on-chain.

Queries need to be executed against a backend. Only mutation operations need to happen on-chain.
- Querying open market orders
- Querying games and NFTs listed by the Novellia DAO

We also expect a need for handling the opening and closing of Hydra heads once this functionality is available on Cardano. It makes sense for a communication layer to exist between Novellia and stake pool operators.

### Separation of concerns

**We don't want stake pool operators to run Novellia. We don't want that tight coupling.**

Why?
- SPOs aren't incentivized to run Novellia. This makes them unreliable.
- Running Novellia uses system resources. This taxes systems already running `cardano-node`.
- The sheer number of API calls made to a Novellia instance would easily take down a stake pool. This kind of thing, especially for gaming at scale, requires sophisticated load balancing.

This doesn't mean an SPO running a Novellia on their stake pool infrastructure is necessarily bad, it just:
- isn't our goal to have SPOs run it.
- is a potential liability for the Cardano network.

So who should run their own Novellia instance?
- Anyone with software that needs constant access to the Novellia Platform and may lose a lot of money if other, trusted microservice instances go down.

For this reason, it is expected that Novellia is easy to deploy without a local instance of `cardano-node`.

## How does it work?

1. Game developer uses Novellia SDK to issue an API call to Novellia.
2. Novellia receives the call and does some processing.
3. Novellia maybe issues some commands to `cardano-node` instances running on other services. This may mean submitting a signed transaction or calling a smart contract.
4. Novellia returns a response to the game developer's application.

As far as the game developer is concerned, Cardano barely exists. At most, they need to surface wallets through their product and indicate transaction fees as a kind of pseudo-tax.

# GraphQL

Novellia connects to Cardano through GraphQL as exposed by [cardano-graphql](https://github.com/input-output-hk/cardano-graphql).

[The specification is for accessible data is here](https://input-output-hk.github.io/cardano-graphql/).
[The code for the specification is here](https://github.com/input-output-hk/cardano-graphql/blob/master/packages/api-cardano-db-hasura/schema.graphql)

We are using [this library](https://github.com/shurcooL/graphql) for querying with reflection.
