### Bin Store

A store for storing data, using a simple key value pair

#### COMMANDS

##### WRITE

    cache the data in-memory

##### STORE

    store the in-memory map into a `.gob` file under the `.store` directory

##### LOAD

    load a `.gob` store into memory

##### READ

    read from the in-memory map
    *TODO the key right now is the unix time, but this will cause a race condition with more clients writing

- [ ] develop the client so that its better...
- [ ] support multiple store?
  - send the store name in the binary
  - you could use an URN as the key i.e. USER:1234

![Workflow](./docs/binstore.jpg "Bin Store Workflow")
