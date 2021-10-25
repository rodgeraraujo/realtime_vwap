# Real time VWAP Calculator

a real-time VWAP (volume-weighted average price) calculation engine for different trading pairs.


## Design Overview

### `cmd/vwap/main.go`

The main entry point will run, and retieves the trading pairs from `coinbase` websocket and calculates the VWAP for each trading pair.


### `cmd/pkg/coinbase`

Has the entire bussines logic, for how to calculate the VWAP of the trading pairs.


## Start using 

See the documentation in the `/docs` repository, the index can be found [here](./docs/README.md)

