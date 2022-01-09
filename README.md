## Current features

- Choose directory via `-path=your/directory` where to run all `.test.js` files or use `-glob=*/**.spec.js` to define your own [glob pattern](https://en.wikipedia.org/wiki/Glob_(programming))
- Prints successful and failed tests with erros
- only jest `test()`function is supported
- can run CommonJS and ESM (but no auto detection yet)
- runs everything in parallel
- test file size is limited (don't know the limit now, but `9.5mb` was too much, see todos)
- currently: performance measuring against Jest, its more than 100% faster in simple test use cases (need to explore more complex scenarios)

## Future features

- use esbuild to be able to run typescript
- in memory sourcemap with correct paths
- good logging
- fake `console.` and `process.stdout` functions to collect the logs and return it to the Go process via the result JSON
- implement all Jest test: `describe`, `describe.each`, `test.each`, `test.todo`, `xtest`, `xdescribe`, etc. (see all https://jestjs.io/docs/api)
- implement Jest setup methods like `beforeEach`, `afterAll`, etc. (see all https://jestjs.io/docs/api)
- include Jest `expect` for easy expectations and thinkg of a own assertion library or search on which solves our need (also for the future WebServer)
- define resources via JavaScript function à la `TestResources([resource1, resourc2]);` for each test file. Therefore we do two runs of every test file: 
    1. Run will inject fake `test();` etc functions which do execute nothing and the `TestResource();` function which will print every resource to stdout, so that the Go process can collect and create all needed resources. (Therefore the creationFunction of the resources have to be the same.)
    2. Run will inject real `test();` etc functions and a fake `TestResource();` function. Then it will execute all test.
- sort test files according to needed resources and run the files with the least needed resources first
- create resources in the background while the other tests are running
- collect logs of resources and tests in one global `log.file`
- build dependency tree of resources; you can define a resource depends on another resource (e.g. the WebApp-Server depnends on the MySql-Database). And execute the test files which the least dependen resources
- define test dependencies like `TestResources();` you get a function `TestDependencies(['./anotherTest.js', ''../a/second/test.js])`
- do `export const aTestDependencies = TestDependencies(['./anotherTest.js']);` to import this dependencies into another file. This helps you to define a dependency tree and this helps the IDE to support you by importing other dependencies.
- build dependency graph of all test files and include resources to get a path how to run your test most efficiently.
- build test history in defined directory. For each test run we will build a JSON file which includes the runtime and name of the test, the testfile, the dependencies of the file and the resources of the file. This file will be saved with the date into that directory.
- use the histories to detect flaky tests
- rerurn flaky tests (or other failing tests)
- create `--settings` (`settings.json`, `settings.js` or `settings.ts`) file which will define how the framework behaves
- autodetect settings and tests frome the directory where you start the process via cli
- build a WebServer `http://localhost:5000` which will have a nice UI to show:
    - what test fails why (IntelliJ like `assert.equal` comparison but better)
    - what resource logs what and connect it to failures (if there is an error or something)
    - what test did not run because a dependent test failed


## Todo

- [ ] Stacktrace should have lines of actual files, not of builded esbuild stuff. Using sourcemap does not work quite good, node prints `eval` as source.

## Performance

```
// Added Glob search to FRAMEWORK
// 100 runs each: 3 Tests, 1 Failing
Running JEST 100 times on /Users/tilo/workspace/test-framework/src/test_examples/simple
Cleared /private/var/folders/61/8w59p_ss67z9799jv3t03w400000gn/T/jest_dx

real    1m9.566s
user    0m54.219s
sys     0m11.987s

Running FRAMEWORK 100 times on /Users/tilo/workspace/test-framework/src/test_examples/simple

real    0m1.152s
user    0m0.355s
sys     0m0.278s
```


```
// 10 runs each: 1502 Tests, random (around 700) Failing
Running JEST 10 times on /Users/tilo/workspace/test-framework/src/test_examples/max
Cleared /private/var/folders/61/8w59p_ss67z9799jv3t03w400000gn/T/jest_dx

real    0m57.378s
user    1m38.066s
sys     0m9.433s

Running FRAMEWORK 10 times on /Users/tilo/workspace/test-framework/src/test_examples/max

real    0m1.137s
user    0m2.914s
sys     0m0.861s
```

```
// Added Goroutines (parallelization) to FRAMEWORK
// 100 runs each: 3 Tests, 1 Failing
Running JEST 100 times on /Users/tilo/workspace/test-framework/src/test_examples/simple
Cleared /private/var/folders/61/8w59p_ss67z9799jv3t03w400000gn/T/jest_dx

real    1m9.524s
user    0m54.273s
sys     0m11.757s

Running FRAMEWORK 100 times on /Users/tilo/workspace/test-framework/src/test_examples/simple

real    0m5.508s
user    0m7.411s
sys     0m1.806s
```

```
// JEST cache is now cleared once before all tests are executed
// 10 runs each: 3 Tests, 1 Failing
Running JEST 10 times on /Users/tilo/workspace/test-framework/src/test_examples/simple
Cleared /private/var/folders/61/8w59p_ss67z9799jv3t03w400000gn/T/jest_dx

real    0m7.247s
user    0m6.539s
sys     0m1.424s

Running FRAMEWORK 10 times on /Users/tilo/workspace/test-framework/src/test_examples/simple

real    0m1.119s
user    0m0.744s
sys     0m0.194s
```

```
// added output logging to FRAMEWORK
// 100 runs each: 3 Tests, 1 Failing
Running JEST 100 times on /Users/tilo/workspace/test-framework/src/test_examples/simple

real    1m10.230s
user    0m53.917s
sys     0m11.941s

Running FRAMEWORK 100 times on /Users/tilo/workspace/test-framework/src/test_examples/simple

real    0m9.883s
user    0m7.516s
sys     0m1.994s
```

```
// 100 runs each: 2 Tests, 0 Failing
Running JEST 100 times on /Users/tilo/workspace/test-framework/src/test_examples/simple

real    1m8.107s
user    0m51.744s
sys     0m11.800s

Running FRAMEWORK 100 times on /Users/tilo/workspace/test-framework/src/test_examples/simple

real    0m9.750s
user    0m7.420s
sys     0m1.976s
```

- Jest: 68.107s / 100 = 0.68107s per run 
- Framework: 9.750s / 100 = 0.0975s per run# test-framework
