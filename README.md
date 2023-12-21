# push-swap

0. [The brief](#0-the-brief)
1. [Getting started](#1-getting-started)
   - a. [Building binaries](#a-building-binaries)
   - b. [A quick note about newline characters](#b-a-quick-note-about-newline-characters)
2. [Audit](#2-audit)
3. [Research](#3-research)
4. [Structure and strategy](#4-structure-and-strategy)
5. [Mathematical observations](#5-mathematical-observations)
   - a. [Swaps and rotations are enough to sort](#a-swaps-and-rotations-are-enough-to-sort)
   - b. [Antipodeal elements: an optimization for stacks of even size](#b-antipodeal-elements-an-optimization-for-stacks-of-even-size)
   - c. [Why does Leo Fu's radix sort always take the same amount of instructions for a given stack size?](#c-why-does-leo-fus-radix-sort-always-take-the-same-amount-of-instructions-for-a-given-stack-size)
6. [Bitmasks: a detour](#6-bitmasks:-a-detour)

## 0. The brief

This is our take on the 01 Edu project `push-swap`, details of which can be found [here](https://github.com/01-edu/public/tree/master/subjects/push-swap).

We're given a list of numbers on a circular stack A, together with an empty stack B, and a set of eleven operations, including rotations of stacks (individually or together in the same direction), swapping the top two elements of a stack, and pushing numbers back and forth from one stack to another.

The object is to write a program that leaves the numbers sorted on stack A in ascending order. There are certain constraints on the length of the sequence of instructions to accomplish this. Only programs that can sort 5 numbers in less than 12 instructions, and 100 numbers in less than 700 instructions will pass.

## 1. Getting started

### a. Building binaries

You'll find the main package for each of the two programs, `checker` and `push-swap`, in the folders of those names in `cmd`.

To compile the `push-swap` program, cd into the corresponding directory and run `go build -o push-swap`. To use the program, enter `./push-swap`, followed by a string of integers to sort, separated by spaces.

To compile the `checker` program, run `go build -o checker`; and, to use it, type `./checker`, followed by a string of integers to sort, piping in the instructions like so:

```
echo -e "rra\npb\nsa\nrra\npa\n" | ./checker "3 2 1 0"
```

Alternatively, you can run `./checker "3 2 1 0"` (with your choice of initial values on stack `a`), then type instructions on the command line, pressing enter after each. When you've typed all your instructions, you can hit enter one last time to let the checker know you've finished.

In case you want to run `main_test.go`, be aware that it expects both these binaries to be built and in their eponymous folders.

### b. A quick note about newline characters

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

## 2. Audit

We've provided a Zsh script to run the audit questions. Make sure the executables are built in the correct folders before running it, as explained in the previous section. If you want to take this shortcut, cd to the `cmd/pushswap` and type `chmod +x audit.zsh`, then execute the script with `./audit.zsh`, assuming you have Zsh. If you have Bash, change the shebang at the beginning of the file to `#!/usr/bin/env bash`.

If auditing this way, be sure to verify that the script does actually do what the audit questions ask, and to consider the subjective questions at the end.

## 3. Research

The technique we used is essentially [Ali Yigit Ogun's "Turk algorithm"](https://medium.com/@ayogun/push-swap-c1f5d2d41e97), with some additional checks to find shorter sequences of instructions that avoid pushes. I recommend this [YouTube video by Thuy Quematon (Thuggonaut)](https://www.youtube.com/watch?v=wRvipSG4Mmk) for more detail.

He used an insertion sort of all but three numbers to stack B, then insertion sort again back to A. A refinement is that instead of always pushing the number at the top of the stack, he does a preliminary calculation before each push. He then pushes the number for which the least amount of rotations is needed to bring it to the top of its stack and its target to the top of the other stack.

As a special case, for 100 numbers, we followed [Fred 1000orion's method](https://www.youtube.com/watch?v=2aMrmWOgLvU): a bucket sort, with three buckets, onto stack B, then insertion sort to stack A with a cost check as per AYO.

I've see three other Medium articles on the subject: one by [Jamie Dawson](https://medium.com/@jamierobertdawson/push-swap-the-least-amount-of-moves-with-two-stacks-d1e76a71789a), one by [Leo Fu](https://medium.com/nerd-for-tech/push-swap-tutorial-fa746e6aba1e), and one by [Julien Caucheteux](https://medium.com/@julien-ctx/push-swap-an-easy-and-efficient-algorithm-to-sort-numbers-4b7049c2639a). JC does something similar to AYO. The others use quite different methods for sorting large stacks.

JC starts by pushing everything to stack B. He then sorts them back to A via insertion sort with cost check, choosing, at each iteration, to push back the number that can be correctly placed on A with the least rotations.

For 100 numbers, JD pushes the lowest 20 numbers to stack B first, then the next lowest, and so on till everything is sorted on B in ascending order (getting bigger as you go down), then pushes them back one by one. (Note that his original approach, that he calls insertion sort, sounds more like selection sort.)

One remark on JD's statement: "We’ll bring those numbers back once the three numbers in Stack A are sorted from smallest to largest." In his example, this happens to rotate stack A into the right position to receive the number at the top of stack B. In other cases, though, it might be counterproductive to rotate stack A all the way till the smallest number is on top. So we omitted this final step and just rotate stack A to where it needs to be before pushing each number back from B.

LF uses base 2 radix sort, of the Least Significant Digit flavor. For each bit, starting with the rightmost, he checks the numbers at the top of stack A. If the relevant bit is 0, he moves pushes to stack B; otherwise he applies `ra` to rotate it out of the way to the bottom of A. In this way, he goes through all the numbers. Then he pushes those that went to stack B back, and procedes in this way through all the bits. He says it didn't get him the highest score; presumably the cost of having to push numbers back and forth on multiple passes was too much. But this is an important technique to learn, particularly as the `push-swap` project description hinted that non-comparative sorting algorithms might be relevant.

It seems the push-swap rules have varied slightly over time and space. We had two write a checker and a push-swap program, as did AYO at 42-Heilbronn; others only had to write push-swap while the checker was provided. The projects I've seen discussed online were written in C or C++ (although the articles focus on strategy rather than implementation). Ours had to be in Go.

Different scoring systems are used by the various schools, which can sometimes offer clues about the performance of these folks' solutions. At Ecole 42, Lyon, in 2021, LF passed by sorting 100 numbers in "about 1084 instructions". He quotes a scoring system in which less than 700 is needed for top marks. He also had to meet a minimum requirement for 500 numbers, and got extra points according to how few instructions he could do it in. AYO says he scored 125/125. As for what this means, he links to a PDF of his school's instructions, but all they say on scoring is that if your list of instructions is "too big" it will fail. (It refers to a "maximum number tolerated" without specifying.) Similarly, at 42 Silicon Valley in 2019, JD needed to pass some requirement for 100 and 500, although he doesn't say how many instructions he was allowed. Of course, a dedicated push-swappist could persuse the commit histories of these various schools' public repos. By 2023, at 01 Founders in London, we'd get an unspecified bonus if we could sort 100 mumbers in less than 700 of the specified operations. No mention is made of 500 numbers in our audit.

On 10,000 tests, our implementation of FO's algorithm took an average of 555 instructions to sort 100 numbers, with a standard deviation of 24. (He says he sorted 100 with a mean of 510 instructions. I don't know how many tests he did. Four buckets performed somewhat worse at 569 instructions. The standard deviation was 23. To sort 500 numbers, he reports a mean of 3750 instructions; our version scored 4216 on 100 trials, with a standard deviation of 121.)

AYO's method achieved a mean of 561 instructions, with a standard deviation of 23, the worst cases being in the low 600s. (Without AYO's cost calculation, the mean was 1387, and the standard deviation 79. Our initial checks to see if the stack can be simply swapped and rotated into order made no difference in this test.)

JC's approach of pushing everything, then insertion sorting with cost checking like AYO on the way back took 584 instructions on average, with a standard deviation of 24.

LF reports "about 1084" instructions for 100 numbers, and "about 6756" for 500, then remarks that he actually always got exactly 6756, no matter how many times he tested it on different random numbers, and poses the question: Why? We'll return to this [shortly](#c-why-does-leo-fus-radix-sort-always-take-the-same-amount-of-instructions-for-a-given-stack-size).

## 4. Structure and strategy

You'll find the source code in several folders: `push-swap`, for the program that generates instructions to sort a given list of numbers, and `checker` for the program to check whether a given sequence of instructions sorts a given list of numbers, both in `cmd`.

The `ps` package is for functions and structs these programs share.

`cmd` also contains a folder called `explorer`, for exploring new ideas. Its `main.go` uses breadth-first search to pre-compute all solutions for five numbers that only use rotations and swaps and are shorter than those found the standard way (using pushes). The current version of the `push-swap` performs this BFS at runtime for stacks of size 6, 7, or 8. For longer lists, and certainly for 100 numbers, BFS is too slow to be practical at runtime, so we make do with checking some of the simpler cases where push-free sorts can be used.

Like AYO and JD, we dealt with initial stacks of less than six numbers as special cases. This was partly because they lend themselves to optimizations that would be prohibitively time-consuming for longer lists, and partly so we could treat these smaller problems as warm up exercises. The six permutations of three elements are easily checked by hand. With a ranking function to simplify the task, it didn't take much longer to find optimal solutions for the 24 permutations of four elements and hardcode them.

At 120 permutation, the five-number challenge seemed like a good place to start automating, particularly as JD uses it as an example to develop intuition for how one might proceed to the general case. Indeed, we followed his method.

Anyway, for the general case, we followed AYO's method, which extends JD's for five numbers. The first two numbers on the stack are pushed indiscriminately from A to B. AYO sorts the rest of the numbers from A to B so that they land in descending order. This makes it simpler to push them back into ascending order on A. Before deciding which number to push from A, he checks each one to find the cheapest.

The cheapest number to move is the one that miminizes

```
A + B - C
```

where `A` is number of rotations needed to bring this number to the top of stack A, `B` is the number of rotations needed to put its target number at the top of stack B, and `C` is any combined rotations. At least in our program, this calculation makes a big saving on instructions for sorting 100 numbers. Sequences of instructions tend to be in the 500s as opposed to the 1400s.

When pushing the cheapest number from A to B, its target is the biggest smaller number or, if there is no smaller number, the maximum of B. When pushing back, the target is the smallest bigger number or, if there is no bigger number, the minimum. For more detail, examples, and illustrations, see AYO's article and the video by TQ.

Indiscriminately pushing the first two numbers from A to B can result in cases where one or both are just pushed right back! We deal with this by checking first to see if the stack is already sorted, and by canceling out any "pb", "pa" subsequence from the list of instructions.

## 5. Mathematical observations

### a. Swaps and rotations are enough to sort

Any permutation of `n` elements, `{1, 2, 3, ..., n}`, can be expressed as a sequence of swaps and rotations, so, if we didn't care about how many instructions it takes to sort a stack, we could just use these two operations.

In the language of group theory, an `n`-cycle, such as `(1 2 3 ... n)` (i.e. the permutation that sends `1` to where `2` was, and `2` to where `3` was, and ..., and `n` to where `1` was), and a transposition of elements that are adjacent in this cycle, such as `(1 2)`, together generate the whole symmetric group on `n` elements. These statements are equivalent because the inverse of `(1 2 3 ... n)` (a reverse rotation) is a composition of `n - 1` instances of `(1 2 3 ... n)`, while `(1 2)` is its own inverse.

To see that they generate the whole symmetric group, we can use the fact every permutation can be expressed as a composition of transpositions. A transposition of neighboring elements can be achieved by rotating them into the top two spots on the stack, then performing the swap operation, then rotating them back to their original positions. To transpose elements that are a positive number `k` steps apart, place one of them at the top of the stack so that the other is no more than half `n` steps below it. (This is the maximum distance apart around the circlular stack that any pair of elements can be.) Then perform a swap, then `k - 1` rotations, each followed by a swap, and then `k - 1` reverse rotations, each also followed by a swap. The result will be that the elements have changed places.

While swaps and rotations are sufficient, it will sometimes be more efficient to push elements to the spare stack B for sorting.

### b. Antipodeal elements: an optimization for stacks of even size

Due to the circular nature of the stacks, the cheapest numbers to push will tend to be those near the top or the bottom. In other words, a number is actually furthest from the top when it's near the middle of the stack. If the stack has `n` elements indexed from `0` at the top, then those whose index is less than or equal to the floor of `n/2` will reach the top sooner when rotated upwards, while, for those whose index is greater than the floor of `n/2`, the top is reached soonest when they're rotated downwards. This means that, when `n` is an even number, there will be a middle element which takes either `n/2` upwards or `n/2` downwards rotations to reach the top. (Think how a clock, where even-numbered 12 is also 0, has such a middle/opposite/antipodeal element: 6.)

One consequence of this is that, if we need to rotate one stack, say, `r` times upwards, and the other stack has `2 * r` elements, then if we need to rotate the second stack `r` times, we can choose to rotate it upwards too, to take advantage of the combined rotation operation.

### c. Why does Leo Fu's radix sort always take the same amount of instructions for a given stack size?

As mentioned above, LF reports that his implementation of base 2 LSD radix sort took "about 1084" instructions for 100 numbers, and "about 6756" for 500. He imediately corrects himself, saying that he actually always got exactly 6756.

To see why this is (and why must also have always got exactly 1084), note first that he takes the convenient step of converting the original values to their rank: 0, 1, 2, 3, ..., 99. Now, ceil(log2(500)) = 9, so there will be 9 passes for the 9 bits needed to label 500 numbers in this way. In the case of 500 numbers, then, there will be at least 500 operations per bit. One might think that half the numbers (250 of them) would take a turn at being pushed to B each time, and thus have to be pushed back (resulting in 9 \* 750 = 6750 instructions)--and this would indeed be the case if 500 was a power of 2. But the bits of 499 are 111110011, so not every 9-bit sequenece of 0s and 1s is represented among the numbers to be sorted. Since it's the highest 12 numbers that are missing from the full total of 2^9 = 512 possible 9-bit sequences, it's the 0s that will be overrepresented in the total, shifting the balance in favor of pushes.

To take a simple example, suppose we had to sort 6 numbers. We'd need ceil(log2(6)) = 3 bits, and the numbers to sort would be expressed in binary form as:

```
000
001
010
011
100
101
```

There would be one operation (`ra` or `pb` for each of these numbers on each pass, and one pass for each of the three bits, hence at least 6 \* 3 = 18 operations. In addition, there will be a `pa` for every number that was pushed to stack B, which is to say, one more operation for every 0 that appears in this list. There are 11 zeros: 3 for the least significant bit, and 4 each for the others. In total, therefore, it will always take 18 + 11 = 29 operations to sort 6 numbers in this way.)

Notice that the two numbers missing to make up the next power of two are:

110
111

For a full power of 2, there are as many 0s as 1s for every bit. Since 0s and 1s are equally represented in the rightmost bit of the missing numbers, they must be equally represented in the rightmost bit of the 6 numbers we have: that is, 6/2 = 3 zeros. But there are no 0s among the other two bits of the missing numbers, so all 8/2 = 4 of the total possible zeros must be present among our 6 numbers. Hence it will take 6 \* 3 + 3 + 4 + 4 = 29 operations.

Similarly, ceil(log2(100)) = 7, so there will be at least 700 operations (rotations and pushes from A to B), and somewhat more than 7 \* 50 extra pushes, representing pushes back to A of those numbers that were moved there. This is a bit further off LF's actual score of 1084, which makes sense given that the difference between 100 and the next highest power of 2--or, equivalently, between 99 and 127 (2^7 - 1)--is greater than the difference between 500 and and 2^9.

## 6. Detour: bitmasks

In the `getInstructions` function of the `checker` program (located in `get-instructions.go`), we wanted to move the cursor up a line to eliminate the blank line that results when the user indicates that they've finished typing instructions by pressing enter on a line with no instructions.

However, we can't unconditionally move up a line because, when the intructions are piped to the program, there is no blank line. Hence we check whether the input is from the terminal before moving up a line.

This is done by checking if the input is from a character device. The method `fi.Mode()` returns a bitmask, i.e. a number that represents a sequence of bits. In the case of `fi.Mode()`, these bits represent the file mode and permissions. One of them indicates whether the file is a character device. This should be true when the input is from the command line.

`os.ModeCharDevice` is a constant that indicates which bit, this is. It's value is 8192. In binary, 8192 is 10000000000000, a 1 followed by 13 zeros. So, being a power of two, 8192 can stand for a single bit, the 14th bit of the sequence.

The `&` is a bitwise AND operator. It's used here to check if the bit that represents a character device is set (i.e. 1). The result of this AND operation will only be zero when the 14th bit of `fi.Mode` is zero. Otherwise it will be 8192.
