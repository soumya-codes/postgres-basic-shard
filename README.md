# postgres-static shard
* Implement simple static sharding using Postgres in the backend.

## Prerequisites
* Docker
* Docker Compose
* Golang 1.22 or later

## Infrastructure
* Machine: MacBook M2 Pro

## Setup/TearDown
* Clone the repository
* The repository contains a `Makefile` with the following targets of interest:
  * `make setup` - Setup the below environment:
    * 2 independent Postgres docker instances with pre-populated heartbeat data for multiple machines.
    * Builds and runs simple Web Application that runs on port **8080** and routes user-requested queries on one of the shards based on 
    machine_id value.
  * `make teardown` - Teardown the above environment

## Shard(Postgres Instances) Coordinates
* shard0:
  * Host: localhost
  * Port: 5432
  * User: test_user
  * Password: test_password
  * Database: shard0

* shard1:
  * Host: localhost
  * Port: 5433
  * User: test_user
  * Password: test_password
  * Database: shard1

## DB Schema
* Schema:
  * Table: `heartbeat`
    * Columns:
      * `machine_id` - TEXT
      * `last_heartbeat` - BIGINT

## DB-Topology and Sharding Strategy
* The setup consists of 2 Postgres instances:
  * shard0 running on port 5432.
  * shard1 running on and 5433.

* The sharding strategy is simple and static:
  * The machines have the following naming convention: `machine_<id>`, for example machine_001.
  * All the even numbered machines are stored in shard0.
  * All the odd numbered machines are stored in the shard1.

## Prepopulated Test Data
* Shard 1:

|  machine_id   |  last_heartbeat  |
|:-------------:|:----------------:|
|  machine_001  |    1622548800    |
|  machine_003  |    1622721600    |
|  machine_005  |    1622894400    |
|  machine_007  |    1623067200    |
|  machine_009  |    1623240000    |

* Shard 2:

|  machine_id  |  last_heartbeat  |
|:------------:|:----------------:|
| machine_002  |    1622635200    |
| machine_004  |    1622808000    |
| machine_006  |    1622980800    |
| machine_008  |    1623153600    |
| machine_0010 |    1623326400    |

## Supported Operations:
* The Web Application supports the following operations:
  * `get` - Get the last heartbeat of a machine.
    * Example Invocation: `curl -X GET http://localhost:8080/heartbeat/machine_001` 
  * `set` - Set the last heartbeat of a machine.
    * Example Invocation: `curl -X PUT http://localhost:8080/heartbeat/machine_001 -H "Content-Type: application/json" -d '{"last_heartbeat": 1625160900}'`

## Example Test Cases
* Requests for even machines, that should land on shard0:
  * Get the last heartbeat of machine_002:
    * `curl -X GET http://localhost:8080/heartbeat/machine_002`
  * Set the last heartbeat of machine_002:
    * `curl -X PUT http://localhost:8080/heartbeat/machine_002 -H "Content-Type: application/json" -d '{"last_heartbeat": 1625160900}'`
* Requests for odd machines, that land on shard1:
  * Get the last heartbeat of machine_001:
    * `curl -X GET http://localhost:8080/heartbeat/machine_001`
  * Set the last heartbeat of machine_001:
    * `curl -X PUT http://localhost:8080/heartbeat/machine_001 -H "Content-Type: application/json" -d '{"last_heartbeat": 1625150700}'`
* **Note:** Verify the requests were successful and landed on the correct shards by checking the DB entries/updates in the corresponding shards.