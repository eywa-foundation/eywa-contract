#!/bin/sh

# set variables for the chain
VALIDATOR_NAME=validator1
CHAIN_ID=eywa-contract
KEY_NAME=eywa-contract-key
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

# create a random Namespace ID for your rollup to post blocks to
NAMESPACE_ID=$(openssl rand -hex 10)
echo $NAMESPACE_ID

# query the DA Layer start height, in this case we are querying
# an RPC endpoint provided by Celestia Labs. The RPC endpoint is
# to allow users to interact with Celestia's core network by querying
# the node's state and broadcasting transactions on the Celestia
# network. This is for Arabica, if using another network, change the RPC.
DA_BLOCK_HEIGHT=$(curl rpc.celestia-mocha.com/block |jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

# build the eywa-contract chain with Rollkit
ignite chain build

# reset any existing genesis/chain data
eywa-contractd tendermint unsafe-reset-all

# initialize the validator with the chain ID you set
eywa-contractd init $VALIDATOR_NAME --chain-id $CHAIN_ID

# add keys for key 1 and key 2 to keyring-backend test
eywa-contractd keys add $KEY_NAME --keyring-backend test

# add these as genesis accounts
eywa-contractd add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test

# set the staking amounts in the genesis transaction
eywa-contractd gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test

# collect genesis transactions
eywa-contractd collect-gentxs

# export the Celestia light node's auth token to allow you to submit
# PayForBlobs to Celestia's data availability network
# this is for Arabica, if using another network, change the network name
export AUTH_TOKEN=$(celestia light auth admin --p2p.network mocha)

# start the chain
eywa-contractd start --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26658","timeout":60000000000,"fee":600000,"gas_limit":6000000,"auth_token":"'$AUTH_TOKEN'"}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT