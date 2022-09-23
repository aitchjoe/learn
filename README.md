# TIL

## Test

* `go test -v ./pkg/shift`

## Fuzz

* `go test -fuzz=Fuzz ./pkg/shift`

Bug found:
```
C:\joe\workspace\aitchjoe\learn>go test -fuzz=Fuzz ./pkg/shift
fuzz: elapsed: 0s, gathering baseline coverage: 0/2 completed
fuzz: elapsed: 0s, gathering baseline coverage: 2/2 completed, now fuzzing with 16 workers
fuzz: minimizing 49-byte failing input file
fuzz: elapsed: 0s, minimizing
--- FAIL: FuzzShift (0.08s)
    --- FAIL: FuzzShift (0.00s)
        shift_test.go:93: Shift(Shift("0", []), []) = "", want "0"

    Failing input written to testdata\fuzz\FuzzShift\8d9227f8196001c8766ef2332f636f0aeb3015755c7c13d9163781ca4f78b2b6
    To re-run:
    go test -run=FuzzShift/8d9227f8196001c8766ef2332f636f0aeb3015755c7c13d9163781ca4f78b2b6
FAIL
exit status 1
FAIL    github.com/aitchjoe/learn/pkg/shift     0.276s
```
