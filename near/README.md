# Lotus NEAR Connector

The NEAR protocol is a sharded, proof-of-stake, layer-one blockchain that is simple to use, secure and scalable. The Lotus NEAR Connector makes use of NEAR Lake Framework to parse events in Rust, then passes them to a Lotus connector written in Go using websockets.

## Events

As of now, the Lotus NEAR connector indexes:

* Blocks
* Transactions
* Receipts
* ExecutionOutcomes
