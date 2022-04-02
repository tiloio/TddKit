# TDDKit

- test adapters for each test case
    - e.g. node for frontend tests, testcafe & java for backend tests
    - adapters specify how to run the tests
- manage teset Resources & Dependencies
    - on multiple levels
        - adapters
            - can wait for resources and dependencies
            - can take into consideration the dependencies result in the logs
        - tests
            - adapters can support this
            - can wait for resources and dependencies
            - can take into consideration the dependencies result in the logs

## Configuration

```yml
tests:
    - name: "frontend"
      type: "node"
      testMatch: [ "**/__tests__/**/*.[jt]s?(x)", "**/?(*.)+(spec|test).[jt]s?(x)" ] 
      resources: ["database", "s3"] # optional - can also be specified in the tests if adapter supports it  
      dependencies: [] # optional - dependencies to other tests
    - name: "e2e"
      type: "testcafe"
      testMatch: [ "**/e2e/**/?(*.)+(spec|test).[jt]s?(x)" ]   
      resources: ["database", "s3"] # optional - can also be specified in the tests if adapter supports it  
      dependencies: ["frontend"] # optional - dependencies to other tests
resources: # use executors format to pass it to executor
    - name: database
      cmd: docker run mysql...
    - name: s3
      cmd: docker run minio...
```

## Performance

Measrure performance with https://github.com/sharkdp/hyperfine.