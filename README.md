# GRPC versus REST 

Wanted to do a blog about GRPC over REST performance wise, to end a discussion.

## Results:

```
go test ./... -bench Bench
testing: warning: no tests to run
PASS
BenchmarkGRPCSetInfo-8      5000            224633 ns/op
BenchmarkRESTSetInfo-8       200           5748596 ns/op
ok      grpc_v_rest     2.991s
```

[Read about it Here][blog]

[blog]: 

