# Demo Parallel Test Run

To see the output:

    $ go test ./... -v
    ?       github.com/baijum/demo-parallel-test-run        [no test files]
    === RUN   TestHello
    2017-10-23 14:27:38.218314966 +0530 IST
    --- PASS: TestHello (0.00s)
    PASS
    ok      github.com/baijum/demo-parallel-test-run/a1     0.010s
    === RUN   TestHello
    2017-10-23 14:27:38.238635035 +0530 IST
    --- PASS: TestHello (0.00s)
    PASS
    ok      github.com/baijum/demo-parallel-test-run/a2     0.002s
    === RUN   TestHello
    2017-10-23 14:27:38.226489923 +0530 IST
    --- PASS: TestHello (0.00s)
    PASS
    ok      github.com/baijum/demo-parallel-test-run/a3     0.011s
    === RUN   TestHello
    2017-10-23 14:27:38.280195323 +0530 IST
    --- PASS: TestHello (0.00s)
    PASS
    ok      github.com/baijum/demo-parallel-test-run/a4     0.002s


You can also specify the packages:

    $ go test ./a1 ./a2 ./a3 ./a4 -v
    === RUN   TestHello
    2017-10-23 14:29:13.244619412 +0530 IST
    --- PASS: TestHello (0.00s)
    PASS
    ok      github.com/baijum/demo-parallel-test-run/a1     0.002s
    === RUN   TestHello
    2017-10-23 14:29:13.240892164 +0530 IST
    --- PASS: TestHello (0.00s)
    PASS
    ok      github.com/baijum/demo-parallel-test-run/a2     0.004s
    === RUN   TestHello
    2017-10-23 14:29:13.235773538 +0530 IST
    --- PASS: TestHello (0.00s)
    PASS
    ok      github.com/baijum/demo-parallel-test-run/a3     0.002s
    === RUN   TestHello
    2017-10-23 14:29:13.219567053 +0530 IST
    --- PASS: TestHello (0.00s)
    PASS
    ok      github.com/baijum/demo-parallel-test-run/a4     0.006s
        
