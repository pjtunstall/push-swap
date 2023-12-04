# push-swap

1. [SET-UP](#1-set-up)
2. [ANOMALIES](#1-audit)

## 1. Set-Up

To compile `checker.go`, run `go build checker.go`; and, to run it, `./checker`, piping in the instructions like so:

```
echo -e "rra\npb\nsa\nrra\npa\n" | ./checker "3 2 1 0"
```

## 2. Anomalies

Notice how, in the example above, every instruction is followed by a newline character, `\n`? This is required according to the project description:

```
Checker will then read instructions on the standard input, each instruction will be followed by \n.
```

However, they then absentmindedly show an example that violates this rule but still results in an `OK`:

```
$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```

So, if you find an example like this that isn't `OK` even though the instructions are right, just append a `\n`, and all should be well.
