## Todo

- [ ] huge-mega: `Command finished with error: fork/exec /Users/tilo/.nvm/versions/node/v16.10.0/bin/node: argument list too long`need to switch from eval to file execution if too large for eval


## Performance

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
