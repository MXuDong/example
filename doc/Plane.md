# The plane for the example

## TODO list

- Support invoke inner project
  
    For test mesh feature, some application has server more than 100.
    So for mock invoke inner application, it must support some method to invoke 
    in the project.

- Support UDP and other protocol

    Some applications like game server, use `UDP` protocol to trans info between 
    the client and server without check. And for mesh framework like `Istio`
    is support the `UDP`, I think the mock example should support some protocols
    like UDP and so on. 
    
    Now the Protocol is support :
  
    - Tcp

## Already Support