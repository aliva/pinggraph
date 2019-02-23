# Ping Graph

Before starting make sure you can connect to all nodes with an ssh connection. `pinggraph` connects to each node with ssh and tries to ping other nodes.

Note that you can also add external nodes, (for example you want to ping `8.8.8.8` from all nodes.) You just need to set `isremote: true` for that node.

## Usage


```
# Using pre-built file
pinggraph -f hosts.yml

# From source
go run *.go -f hosts.yml
```

### hosts.yml keys

(check `hosts.yml.example` for sample)

| Key      | Default      | Description |
| ---      | ------------ | ----------- |
| host     | **Required** | Public IP or domain of your node |
| user     | `root`         | User used for ssh connection into node
| name     | if empty uses host value | Name given to this node |
| isremote | false        | pinggraph doesn't ssh connect into remote nodes, but other nodes will ping it, (when you want to check if all nodes can connect into an external server like `8.8.8.8`)
