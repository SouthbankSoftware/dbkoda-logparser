# dbKoda Log Parser
*Proof of concept MongoDB log parser for all platforms*

## Usage
1. `go build`
2. `./dbkoda-logparser PATH_TO_MONGO_DB_LOGS`

## PATH_TO_MONGO_DB_LOGS
You can get the path using `db.serverCmdLineOpts().parsed.systemLog.path`
