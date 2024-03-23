# compliance-blockchain
The compliance blockchain is a way for new servers to be posted onto a blockchain for compliance testing

## Install ignite
Install ignite by running the following command:
```
curl https://get.ignite.com/cli! | bash
```
Ignite CLI is a command-line interface for creating, managing, and deploying blockchain applications.
It is built on top of Starport, a blockchain application scaffold for Cosmos SDK. 
More information can be found at [Ignite CLI](https://ignite.com/cli).

## Get started
Build the cli by running the following command:
```
ignite chain build
```

To serve the blockchain, run the following command:
```
ignite chain serve
```
The `serve` command installs dependencies, builds, initializes, and starts your blockchain in development.
After the blockchain is running, you can start creating, updating, and querying data on the blockchain.

To post an approval to the blockchain, run the following command:
```
complianced tx compliance create-approval "test-driver-id-1" "www.test-1.com"  --from alice --chain-id compliance
```

To update an approval on the blockchain, run the following command:
```
complianced tx compliance update-approval "0" "approver-driver-id-2" "www.test-2.com" true  --from alice --chain-id compliance
```

To query the blockchain for all approvals, run the following command:
```
complianced q compliance list-all-approvals
```

To query the blockchain for a pending approvals, run the following command:
```
complianced q compliance list-pending-approvals
```

To query the blockchain for a specific approval, run the following command:
```
complianced q compliance list-approved-approvals
```

To query the blockchain for a rejected approvals, run the following command:
```
complianced q compliance list-rejected-approvals
```

To get a specific approval, run the following command:
```
complianced q compliance get-approval-by-id "0"
```