# push-swap

1. [Getting started](#1-getting-started)
2. [Research](#-research)
3. [Anomalies](#1-anomalies)
4. [Other people](#-other-people)
5. [Bitmasks: a detour](#-bitmasks:-a-detour)

## 1. Getting Started

You'll find the main package for each of the two programs, `checker` and `push-swap`, in the folders of those names in `cmd`.

To compile the `push-swap` program, cd into the corresponding directory and run `go build -o push-swap main.go`. To use the program, enter `./push-swap`, followed by a string of integers to sort, separated by spaces.

To compile the `checker` program, run `go build -o checker main.go`; and, to use it, type `./checker`, followed by a string of integers to sort, piping in the instructions like so:

```
echo -e "rra\npb\nsa\nrra\npa\n" | ./checker "3 2 1 0"
```

Alternatively, you can run `./checker "3 2 1 0"` (with your choice of initial values on stack `a`), then type instructions on the command line, pressing enter after each. When you've typed all your instructions, you can hit enter one last time to let the checker know you've finished.

In case you want to run `main_test.go`, be aware that it expects both these binaries to be built and in their respective folders.

## 2. Research

## 3. Anomalies

In the `checker` example above, every instruction is followed by a newline character, `\n`, as it should be according to the project description:

```
Checker will then read instructions on the standard input, each instruction will be followed by \n.
```

The project description then shows an example that violates this rule but still results in an `OK`:

```
$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```

Presumably the contradiction is due to a typo: either a missing `\n` from the end of the instructions or an `OK` where they meant `KO`. The audit questions support the statement in the project description that every instruction, including the final one, should have a `\n`.

## 4. Other People

At least three fellow 01 students from around the world have written Medium articles on their solutions: Jamie Dawson, Leo Fu, and Ali Yigit Ogun.

Reading these, it seems that the rules have varied slightly over time and space. We had two write a checker and a push-swap program, as did Ali Yigit Ogun at 42-Heilbronn, apparently; others only had to write push-swap while the checker was provided. Most, if not all, of the examples I've seen discussed online were required to write their program in C or C++. Ours had to be in Go. More significant: By 2023, at 01 Founders in London, we'd fail unless we could sort 100 mumbers in less than 700 of the specified operations. At Ecole 42, Lyon, in 2021, Leo Fu passed by sorting 100 numbers in "about 1084 instructions". On the other hand, he had to meet a minimum requirement for 500 numbers, whereas we weren't tested on 500 at all. Ali was happy to score 125/125. As for what this means, he links to a PDF of his school's instructions, but all they say on scoring is that if your list of instructions is "too big" it will fail. (It refers to a "maximum number tolerated" without specifying.) Similarly, at 42 Silicon Valley, 2019, Jamie Dawson needed to pass some requirement for 100 and 500, althouhgt he doesn't say how many instructions he was allowed. Of course, a dedicated push-swappist could persuse the commit histories of these various schools' public repos.

## 5. Bitmasks: a detour

In the `getInstructions` function of the `checker` program (located in `get-instructions.go`), we wanted to move the cursor up a line to eliminate the blank line that results when the user indicates that they've finished typing instructions by pressing enter on a line with no instructions.

However, we can't unconditionally move up a line because when the intructions are piped to the program, there is no blank line. Hence we check whether the input is from the terminal before moving up a line.

This is done by checking if the input is from a character device. The method `fi.Mode()` returns a bitmask, i.e. a number that represents a sequence of bits. In the case of `fi.Mode()`, these bits represent the file mode and permissions. One of them indicates whether the file is a character device. This should be true when the input is from the command line.

`os.ModeCharDevice` is a constant that indicates which bit, this is. It's value is 8192. In binary, 8192 is 10000000000000, a 1 followed by 13 zeros. So, being a power of two, 8192 can stand for a single bit, the 14th bit of the sequence.

The `&` is a bitwise AND operator. It's used here to check if the bit that represents a character device is set (i.e. 1). The result of this AND operation will only be zero when the 14th bit of `fi.Mode` is zero. Otherwise it will be 8192.
