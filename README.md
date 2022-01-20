## Current features

- Choose directory via `-path=your/directory` where to run all `.test.js` files or use `-glob=*/**.spec.js` to define your own [glob pattern](https://en.wikipedia.org/wiki/Glob_(programming))
- Prints successful and failed tests with erros
- only jest `test()`function is supported
- can run CommonJS and ESM (but no auto detection yet)
- runs everything in parallel
- currently: performance measuring against Jest, its more than 100% faster in simple test use cases (need to explore more complex scenarios)

## Future features

Test phases (defined through environment variable `TEST_FRAMEWORK_PHASE` in the node process):
1. Collect test files.
2. Build every file with `esbuild` to support typescript, esm and cjs. (Maybe only do this step if the file is a typescript file to save performance).
3. Run test files with `TEST_FRAMEWORK_PHASE=DISCOVERY` to
    - collect all resources
    - and collect all dependencies.
4. (Not execute for single file tests.) Build graph of dependencies and resources and in parallel<br>
   Create the needed resources    
5. Run test files with `TEST_FRAMEWORK_PHASE=TEST_RUN` and `TEST_FRAMEWORK_ID=SOME_ID` to
    - optinal wait until all exclusive resources are free (then lock them for this run),
    - initialize the needed resources with the `TEST_FRAMEWORK_ID` before each test executes and
    - wait for tests finishes.<br>
   In parallel:<br>
   Log everything from each resource and test to a specific log file for this run (e.g. `run-log-xxx.json`).
6. Display results (format is not clear but should be configurabel).

Each phase should become their own module in the code to be able to swap and change it independently.

### Test execution

- [x] use esbuild to be able to run typescript
- [ ] in memory sourcemap with correct paths
- [ ] good logging
- [ ] fake `console.` and `process.stdout` functions to collect the logs and return it to the Go process via the result JSON

Two different run scenarios:
1. TDD; you want ultra fast feedback and only execute one file or one test.
2. Overall Tests; you want to run all tests to check if everything is still working and get excellent feedback what is why wrong.

#### Jest legacy support

To support projects which use the jest syntax we will implement a flag `-jest` which will inject stuff into the test. The none jest way is to always import what you need to have (like `import { test } from 'test-framework'`).

- implement all Jest test: `describe`, `describe.each`, `test.each`, `test.todo`, `xtest`, `xdescribe`, etc. (see all https://jestjs.io/docs/api)
- implement Jest setup methods like `beforeEach`, `afterAll`, etc. (see all https://jestjs.io/docs/api)
- include Jest `expect` for easy expectations and thinkg of a own assertion library or search on which solves our need (also for the future WebServer)

### Test resources

- define resources via JavaScript function à la `TestResources([resource1, resourc2]);` for each test file. Therefore we do two runs of every test file: 
    1. We will execute the tests with the environment variable `TEST_FRAMEWORK_PHASE=DISCOVERY`.
    2. All `test();` etc functions will execute nothing.
    3. The `TestResource();` function will print every resource to stdout, so that the Go process can collect and create all needed resources. (Therefore the creationFunction of the resources has to be the same.)
    4. Test run will set the environment variable `TEST_FRAMEWORK_PHASE=TEST_RUN` (or the variable is not set which should do the same). 
    5. All `test();` etc functions will work like they should, only the `TestResource();` function will do nothing.
- sort test files according to needed resources and run the files with the least needed resources first.
- create resources in the background while the other tests are running.
- collect logs of resources and tests in one global `log.json` file.
- build dependency tree of resources; you can define a resource depends on another resource (e.g. the WebApp-Server depnends on the MySql-Database). And execute the test files which the least dependen resources.
- Some Resource are able to handle multiple instances at once (like a MySQL server), but other resources need to be created to have separate instances (like a web server). We call theses resources **exclusive resources** Therefore we need something like a waiting pool which limits the creation of these resources and locks the created resources for each run where they are needed. 

### Test dependencies

- define test dependencies like `TestResources();` you get a function `TestDependencies(['./anotherTest.js', ''../a/second/test.js])`
- do `export const aTestDependencies = TestDependencies(['./anotherTest.js']);` to import this dependencies into another file. This helps you to define a dependency tree and this helps the IDE to support you by importing other dependencies.
- build dependency graph of all test files and include resources to get a path how to run your test most efficiently.
- build test history in defined directory. For each test run we will build a JSON file which includes the runtime and name of the test, the testfile, the dependencies of the file and the resources of the file. This file will be saved with the date into that directory.
- use the histories to detect flaky tests
- rerurn flaky tests (or other failing tests)
- create `--settings` (`settings.json`, `settings.js` or `settings.ts`) file which will define how the framework behaves
- autodetect settings and tests frome the directory where you start the process via cli

### Webserver

Build a WebServer `http://localhost:5000` which will have a nice UI to show:
- what test fails why (IntelliJ like `assert.equal` comparison but better)
- what resource logs what and connect it to failures (if there is an error or something)
- what test did not run because a dependent test failed

Server has to run everytime in background
- Communication via files in the project (e.g. `.test-framework/run-xxx.json`)  
- Test process starts the server in a detached process
- Test process checks via http call if test server is already running, if yes it also gets configured where the files live in that http call


## Todo

- [ ] Stacktrace should have lines of actual files, not of builded esbuild stuff. Using sourcemap does not work quite good, node prints `eval` as source. Maybe solveable if we search all `[stdin]` logs and replace them with the last `// ./file.xy` comment of the mentioned line. Then substract the comment rowcount from that line...
- [ ] `Run /Users/tilo/workspace/test-framework/src/framework go run . -path "../test_examples/max" -glob "/**/*.test.[tj]s" -esm=true` fails with `Error: Dynamic require of "assert" is not supported` -> does not appear if we run it on the file system 🤷🏽‍♂️
- [ ] exclude efficiently node_modules folder while searching for tests - or allow folders to ignore.
- [ ] jest - support option (injects test, describe, expect, etc.) - our version should use a explicit import to avoid the injection at beginning of the tests. This will also allow to disable the tests in the discover phase with a environment variable like `TEST_FRAMEWORK_PHASE=DISCOVERY`.
- [ ] heavy init test example: A example which tests if its faster to run jest or this framework if there is a initialization thing which would only be once by jest.
- [x] multiple calls of `TestDependencies` and `TestRescources` in one test file. `-->` use last log to be able to use `TestDependencies` and depend on it.
- [x] when we use the Object from `TestDependencies` to describe dependencies and import it into another test file we end up in stacking these files together. So if test `a` uses test `b` and we build and execute test `a` test `b` would be execute before test `a` is executed. This is not a problem, this would even help us. But if we have also a test `c` which depends on test `b`, then both test `a` and `c` would execute test `b`. This will lead to unessecary executions. `-->` We set a env variable with the dependency name and make the test only active after the dependency creation of the given ID has run.
- [ ] Error if there are two dependencies with the same name.
- [ ] Try to use only one esbuild call with all found test files as multiple entry points. This maybe speeds everything up and can be handle on file level. Maybe we also can use a virtual file system like FS.embed does to not rely on disk writes and reads which will definitly slow the process.
- [ ] Count tests in discovery
- [ ] Visualize dependencies
- [ ] discover cross dependencies 

## Performance

```
Running FRAMEWORK with stdin 100 times on /Users/tilo/workspace/test-framework/src/test_examples/max

user    0m53.675s
sys     0m12.572s

Running FRAMEWORK with file creation 100 times on /Users/tilo/workspace/test-framework/src/test_examples/max

real    0m31.063s
user    1m17.018s
sys     0m13.433s
```

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
