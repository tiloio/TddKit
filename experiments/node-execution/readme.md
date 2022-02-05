# Nod execution experiment


# 100 Processes with `console.log("Hello World")`

- Multi processes
    - 1,92s user 0,41s system 91% cpu 2,549 total
- One process
    - 0,04s user 0,01s system 20% cpu 0,259 total   

Difference 2,29 seconds = 2290 ms / 100 runs = 22,9ms overhead

Currently we need min. 2 process per file (`discovery`, `runTest`). For each resource we need another process.

Example project with 12 resources and 1000 tests = 2012 runs.

This would make 2012 run * 22,9 ms = 46.074,8 ms = **~46 seconds overhead**.
IF we parallelize with 4 runs this will cut into 503 runs * 22,9 ms = 11.518,7 ms = **~11,5 seconds overhead in parallelization**.


# 100 Processes with `console.log("Hello World")` in `go`

- Multi processes
    - real    0m2.830s
    - user    0m2.021s
    - sys     0m0.574s
- One process
    - real    0m0.002s
    - user    0m0.000s
    - sys     0m0.001s

Difference 2,83 seconds - 0.002 seconds = 2828 ms / 100 runs = 28,28ms overhead

# 100 Processes with `console.log("Hello World")` in `go` with one go routines for each

- Multi processes
    - real    0m0.972s
    - user    0m3.006s
    - sys     0m1.230s
- One process
    - real    0m0.002s
    - user    0m0.000s
    - sys     0m0.001s

Difference 0,972 seconds - 0,002 seconds = 970 ms / 100 runs = 9,7ms overhead

# 100 Processes with `console.log("Hello World")` in `go` with amount of go routines amount of threads
- Multi processes
    - real    0m0.972s
    - user    0m3.006s
    - sys     0m1.230s
- One process
    - real    0m0.002s
    - user    0m0.000s
    - sys     0m0.001s

Difference 0,972 seconds - 0,002 seconds = 970 ms / 100 runs = 9,7ms overhead

# 100 Processes with `console.log("Hello World")` in `go` with combined files of amount of threads in go routines
- Multi processes
    - real    0m0.452s
    - user    0m0.467s
    - sys     0m0.254s
- One process
    - real    0m0.002s
    - user    0m0.000s
    - sys     0m0.001s

Difference 0,452 seconds - 0,002 seconds = 450 ms / 100 runs = 4,5ms overhead


