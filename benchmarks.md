# benchmarks

Hello these are my benchmarks as I refactor!

## Min/Max

```shell
λ ~/code/playground/kv/ go test bench_min_max_test.go -bench=. -benchmem
goos: linux
goarch: amd64
cpu: 13th Gen Intel(R) Core(TM) i7-13700F
BenchmarkKVIntValidation/MaxInt/Valid-24                1000000000               0.5949 ns/op          0 B/op          0 allocs/op
BenchmarkKVIntValidation/MaxInt/Invalid-24              1000000000               0.8908 ns/op          0 B/op          0 allocs/op
BenchmarkKVIntValidation/MinInt/Valid-24                1000000000               0.5921 ns/op          0 B/op          0 allocs/op
BenchmarkKVIntValidation/MinInt/Invalid-24              1000000000               0.9888 ns/op          0 B/op          0 allocs/op
BenchmarkKVUintValidation/MaxUint/Valid-24              1000000000               0.6228 ns/op          0 B/op          0 allocs/op
BenchmarkKVUintValidation/MaxUint/Invalid-24            1000000000               0.8916 ns/op          0 B/op          0 allocs/op
BenchmarkKVUintValidation/MinUint/Valid-24              1000000000               0.6086 ns/op          0 B/op          0 allocs/op
BenchmarkKVUintValidation/MinUint/Invalid-24            1000000000               0.9894 ns/op          0 B/op          0 allocs/op
BenchmarkKVFloat64Validation/MaxFloat/Valid-24          1000000000               0.8509 ns/op          0 B/op          0 allocs/op
BenchmarkKVFloat64Validation/MaxFloat/Invalid-24        1000000000               1.100 ns/op           0 B/op          0 allocs/op
BenchmarkKVFloat64Validation/MinFloat/Valid-24          1000000000               0.8160 ns/op          0 B/op          0 allocs/op
BenchmarkKVFloat64Validation/MinFloat/Invalid-24        1000000000               1.106 ns/op           0 B/op          0 allocs/op
BenchmarkKVTimeValidation/MaxTime/Valid-24              464188100                2.589 ns/op           0 B/op          0 allocs/op
BenchmarkKVTimeValidation/MaxTime/Invalid-24            302106501                3.970 ns/op           0 B/op          0 allocs/op
BenchmarkKVTimeValidation/MinTime/Valid-24              466188385                2.568 ns/op           0 B/op          0 allocs/op
BenchmarkKVTimeValidation/MinTime/Invalid-24            303379029                3.944 ns/op           0 B/op          0 allocs/op
BenchmarkOzzoIntValidation/MaxInt/Valid-24              92836875                12.68 ns/op            0 B/op          0 allocs/op
BenchmarkOzzoIntValidation/MaxInt/Invalid-24             4246437               398.3 ns/op           384 B/op          3 allocs/op
BenchmarkOzzoIntValidation/MinInt/Valid-24              95015923                12.49 ns/op            0 B/op          0 allocs/op
BenchmarkOzzoIntValidation/MinInt/Invalid-24             4734108               367.7 ns/op           384 B/op          3 allocs/op
BenchmarkOzzoUintValidation/MaxUint/Valid-24            95561800                12.57 ns/op            0 B/op          0 allocs/op
BenchmarkOzzoUintValidation/MaxUint/Invalid-24           3485271               406.0 ns/op           384 B/op          3 allocs/op
BenchmarkOzzoUintValidation/MinUint/Valid-24            94816707                12.56 ns/op            0 B/op          0 allocs/op
BenchmarkOzzoUintValidation/MinUint/Invalid-24           7025714               378.8 ns/op           384 B/op          3 allocs/op
BenchmarkOzzoFloat64Validation/MaxFloat/Valid-24        55031554                20.27 ns/op            8 B/op          1 allocs/op
BenchmarkOzzoFloat64Validation/MaxFloat/Invalid-24               3248820               440.3 ns/op           392 B/op          4 allocs/op
BenchmarkOzzoFloat64Validation/MinFloat/Valid-24                53132193                20.99 ns/op            8 B/op          1 allocs/op
BenchmarkOzzoFloat64Validation/MinFloat/Invalid-24               2991337               448.2 ns/op           392 B/op          4 allocs/op
BenchmarkOzzoTimeValidation/MaxTime/Valid-24                     3105169               370.6 ns/op            24 B/op          1 allocs/op
BenchmarkOzzoTimeValidation/MaxTime/Invalid-24                   1350711               970.8 ns/op           408 B/op          4 allocs/op
BenchmarkOzzoTimeValidation/MinTime/Valid-24                     3208317               368.4 ns/op            24 B/op          1 allocs/op
BenchmarkOzzoTimeValidation/MinTime/Invalid-24                   1413831               916.3 ns/op           408 B/op          4 allocs/op
PASS
ok      command-line-arguments  45.852s
```

## Nil/Empty

```shell
λ ~/code/playground/kv/ go test bench_absent_test.go -bench=. -benchmem
goos: linux
goarch: amd64
cpu: 13th Gen Intel(R) Core(TM) i7-13700F
BenchmarkKVAbsentValidation/Nil/Valid-24                1000000000               0.2966 ns/op          0 B/op          0 allocs/op
BenchmarkKVAbsentValidation/Nil/Invalid-24              1000000000               0.7904 ns/op          0 B/op          0 allocs/op
BenchmarkKVAbsentValidation/Empty/Valid-24              1000000000               0.3970 ns/op          0 B/op          0 allocs/op
BenchmarkKVAbsentValidation/Empty/Invalid-24            1000000000               0.9146 ns/op          0 B/op          0 allocs/op
BenchmarkOzzoAbsentValidation/Nil/Valid-24              313988296                3.802 ns/op           0 B/op          0 allocs/op
BenchmarkOzzoAbsentValidation/Nil/Invalid-24            32311728                37.21 ns/op           16 B/op          1 allocs/op
BenchmarkOzzoAbsentValidation/Empty/Valid-24            30234775                42.67 ns/op           16 B/op          1 allocs/op
BenchmarkOzzoAbsentValidation/Empty/Invalid-24          28358311                39.10 ns/op           16 B/op          1 allocs/op
PASS
ok      command-line-arguments  8.972s
```
