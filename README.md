### Bin Store

A store for storing data, using a simple key value pair

#### COMMANDS

##### [0] WRITE

    cache the data in-memory

##### [1] READ

    read from the in-memory map
    *TODO the key right now is the unix time, but this will cause a race condition with more clients writing

##### [2] STORE

    store the in-memory map into a `.gob` file on disk under the `.store` directory

##### [3] LOAD

    load a `.gob` store into memory

##### [4] MERGE

    merges a store on disk with the data in-memory

##### [5] CLEAR

    clears the in-memory data
* Need to implement 

- [ ] develop the client so that its better...
- [ ] support multiple store?
  - send the store name in the binary
  - you could use an URN as the key i.e. USER:1234

![Workflow](./docs/binstore.jpg "Bin Store Workflow")
