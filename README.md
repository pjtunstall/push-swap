# push-swap

0. [The brief](#0-the-brief)
1. [Getting started](#1-getting-started)
2. [A quick note about newline characters](#2-a-quick-note-about-newline-characters)
3. [Research](#-research)
4. [Structure and strategy](#-structure-and-strategy)
5. [Mathematical](#-mathematical-observations)
6. [Bitmasks: a detour](#-bitmasks:-a-detour)

## 0. The brief

This is our take on the 01 Edu project `push-swap`, details of which can be found [here](https://github.com/01-edu/public/tree/master/subjects/push-swap).

We're given a list of numbers on a circular stack A, together with an empty stack B, and a set of eleven operations, including rotations of stacks (individually or together in the same direction), swapping the top two elements of a stack, and pushing numbers back and forth from one stack to another.

The object is to write a program that leaves the numbers sorted on stack A in ascending order. There are certain constraints on the length of the sequence of instructions to accomplish this. Only programs that can sort 5 numbers in less than 12 instructions, and 100 numbers in less than 700 instructions will pass.

## 1. Getting started

You'll find the main package for each of the two programs, `checker` and `push-swap`, in the folders of those names in `cmd`.

To compile the `push-swap` program, cd into the corresponding directory and run `go build -o push-swap`. To use the program, enter `./push-swap`, followed by a string of integers to sort, separated by spaces.

To compile the `checker` program, run `go build -o checker`; and, to use it, type `./checker`, followed by a string of integers to sort, piping in the instructions like so:

```
echo -e "rra\npb\nsa\nrra\npa\n" | ./checker "3 2 1 0"
```

Alternatively, you can run `./checker "3 2 1 0"` (with your choice of initial values on stack `a`), then type instructions on the command line, pressing enter after each. When you've typed all your instructions, you can hit enter one last time to let the checker know you've finished.

In case you want to run `main_test.go`, be aware that it expects both these binaries to be built and in their eponymous folders.

We've provided a Zsh script to run the audit questions. Make sure the executables are built in the correct folders before running it. If you want to take this shortcut, cd to the `cmd/pushswap` and type `chmod +x audit.zsh`, then execute the audit script with `./audit.zsh`, assuming you have Zsh. If you have Bash, change the shebang at the beginning of the file to `#!/usr/bin/env bash`. (Don't worry about the extension.)

If auditing this way, be sure to verify that the script does actually do what the audit questions ask, and to consider the subjective questions at the end.

## 2. A quick note about newline characters

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

## 3. Research

The technique we used is essentially [Ali Yigit Ogun's "Turk algorithm"](https://medium.com/@ayogun/push-swap-c1f5d2d41e97), with some additional checks to find shorter sequences of instructions that avoid pushes. I recommend this [YouTube video by Thuy Quematon (Thuggonaut)](https://www.youtube.com/watch?v=wRvipSG4Mmk) for more detail.

Two other Medium articles on the subject are also worth consulting as they use quite different methods to sort larger stacks: one by [Jamie Dawson](https://medium.com/@jamierobertdawson/push-swap-the-least-amount-of-moves-with-two-stacks-d1e76a71789a) and one by [Leo Fu](https://medium.com/nerd-for-tech/push-swap-tutorial-fa746e6aba1e).

For 100 numbers, JD pushes the lowest 20 numbers to stack B first, then the next lowest, and so on till everything is sorted on stack B in ascending order (getting bigger as you go down), then pushes them back one by one.

By contrast, AYO sorts all but three numbers onto stack B in descending order. That way, he can push them back without having to make an extra rotation each time to bring the right number to the top, and they'll arrive naturally in the right order. He calculates the cheapest element to push from stack A each time, based on how many rotations it would take to bring both it and its target number in B to the tops of their respective stacks.

LF uses radix sort. He says it didn't get him the highest score; presumably the cost of having to push numbers back and forth on multiple passes was too much. But this is an important technique to learn, and his article makes an interesting read, particularly as the `push-swap` project description hinted that non-comparative sorting algorithms might be relevant.

It seems the push-swap rules have varied slightly over time and space. We had two write a checker and a push-swap program, as did AYO at 42-Heilbronn; others only had to write push-swap while the checker was provided. The projects I've seen discussed online were written in C or C++ (although the articles focus on strategy rather than implementation). Ours had to be in Go.

Different scoring systems are used by the various schools, which can sometimes offer clues about the performance of these folks' solutions. At Ecole 42, Lyon, in 2021, LF passed by sorting 100 numbers in "about 1084 instructions". He quotes a scoring system in which less than 700 is needed for top marks. He also had to meet a minimum requirement for 500 numbers, and got extra points according to how few instructions he could do it in. AYO says he scored 125/125. As for what this means, he links to a PDF of his school's instructions, but all they say on scoring is that if your list of instructions is "too big" it will fail. (It refers to a "maximum number tolerated" without specifying.) Similarly, at 42 Silicon Valley in 2019, JD needed to pass some requirement for 100 and 500, although he doesn't say how many instructions he was allowed. Of course, a dedicated push-swappist could persuse the commit histories of these various schools' public repos. By 2023, at 01 Founders in London, we'd get an unspecified bonus if we could sort 100 mumbers in less than 700 of the specified operations. No mention is made of 500 numbers in our audit.

In any case, testing our program on 1000 instances of 100 random numbers, the mean number of instructions produced is consistently 560 or 561, with a standard deviation of 23 or 24. There are usually between 50 and 60 that score in the low 600s. I've never seen any worse than that.

Without AYO's cost calculation to find the cheapest number to push from stack A to stack B while sorting, if we always push the number at the top of stack A, every one of these 1000 instances of 100 numbers takes more than 700 instructions. The mean number of instructions is in the 1380s, and the standard deviation in the high 70s or low 80s.

## 4. Structure and strategy

You'll find the source code in several folders: `push-swap`, for the program that generates instructions to sort a given list of numbers, and `checker` for the program to check whether a given sequence of instructions sorts a given list of numbers, both in `cmd`.

The `ps` package is for functions and structs these programs share.

`cmd` also contains a folder called `explorer`, for exploring new ideas. Its `main.go` uses breadth-first search to pre-compute all solutions for five numbers that only use rotations and swaps and are shorter than those found the standard way (using pushes). The current version of the `push-swap` performs this BFS at runtime for stacks of size 6, 7, or 8. For longer lists, and certainly for 100 numbers, BFS is too slow to be practical at runtime, so we make do with checking some of the simpler cases where push-free sorts can be used.

Like AYO and JD, we dealt with initial stacks of less than six numbers as special cases. This was partly because they lend themselves to optimizations that would be prohibitively time-consuming for longer lists, and partly so we could treat these smaller problems as warm up exercises. The six permutations of three elements are easily checked by hand. With a ranking function to simplify the task, it didn't take much longer to find optimal solutions for the 24 permutations of four elements and hardcode them.

At 120 permutation, the five-number challenge seemed like a good place to start automating, particularly as JD uses it as an example to develop intuition for how one might proceed to the general case. Indeed, we followed his method.

One remark on JD's statement: "We’ll bring those numbers back once the three numbers in Stack A are sorted from smallest to largest." In his example, this happens to rotate stack A into the right position to receive the number at the top of stack B. In other cases, though, it might be counterproductive to rotate stack A all the way till the smallest number is on top. So we omit this final step and just rotate stack A to where it needs to be before pushing each number back from B.

Anyway, for the general case, we followed AYO's method, which extends JD's for five numbers. The first two numbers on the stack are pushed indiscriminately from A to B. AYO sorts the rest of the numbers from A to B so that they land in descending order. This makes it simpler to push them back into ascending order on A. Before deciding which number to push from A, he checks each one to find the cheapest.

The cheapest number to move is the one that miminizes

```
A + B - C
```

where `A` is number of rotations needed to bring this number to the top of stack A, `B` is the number of rotations needed to put its target number at the top of stack B, and `C` is any combined rotations. At least in our program, this calculation makes a big saving on instructions for sorting 100 numbers. Sequences of instructions tend to be in the 500s as opposed to the 1400s.

When pushing the cheapest number from A to B, its target is the biggest smaller number or, if there is no smaller number, the maximum of B. When pushing back, the target is the smallest bigger number or, if there is no bigger number, the minimum. For more detail, examples, and illustrations, see AYO's article and the video by TQ.

Indiscriminately pushing the first two numbers from A to B can result in cases where one or both are just pushed right back! We deal with this by checking first to see if the stack is already sorted, and by canceling out any "pb", "pa" subsequence from the list of instructions.

## 5. Mathematical observations

Any permutation of `n` elements, `{1, 2, 3, ..., n}`, can be expressed as a sequence of swaps and rotations, so, if we didn't care about how many instructions it takes to sort a stack, we could just use these two operations.

In the language of group theory, an `n`-cycle, such as `(1 2 3 ... n)` (i.e. the permutation that sends `1` to where `2` was, and `2` to where `3` was, and ..., and `n` to where `1` was), and a transposition of elements that are adjacent in this cycle, such as `(1 2)`, together generate the whole symmetric group on `n` elements. These statements are equivalent because the inverse of `(1 2 3 ... n)` (a reverse rotation) is a composition of `n - 1` instances of `(1 2 3 ... n)`, while `(1 2)` is its own inverse.

To see that they generate the whole symmetric group, we can use the fact every permutation can be expressed as a composition of transpositions. A transposition of neighboring elements can be achieved by rotating them into the top two spots on the stack, then performing the swap operation, then rotating them back to their original positions. To transpose elements that are a positive number `k` steps apart, place one of them at the top of the stack so that the other is no more than half `n` steps below it. (This is the maximum distance apart around the circlular stack that any pair of elements can be.) Then perform a swap, then `k - 1` rotations, each followed by a swap, and then `k - 1` reverse rotations, each also followed by a swap. The result will be that the elements have changed places.

While swaps and rotations are sufficient, it will sometimes be more efficient to push elements to the spare stack B for sorting.

Due to the circular nature of the stacks, the cheapest numbers to push will tend to be those near the top or the bottom. In other words, a number is actually furthest from the top when it's near the middle of the stack. If the stack has `n` elements indexed from `0` at the top, then those whose index is less than or equal to the floor of `n/2` will reach the top sooner when rotated upwards, while, for those whose index is greater than the floor of `n/2`, the top is reached soonest when they're rotated downwards. This means that, when `n` is an even number, there will be a middle element which takes either `n/2` upwards or `n/2` downwards rotations to reach the top. (Think how a clock, where even-numbered 12 is also 0, has such a middle/opposite/antipodeal element: 6.)

One consequence of this is that, if we need to rotate one stack, say, `r` times upwards, and the other stack has `2 * r` elements, then if we need to rotate the second stack `r` times, we can choose to rotate it upwards too, to take advantage of the combined rotation operation.

## 6. Detour: bitmasks

In the `getInstructions` function of the `checker` program (located in `get-instructions.go`), we wanted to move the cursor up a line to eliminate the blank line that results when the user indicates that they've finished typing instructions by pressing enter on a line with no instructions.

However, we can't unconditionally move up a line because, when the intructions are piped to the program, there is no blank line. Hence we check whether the input is from the terminal before moving up a line.

This is done by checking if the input is from a character device. The method `fi.Mode()` returns a bitmask, i.e. a number that represents a sequence of bits. In the case of `fi.Mode()`, these bits represent the file mode and permissions. One of them indicates whether the file is a character device. This should be true when the input is from the command line.

`os.ModeCharDevice` is a constant that indicates which bit, this is. It's value is 8192. In binary, 8192 is 10000000000000, a 1 followed by 13 zeros. So, being a power of two, 8192 can stand for a single bit, the 14th bit of the sequence.

The `&` is a bitwise AND operator. It's used here to check if the bit that represents a character device is set (i.e. 1). The result of this AND operation will only be zero when the 14th bit of `fi.Mode` is zero. Otherwise it will be 8192.
