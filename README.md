# Ping Graph

Before starting make sure you can connect to all nodes with an ssh connection. `pinggraph` connects to each node with ssh and tries to ping other nodes.

Note that you can also add external nodes, (for example you want to ping `8.8.8.8` from all nodes.) You just need to set `isremote: true` for that node.

## Usage


```
# Using pre-built file
pinggraph -f nodes.yml

# From source
go run *.go -f nodes.yml
```

### nodes.yml keys

(check `nodes.yml.example` for sample)

| Key      | Default      | Description |
| ---      | ------------ | ----------- |
| host     | **Required** | Public IP or domain of your node |
| user     | `root`         | User used for ssh connection into node
| name     | if empty uses host value | Name given to this node |
| isremote | false        | pinggraph doesn't ssh into remote nodes, but other nodes will ping it, (when you want to check if all nodes can connect into an external server like `8.8.8.8`)

## Screenshot

![screenshot](/screenshot.png)
