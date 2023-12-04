# push-swap

1. [Getting Started](#1-getting-started)
2. [Anomalies](#1-audit)
3. [Bitmasks: a Detour](#-bitmasks:-a-detour)

## 1. Getting Started

To compile `checker.go`, run `go build checker.go`; and, to run it, `./checker`, piping in the instructions like so:

```
echo -e "rra\npb\nsa\nrra\npa\n" | ./checker "3 2 1 0"
```

or you can run `./checker "3 2 1 0"` (with your choice of initial values on stack `a`), then type instructions on the command line, pressing enter after each. When you've typed all your instructions, you can press enter once more to let the checker know you've finished.

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

## Bitmasks: a Detour

In the `readInstructions` function in `checker`, we wanted to move the cursor up a line to eliminate the blank line that results when the user indicates that they've finished typing instructions by pressing enter on a line with no instructions.

However, we can't unconditionally move up a line because when the intructions are piped to the program, there is no blank line. Hence we check whether the input is from the terminal before moving up a line.

This is done by checking if the input is from a character device. The method `fi.Mode()` returns a bitmask, i.e. a number that represents a sequence of bits. In the case of `fi.Mode()`, these bits represent the file mode and permissions. One of them indicates whether the file is a character device. This should be true when the input is from the command line.

`os.ModeCharDevice` is a constant that indicates which bit, this is. It's value is 8192. In binary, 8192 is 10000000000000, a 1 followed by 13 zeros. So, being a power of two, 8192 can stand for a single bit, the 14th bit of the sequence.

The `&` is a bitwise AND operator. It's used here to check if the bit that represents a character device is set (i.e. 1). The result of this AND operation will only be zero when the 14th bit of `fi.Mode` is zero. Otherwise it will be 8192.
