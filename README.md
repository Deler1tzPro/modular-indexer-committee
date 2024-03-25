# Modular Indexer (Committee) [![Join Nubit Discord Community](https://img.shields.io/discord/916984413944967180?logo=discord&style=flat)](https://discord.gg/5sVBzYa4Sg) [![Follow Nubit On X](https://img.shields.io/twitter/follow/nubit_org)](https://twitter.com/Nubit_org)

<img src="assets/logo.svg" width="400px" alt="Nubit Logo" />

***Warning!*** *This release is specifically for the Pre-alpha Testnet and may include changes that are not backward compatible in the future.*

## Background
The Modular Indexer, which includes the Modular Indexer (Committee) and the [Modular Indexer (Light)](https://github.com/RiemaLabs/modular-indexer-light), introduces a fully user-verified execution layer for Bitcoin's meta-protocols. By leveraging the immutable and decentralized characteristics of Bitcoin, it provides a Turing-complete execution layer, going beyond the limitations of Bitcoin's script language.

Our innovative approach uses Verkle trees for trusted, decentralized data integrity. Even with a majority of hostile modular-indexer-committees, the Modular Indexer reliably connects Bitcoin with complex applications like BRC-20, propelling the ecosystem forward. 

For a detailed understanding, refer to our paper: ["Modular Indexer: Fully User-Verified Execution Layer for Meta-protocols on Bitcoin"](https://eprint.iacr.org/2024/408). Stay updated on the latest progress in our [L1F Discourse Group](https://l1f.discourse.group/t/modular-indexer-fully-user-verified-execution-layer-for-meta-protocols-on-bitcoin/598).


## What is modular-indexer-committee?
modular-indexer-committee serves as a key component of Modular Indexer, and is responsible for reading each block of Bitcoin, calculating protocol states, and summarizing these states as a polynomial commitment namely checkpoint. Whenever the modular-indexer-committee obtains a new Bitcoin block, it generates a new checkpoint for the protocol and publishes it to the data availability layer for users to access. It is permissionless; anyone can operate its modular-indexer-committee for a given meta-protocol.

## Getting Started
Welcome to the Modular Indexer setup. Ensure your machine is equipped for optimal performance with the recommended specifications:

| Metric     | Minimum Recommended | Notes                          |
|------------|---------------------|--------------------------------|
| **CPU**    | [CPU Specification] | For efficient processing       |
| **Memory** | [Memory Specification] | To handle operations effectively |
| **Disk**   | [Disk Specification] | Adequate storage for speed     |
| **Bandwidth** | [Bandwidth Specification] | Stable network connection    |

Modular Indexer is built with Golang. You can run your own modular Indexer by following the procedure below. `Go` version `1.22.0` is required for running repository. Please visit [Golang download Page](https://go.dev/doc/install) to get latest Golang installed.

### 1. Install Dependence
Golang is easy to install all dependence. Fetch all required package by simply running.
```Bash
go mod tidy
```

### 2. Prepare config.json
```Bash
cp config.example.json config.json
# Tailor config.json according to your setup
``` 
See [Details](#preparing-configjson) of how to set up your own `config.json`.

### 3. Run with Command Flag

```Bash
# Build the modular-indexer-committee
go build

# Run the modular-indexer-committee
./indexer-committee --committee --service

# Run the modular-indexer-committee in test modee
./indexer-committee --committee --service -t --blockheight 780010
```
Below are the explanation for each of the command flags.
- `--service` `(-s)`: Use this flag to activate web service from moduler indexer. When enabled, the moduler indexer will provide web service for incoming query.

- `--committee`: This flag activates the modular-indexer-committee functionality. When enabled, the moduler indexer will provide checkpoints to the DA layer.

- `--cache`: By default, the state root cache is enabled, facilitating efficient verkle tree storage. This flag ensures that the application starts with the cache service activated, and will therefore fasten the initialization speed next time.

- `-t` `(--test)`: Enable this flag to activate test mode, allowing the indexer to operate up to a specified block height limit. This mode is useful for development and testing by simulating the indexer's behavior without catching up to the real latest block.

- `--blockheight`: When test mode is enabled with -t, this flag sets a fixed maximum block height limit for the indexer's operations. It allows for focused testing and performance tuning by limiting the range of blocks the indexer processes.

## Preparing Config.json
Proper configuration of config.json is key for the smooth operation of the Modular Indexer.

### Setting Up `database` Configuration
The database section requires connection details to the OPI database. If you're running an OPI full node, ensure to provide the correct details as follows:
- `host`: The IP address or hostname of the machine where database is running.
- `user`: The username for accessing the database.
- `password`: The password associated with the above user account.
- `dbname`: The name of the database you're connecting to.
- `port`: The port number on which your database service is listening.

### Setting Up `report` Configuration
Define where and how to store the checkpoints generated by your modular-indexer-committee. The report section currently supports AWS S3 and the Data Availability (DA) layer.

- `method`: Choose between `DA` and `S3` for storage method.
- `timeout`: Timeout setting in milliseconds for publishing checkpoints.

**S3 Configuration:**
- `region`: Specify the AWS S3 region for storage.
- `bucket`: Name of the S3 bucket where checkpoints are stored.
- `accessKey`: Your AWS access key ID.
- `secretKey`: Your AWS secret access key.

**DA Configuration:**
- `network`: Specify the network (current: 'Pre-Alpha Testnet').
- `namespaceID`: Your designated namespace identifier.
- `gasCode`: Custom code for managing transaction fees.
- `privateKey`: Your private key for secure transactions.

### Setting Up `service` Configuration
The service section specifies the details of your API service, enabling access to the Modular Indexer functionalities.

- `url`: The URL where your API service is hosted and accessible.
- `name`: A unique name for your Indexer instance.
- `metaProtocol`: Specify the meta-protocol served by your Indexer (default 'brc-20').
- `version`: Specify the version of your modular-indexer-committee (current: 'v0.1.0-rc.0').

<!-- ## Service API -->

## Useful Links
:spider_web: <https://www.nubit.org>  
:octocat: <https://github.com/Wechaty/wechaty>  
:beetle: <https://github.com/RiemaLabs/indexer-committee/issues>  
:book: <https://docs.nubit.org/developer-guides/introduction>  


## FAQ
- **Is there a consensus mechanism among modular-indexer-committees?**
    - No, within the modular-indexer-committee, only one honest indexer needs to be available in the network to satisfy the 1-of-N trust assumption, allowing the light indexer to detect checkpoint inconsistencies and thus proceed with the verification process.
- **How is the set of modular-indexer-committees determined?**
    - modular-indexer-committees must publish checkpoints to the DA Layer for access by other participants. Users can maintain their list of modular-indexer-committees. Since the user's light indexer can verify the correctness of checkpoints, attackers can be removed from the modular-indexer-committee set upon detection of malicious behavior; the judgment of malicious behavior is not based on a 51% vote but on a challenge-proof mechanism. Even if the vast majority of modular-indexer-committees are malicious, if there is one honest modular-indexer-committee, the correct checkpoint can be calculated/verified, allowing the service to continue.
- **Why do users need to verify data through checkpoints instead of looking at the simple majority of the indexer network?**
    - This would lead to Sybil attacks: joining the indexer network is permissionless, without a staking model or proof of work, so the economic cost of setting up an indexer attacker cluster is very low, requiring only the cost of server resources. This allows attackers to achieve a simple majority at a low economic cost; even by introducing historical reputation proof, without a slashing mechanism, attackers can still achieve a 51% attack at a very low cost.
- **Why are there no attacks like double-spending in the Modular Indexer architecture?**
    - Bitcoin itself provides transaction ordering and finality for meta-protocols (such as BRC-20). It is only necessary to ensure the correctness of the indexer's state transition rules and execution to avoid double-spending attacks (there might be block reorganizations, but indexers can correctly handle them).
- **Why upload checkpoints to the DA Layer instead of a centralized server or Bitcoin?**
    - For a centralized server, if checkpoints are stored on a centralized network, the service loses availability in the event of downtime, and there is also the situation where the centralized server withholds checkpoints submitted by honest indexers, invalidating the 1-of-N trust assumption.
    - For indexers, checkpoints are frequently updated, time-sensitive data:
        - The state of the Indexer updates with block height and block hash, leading to frequent updates of checkpoints (~10 minutes).
        - The cost of publishing data on Bitcoin in terms of transaction fees is very high.
        - The data throughput demand for hundreds or even thousands of meta-protocol indexers storing checkpoints is huge, and the throughput of Bitcoin cannot support it.
- **What are the mainstream meta-protocols on Bitcoin currently?**
    - The mainstream meta-protocols are all based on the Ordinals protocol, which allows users to store raw data on Bitcoin. BRC-20, Bitmap, SatsNames, etc., are mainstream meta-protocols. More meta-protocols and information can be found [here](https://l1f.discourse.group/latest)
- **What kind of ecosystem support has this proposal received?**
    - The proposal is put forward by Nubit as a long-term supporter and builder of the Bitcoin ecosystem. We have also exchanged ideas with many ecosystem partners and hope to jointly promote the progress and improvement of the modular indexer architecture.

<!-- ## License -->
