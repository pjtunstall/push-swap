# push-swap

1. [SET-UP](#1-set-up)
2. [ANOMALIES](#1-audit)

## 1. Set-Up

To compile `checker.go`, run `go build checker.go`.

## 2. Anomalies

The project description clearly state that every instruction input to `checker` must be followed by a newline character, `\n`:

`Checker will then read instructions on the standard input, each instruction will be followed by \n.`

They then absentmindedly show an example that violates this rule but still results in an `OK`:

```$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```

If one followes the project specification to the letter, this will not result in `OK` because the final instruction is not followed by a `\n`. But if a newline character is appended, our `checker` will indeed output `OK`.
