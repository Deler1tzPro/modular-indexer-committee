# Committee Indexer [![Join Nubit Discord Community](https://img.shields.io/discord/916984413944967180?logo=discord&style=flat)](https://discord.gg/5sVBzYa4Sg) [![Follow Nubit On X](https://img.shields.io/twitter/follow/nubit_org)](https://twitter.com/Nubit_org)

<!-- <p align='left'>
  <img width='70%' src='./img/logo.png' />
</p> -->
![Logo](img/logo.png)
## Background
Modular Indexer introduces a fully user-verified execution layer for meta-protocols on Bitcoin. Leveraging the immutable and decentralized nature of Bitcoin, the Modular Indexer provides a Turing-complete execution layer capable of running complex logic that cannot be directly executed on Bitcoin due to its script language's limitations.

## What is Committee Indexer?
Committee indexer serves as a key component of Modular Indexer, and is responsible for reading each block of Bitcoin, calculating protocol states, and summarizing these states as a polynomial commitment namely checkpoint. Whenever the committee indexer obtains a new Bitcoin block, it generates a new checkpoint for the protocol and publishes it to the data availability layer for users to access. It is permissionless; anyone can operate its committee indexer for a given meta-protocol.

## Getting Started
Modular Indexer is built with Golang. You can run your own modular Indexer by following the procedure below. `Go` version `1.22.0` is required for running repository. Please visit [Golang download Page](https://go.dev/doc/install) to get latest Golang installed.

### 1. Install Dependence
Golang is easy to install all dependence. Fetch all required package by simply running.
```Bash
go mod tidy
```

### 2. Prepare config.json
```Bash
cp config.example.json config.json
# Next you should tailor config.json by yourself
``` 
See [Details](#preparing-configjson) of how to set up your own `config.json`.

### 3. Run with Command Flag
```Bash
go run . --service --committee --cache
```
Below are the explanation for each of the command flags.
- `--service` `(-s)`: Use this flag to activate web service from moduler indexer. When enabled, the moduler indexer will provide web service for incoming query.

- `--committee`: This flag activates the committee indexer functionality. When enabled, the moduler indexer will provide checkpoints to the DA layer.

- `--cache`: By default, the state root cache is enabled, facilitating efficient verkle tree storage. This flag ensures that the application starts with the cache service activated, and will therefore fasten the initialization speed next time.

## Preparing Config.json
To ensure the Modular Indexer runs smoothly in your environment, it's crucial to properly configure the config.json file.

### Setting Up `database` Configuration
The database section requires connection details to the OPI database. If you're running an OPI full node, ensure to provide the correct details as follows:
- "host": The IP address or hostname of the machine where database is running.
- "user": The username for accessing the database.
- "password": The password associated with the above user account.
- "dbname": The name of the database you're connecting to.
- "port": The port number on which your database service is listening.

### Setting Up `report` Configuration
The report section lets you define where and how to store the checkpoints generated by your Committee Indexer. Currently, support is included for AWS S3, with a Data Availability (DA) layer option coming soon. Use `S3` or `DA` as the value.

S3 Configuration:
- "bucket": The name of the S3 bucket where checkpoints will be stored.
- "accessKey": Your AWS access key ID for authentication.
- "secretKey": Your AWS secret key for authentication.

### Setting Up `service` Configuration
The service section defines the details about your API service that provides access to the Modular Indexer functionalities.

- "url": The IP address or hostname where your API service is accessible. You should ensure this is reachable by the clients intended to use the service.
- "name": A unique identifier for your Indexer instance. This can be any name you choose and will be used in the checkpoint filename to identify the source.
- "metaProtocol": The meta-protocol your Indexer is serving. For this configuration, it's predefined as "brc-20".

<!-- ## Service API -->

## Useful Links
:spider_web: <https://www.nubit.org>  
:octocat: <https://github.com/Wechaty/wechaty>  
:beetle: <https://github.com/RiemaLabs/indexer-committee/issues>  
:book: <https://docs.nubit.org/developer-guides/introduction>  


## FAQ
- **Is there a consensus mechanism among committee indexers?**
    - No, within the committee indexer, only one honest indexer needs to be available in the network to satisfy the 1-of-N trust assumption, allowing the light indexer to detect checkpoint inconsistencies and thus proceed with the verification process.
- **How is the set of committee indexers determined?**
    - Committee indexers must publish checkpoints to the DA Layer for access by other participants. Users can maintain their list of committee indexers. Since the user's light indexer can verify the correctness of checkpoints, attackers can be removed from the committee indexer set upon detection of malicious behavior; the judgment of malicious behavior is not based on a 51% vote but on a challenge-proof mechanism. Even if the vast majority of committee indexers are malicious, if there is one honest committee indexer, the correct checkpoint can be calculated/verified, allowing the service to continue.
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
