# push-swap

0. [The brief](#0-the-brief)
1. [Getting started](#1-getting-started)
   - a. [Building binaries](#a-building-binaries)
   - b. [A quick note about newline characters](#b-a-quick-note-about-newline-characters)
2. [Audit](#2-audit)
3. [Research](#3-research)
   - a. [Our peers at affiliated schools](#a-our-peers-at-affiliated-schools)
   - b. [Grading systems](#b-grading-systems)
   - c. [Results](#c-results)
4. [Structure and strategy](#4-structure-and-strategy)
5. [Mathematical curios](#5-mathematical-curios)
   - a. [Swaps and rotations are enough to sort](#a-swaps-and-rotations-are-enough-to-sort)
   - b. [Antipodeal elements: an optimization for stacks of even size](#b-antipodeal-elements-an-optimization-for-stacks-of-even-size)
   - c. [Why Leo Fu always got the same amount of instructions for a given stack size](#c-why-leo-fu-always-got-the-same-amount-of-instructions-for-a-given-stack-size)
6. [Detour: bitmasks](#6-detour-bitmasks)

## 0. The brief

This is our take on the 01 Edu project `push-swap`, details of which can be found [here](https://github.com/01-edu/public/tree/master/subjects/push-swap).

We're given a list of numbers on a circular stack A, together with an empty stack B, and a set of eleven operations, including rotations of stacks (individually or together in the same direction), swapping the top two elements of a stack, and pushing numbers back and forth from one stack to another.

The object is to write a program that leaves the numbers sorted on stack A in ascending order. There are certain constraints on the length of the sequence of instructions to accomplish this.

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

If you're running tests, you might want to use the `-count=1` flag to make sure the test is actually performed! Sometimes cached results are given. This can happen if you run the same test on the same code (perhaps hoping the test will generate different random numbers), or if you make a change to code in a different package from the test (which can be a problem if your code depends on the code you changed in that external package). Cached results can also be given if the test depends on something outside of the code, such as environment variables, command-line flags, or some other file or database. The `-count=1` flag ensures that the test is definitely run once. For example, to test the general function:

```
go test -run=TestGeneral -count=1
```

This will ensure that the test actually reflect the current state of the code!

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

### a. Our peers at affiliated schools

Our general technique is [Fred 1000orion's](https://www.youtube.com/watch?v=2aMrmWOgLvU): a bucket sort, with three buckets (that is, splitting the numbers into three equal chunks, then pushing all but three numbers to B in such a way that the bucket consisting of the smallest numbers is at the bottom of B, and that consisting of the largest numbers is at the top of B), then insertion sort to stack A with a cost check. That is to say, rather than just pushing whatever is at the top of stack B, we perform a preliminary calculation before each push and only pushes the number for which the least amount of rotations is needed to bring it to the top of B and its target to the top of A.

This is currently the best way we've found for large stacks. It assumes, though, that the there are at least 8 numbers. We've treated smaller stacks as special cases. We've hardcoded 3 and 4. For 5, we followed [Jamie Dawson](https://medium.com/@jamierobertdawson/push-swap-the-least-amount-of-moves-with-two-stacks-d1e76a71789a), who also treated it as a special case: push the top two numbers from A to B, sort the three on A, swap (or equivalently rotate) the two on B, insert back at the correct places, and rotate till the smallest is on top.

For 6 and 7, we use [Ali Yigit Ogun's self-titled "Turk algorithm"](https://medium.com/@ayogun/push-swap-c1f5d2d41e97). I recommend this [YouTube video by Thuy Quematon (Thuggonaut)](https://www.youtube.com/watch?v=wRvipSG4Mmk) for more detail. AYO used an [insertion sort](https://en.wikipedia.org/wiki/Insertion_sort) of all but three numbers to stack B, then insertion sort again back to A with a cost check as per FO. AYO explains the cost check well with an example.

For stack size less than 9, after finding our provisional result, we run a cheeky BFS to see if there's a shorter sequences of instructions that avoids pushes. Perhaps surprisingly, there often is. BFS beats Turk on 74 out of the 120 possible permutations of 5 numbers, 567 out of the 720 permutations of 6 numbers, and 3683 out of 5040 for 7. BFS becomes prohibitively time-consuming for larger stacks, so we just do some simple checks for low-hanging fruit: stacks that can be sorted with a swap, possibly preceded and possibly succeeded by rotations.

With the methods above, our program sorts all but three permutations of the numbers from 1 to 6 (inclusive) in under 13 instructions: "4 3 2 1 6 5", "2 6 5 4 3 1", "3 1 2 6 5 4". All take exactly 13. For the sake of neatness, we hardcode them to take less than 13.

I've read several other Medium articles on the subject: by [Leo Fu](https://medium.com/nerd-for-tech/push-swap-tutorial-fa746e6aba1e), [Julien Caucheteux](https://medium.com/@julien-ctx/push-swap-an-easy-and-efficient-algorithm-to-sort-numbers-4b7049c2639a), [Dan Sylvain](https://medium.com/@dansylvain84/my-implementation-of-the-42-push-swap-project-2706fd8c2e9f), and [YYBer](https://medium.com/@YYBer/my-one-month-push-swap-journey-explore-an-easily-understand-and-efficient-algorithm-11449eb17752).

DS's distinct feature is that he starts by finding the longest increasing subsequence (LIS) in the initial arrangement of A (ignoring its circularity, that is, and not wrapping around). He pushes everthing else to B in two buckets: numbers greater than the median go to the top half of B, the rest to the bottom. They're then insertion sorted back using a cost calculation as AYO does.

JC does something similar to AYO. His method is a bit simpler and a bit less effective, at least in our implementation. He starts by pushing everything to stack B. He then sorts them back to A via insertion sort with cost check, choosing, at each iteration, to push back the number that can be correctly placed on A with the least rotations.

As mentioned, we used JD's special-case technique for 5 numbers. JD's solution for lerger numbers, like FO's, involves splitting the numbers into buckets. In Jd's case, 5 for 100 numbers, 11 for 500. After a passable but less than optimally scoring attempt that he calls insertion sort, but which sounds more like [selection sort](https://en.wikipedia.org/wiki/Selection_sort), he settled on a form of bucket sort where he inserts each bucket's contents at the correct location inside the bucket as he fills it rather than, as FO does, first just dealing the numbers into buckets on stack B, and only after all the buckets are filled, inserting their contents into their correct locations on stack A.

Thus, for 100 numbers, JD pushes the lowest 20 numbers to stack B first, then the next lowest, and so on till everything is sorted on B in ascending order. (Why not descending, so that they'll naturally fall into ascending order when pushed back to A?) At each iteration, he pushes the number in the current range that takes least rotations to reach the top of stack A. Finally he pushes them back one by one, presumably rotating B between pushes to account for that fact that, within buckets, ascending order prevails. As with JC's method, all but trivial sorting happens in one direction. For JD, the sorting happens from A to B; for JC, from B to A.

One remark on JD's statement: "We’ll bring those numbers back once the three numbers in Stack A are sorted from smallest to largest." In his example, this happens to rotate stack A into the right position to receive the number at the top of stack B. In other cases, though, it might be counterproductive to rotate stack A all the way till the smallest number is on top. So we omitted this final step and just rotate stack A to where it needs to be before pushing each number back from B.

YYB, like many push-swappers, follows the convenience of ranking the numbers. She deals them from A into 8 equal-sized buckets on B for 100 numbers, or 12 for 500, then pushes the remainder. She then uses an ingenious cost checking procedure to push everything back to A in order. First the maximum number is pushed to A. Then, at each iteration, if the cheapest number to push back is the one that belongs on top of the top element of A, she pushes that and leaves it there; if the cheapest number is greater than the element on the bottom of A, she pushes it and rotates it with `ra` to the bottom of A. If, after a push, the element on the bottom of A turns out to be now the one that belongs on top of the top element, she rotates it up into place with `rra`. In this way she sorts everything back to A.

LF followes a quite different approach. He uses base 2 radix sort, of the Least Significant Digit flavor. For each bit, starting with the least significant (rightmost), he checks the number at the top of stack A. If the relevant bit is 0, he pushes the top number from A to B with `pb`; otherwise he applies `ra` to rotate it out of the way to the bottom of A. In this way, he goes through all the numbers in A. Then he pushes back everything from B with `pa`, and procedes in this way through all the bits. He says it didn't get him the highest score; presumably the cost of having to push all those numbers back and forth on multiple passes was too much. But radix sort is an important technique to learn, and it should be mentioned that the `push-swap` project description does direct our attention towards non-comparative sorting algorithms.

No doubt [btilly's idea](https://stackoverflow.com/questions/75100698/push-swap-what-is-the-most-efficient-way-to-sort-given-values-using-a-limited-s) of sharing out runs between the two stacks and merging them back and forth would suffer from the same issue: too much pushing.

A few others have made their solutions public on GitHub, such as [Adrian Roque](https://github.com/AdrianWR/push_swap), who credits [Anya Schukin](https://github.com/anyaschukin/Push_Swap) with this idea. They deal into buckets on stack B, like FO, except that Anya says she used 2 buckets when there are no more than 100 numbers, or else 4 buckets when there are 500 or less. The numbers are then pushed back to A using a kind of selection sort, except with the refinement that the minimum OR THE MAXIMUM are valid options to push, depending on which is cheaper, thus taking advantage of the circularity of the stack.

### b. Grading systems

It seems the push-swap rules have varied slightly over time and space. We had two write a checker and a push-swap program, as did AYO at 42-Heilbronn; others only had to write push-swap while the checker was provided. The projects I've seen discussed online were written in C or C++ (although the articles focus on strategy rather than implementation). Ours had to be in Go.

Different scoring systems are used by the various schools, which can sometimes offer clues about the performance of folks' solutions when they don't go into specifics. At Ecole 42, Lyon, in 2021, [Leo Fu](https://medium.com/nerd-for-tech/push-swap-tutorial-fa746e6aba1e) passed by sorting 100 numbers in "about 1084 instructions". He quotes a scoring system in which top marks are gained by sorting 100 numbers in less than 700 instructions, and 500 in less than 5500.

AYO says he scored 125/125. As for what this means, he links to a PDF of his school's instructions, but all they say on scoring is that if your list of instructions is "too big" it will fail. (It refers to a "maximum number tolerated" without specifying.) Similarly, at 42 Silicon Valley in 2019, JD needed to pass some requirement for 100 and 500, although he doesn't say how many instructions he was allowed. Of course, a dedicated push-swappist could persuse the commit histories of these various schools' public repos.

[Dan Sylvain](https://medium.com/@dansylvain84/my-implementation-of-the-42-push-swap-project-2706fd8c2e9f) presents a similar grading system to LF, but talks about a "base project" having to be perfect before you can conplete the "bonus part", as if either the 100 or the 500 test are bonuses.

At any rate, by December 2023, at 01 Founders in London, we'd get an unspecified bonus if we could sort 100 mumbers in less than 700 of the specified operations. No mention is made of 500 numbers in our [audit questions](https://github.com/01-edu/public/tree/master/subjects/push-swap/audit). Altogether, it looks like we could have an easier ride--if we wanted it..

### Results

On 10,000 tests, our implementation of FO's algorithm took an average of 555 instructions to sort 100 numbers, with a standard deviation of 24. Of all the ways we've tried so far, this is the winner, and it sounds like FO may have achieved some further optimization that wasn't revealed in his summary, unless I've overlooked it. He says he sorted 100 with a mean of 510 instructions. I don't know how many tests he did. To sort 500 numbers, he reports a mean of 3750 instructions; our version scored 4216 on 100 trials, with a standard deviation of 121.

We tried varying the number of buckets used in this technique: Four buckets performed somewhat worse at 569 instructions (standard deviation: 23), and two buckets even worse at 573 (standard deviation: 26).

AYO's method achieved a mean of 561 instructions, with a standard deviation of 23, the worst cases being in the low 600s. (Without AYO's cost calculation, the mean was 1387, and the standard deviation 79. Our initial checks to see if the stack can be simply swapped and rotated into order made no difference in this test.)

DS did pretty well with his strategy of and sorting all but the longest increasing subsequence into two buckets on B, then insertion sorting back to A with a cost check. Our implementation scored a mean of 566, with a standard deviation of 21.

JC's approach of pushing everything, then insertion sorting with cost checking like AYO on the way back took 584 instructions on average, with a standard deviation of 24.

YYB first scored 750 for 100 numbers by pushing the smallest half to A, then the smallest half of the rest, and so on till A is empty, and then selection sorting them back to A. After this attempt, she switched to a new strategy: bucket sort to B with 8 equal-sized buckets in the case of 100 numbers, or 12 in the case of 500 (although it sounds like the buckets are not arranged in descending order on B?), then push whatever is left till A is empty. Then selection sort back to A. This required less than 700 instructions for 100 numbers, and about 6500 for 500. As a final optimization, she replaced selection sort with the ingenious cost-checking procedure, described above, that reduced her mean score for 500 numbers to "5300-5400 steps".

LF reports "about 1084" instructions for 100 numbers, and "about 6756" for 500, then remarks that he actually always got exactly 6756, no matter how many times he tested it on different random numbers, and poses the question: Why? We'll return to this [shortly](#c-why-does-leo-fus-radix-sort-always-take-the-same-amount-of-instructions-for-a-given-stack-size).

LR: We also tried leaving the longest run (i.e. the longest sequence of numbers adjacent in the initial stack, such that their ranks are consecutive integers) on A and pushing everything else into two buckets on B, after which we insertion sorted them back with a cost check. This resulted in a mean of 577 instructions and a standard deviation of 25. But then, the length of the longest run, for 100 uniformly distributed random numbers, is mostly 1 or 2.

FO(3) 555
AYO 561
DS(2) 566
FO(4) 569
FO(2) 573
LR 577
DS(3) 578
JC 584
LF 1084
AYO(-cost) 1387

## 4. Structure and strategy

You'll find the source code in several folders: `push-swap`, for the program that generates instructions to sort a given list of numbers, and `checker` for the program to check whether a given sequence of instructions sorts a given list of numbers, both in `cmd`.

The `ps` package is for functions and structs these programs share.

`cmd` also contains a folder called `explorer`, for exploring new ideas, such as the advantages of BFS.

Like JD (and many of the others), we dealt with initial stacks of less than six numbers as special cases. This was partly because they lend themselves to optimizations that would be prohibitively time-consuming for longer lists, and partly so we could treat these smaller problems as warm up exercises. The six permutations of three elements are easily checked by hand. With a ranking function to simplify the task, it didn't take much longer to find optimal solutions for the 24 permutations of four elements and hardcode them.

At 120 permutation, the five-number challenge seemed like a good place to start automating, particularly as JD uses it as an example to develop intuition for how one might proceed to the general case. Indeed, we followed his method.

Anyway, for the general case, we followed AYO's method, which resembles JD's for five numbers, except for AYO's cost check. The first two numbers on the stack are pushed indiscriminately from A to B. AYO sorts the rest of the numbers from A to B so that they land in descending order. This makes it simpler to push them back into ascending order on A. Before deciding which number to push from A, he checks each one to find the cheapest.

The cheapest number to move is the one that miminizes

```
A + B - C
```

where `A` is number of rotations needed to bring this number to the top of stack A, `B` is the number of rotations needed to put its target number at the top of stack B, and `C` is any combined rotations. At least in our program, this calculation makes a big saving on instructions for sorting 100 numbers, as detailed in the previous section.

When pushing the cheapest number from A to B, its target is the biggest smaller number or, if there is no smaller number, the maximum of B. When pushing back, the target is the smallest bigger number or, if there is no bigger number, the minimum. For more detail, examples, and illustrations, see AYO's article and the video by TQ.

Indiscriminately pushing the first two numbers from A to B can result in cases where one or both are just pushed right back! We deal with this by checking first to see if the stack is already sorted, and by canceling out any "pb", "pa" subsequence from the list of instructions.

## 5. Mathematical curios

### a. Swaps and rotations are enough to sort

Any permutation of `n` elements, `{1, 2, 3, ..., n}`, can be expressed as a sequence of swaps and rotations, so, if we didn't care about how many instructions it takes to sort a stack, we could just use these two operations.

In the language of group theory, an `n`-cycle, such as `(1 2 3 ... n)` (i.e. the permutation that sends `1` to where `2` was, and `2` to where `3` was, and ..., and `n` to where `1` was), and a transposition of elements that are adjacent in this cycle, such as `(1 2)`, together generate the whole symmetric group on `n` elements. These statements are equivalent because the inverse of `(1 2 3 ... n)` (a reverse rotation) is a composition of `n - 1` instances of `(1 2 3 ... n)`, while `(1 2)` is its own inverse.

To see that they generate the whole symmetric group, we can use the fact every permutation can be expressed as a composition of transpositions. A transposition of neighboring elements can be achieved by rotating them into the top two spots on the stack, then performing the swap operation, then rotating them back to their original positions. To transpose elements that are a positive number `k` steps apart, place one of them at the top of the stack so that the other is no more than half `n` steps below it. (This is the maximum distance apart around the circlular stack that any pair of elements can be.) Then perform a swap, then `k - 1` rotations, each followed by a swap, and then `k - 1` reverse rotations, each also followed by a swap. The result will be that the elements have changed places.

Since any permutation can be sorted on one stack, it follows that it can be sorted when two are available, an unspoken assumption of this project!

### b. Antipodeal elements: an optimization for stacks of even size

Due to the circular nature of the stacks, the cheapest numbers to push will tend to be those near the top or the bottom. In other words, a number is actually furthest from the top when it's near the middle of the stack. If the stack has `n` elements indexed from `0` at the top, then those whose index is less than or equal to the floor of `n/2` will reach the top sooner when rotated upwards, while, for those whose index is greater than the floor of `n/2`, the top is reached soonest when they're rotated downwards. This means that, when `n` is an even number, there will be a middle element (we could call it opposite or antipodeal) which takes either `n/2` upwards or `n/2` downwards rotations to reach the top. (Think how a clock, where even-numbered 12 is also 0, has such an opposite element: 6.)

One consequence of this is that, if we need to rotate one stack, say, `r` times upwards, and the other stack has `2 * r` elements, then if we need to rotate the second stack `r` times, we can choose to rotate it upwards too, to take advantage of the combined rotation operation.

### c. Why Leo Fu always got the same amount of instructions for a given stack size

As mentioned above, LF reports that his implementation of base 2 LSD radix sort took "about 1084" instructions for 100 numbers, and "about 6756" for 500. He imediately corrects himself, saying that he actually always got exactly 6756.

To see how this can be, and why he must also have always got exactly 1084 for 100 numbers, note first that LF takes the convenient step of converting the original values to their rank: 0, 1, 2, 3, 4, ..., 98, 99.

Ceil(log2(99)) = 7, so it takes 7 bits to express these numbers. Picture them listed in binary form, padding the smaller ones with leading zeros:

0000000
0000001
0000010
0000011
0000100
...
1100010
1100011

In sorting, there is one pass (iteration) for every bit, starting with the least significant, i.e. at the right. At every pass, all the numbers with a 0 at that bit will be pushed to B with `pb` and eventually back with `pa`, while all the numbers with a 1 at that bit will be rotated out of the way with `ra`.

If the size of the stack was a power of 2, there would be an equal number of ones and zeros in the list of numbers (in binary form) to be sorted. But the bits of 99 are 1100011, so not every 7-bit sequenece of 0s and 1s is represented among the numbers to be sorted. Since it's the greatest 28 numbers that are missing between 99 and 127 (one less than the next power of 2), i.e. from the full total of all possible 7-bit sequences, it's the 0s that will be overrepresented in the total, shifting the balance in favor of pushes.

We can calculate the exact amount of instructions as the number of ones in the list, plus twice the number of zeros, `ones + 2 * zeros`, or alternatively as `bits * n + zeros` (with n = 100 in this case; LF gives the latter formula in a comment on his article), since every number requires at least one instruction per bit (whether a rotation or a push), and then those with zero at that bit need one extra instruction to push them back. Sure enough, both formulas give 1084 for n = 100.

Similarly, ceil(log2(499)) = 9, so it takes 9 bits to represent the numbers from 0 to 499. It's the highest 12 numbers that are missing from the full total of 2^9 = 512 possible 9-bit sequences, so, as always when the stack size is not a power of 2, the zeros will be overrepresented, and there will be more pushes than rotations.

At this point, though, we have a a bit of a mystery. By my calculation, `9 * 500 + zeros` and `ones + 2 * zeros` both give 6784 rather than LF's 6756. I'm quite puzzled by this. In his article, LF actually shows a screenshot of a checker program showing the result 6756 for every random sequence of 500 numbers it tested, so it's not typo. I wonder if he introduced some optimization for stack sizes greater than 100 that I overlooked or that he didn't mention. Maybe I'm just missing something. I've written to him, asking if he has any insight. I'll update this if I learn more.

## 6. Detour: bitmasks

In the `getInstructions` function of the `checker` program (located in `get-instructions.go`), we wanted to move the cursor up a line to eliminate the blank line that results when the user indicates that they've finished typing instructions by pressing enter on a line with no instructions.

However, we can't unconditionally move up a line because, when the intructions are piped to the program, there is no blank line. Hence we check whether the input is from the terminal before moving up a line.

This is done by checking if the input is from a character device. The method `fi.Mode()` returns a bitmask, i.e. a number that represents a sequence of bits. In the case of `fi.Mode()`, these bits represent the file mode and permissions. One of them indicates whether the file is a character device. This should be true when the input is from the command line.

`os.ModeCharDevice` is a constant that indicates which bit, this is. It's value is 8192. In binary, 8192 is 10000000000000, a 1 followed by 13 zeros. So, being a power of two, 8192 can stand for a single bit, the 14th bit of the sequence.

The `&` is a bitwise AND operator. It's used here to check if the bit that represents a character device is set (i.e. 1). The result of this AND operation will only be zero when the 14th bit of `fi.Mode` is zero. Otherwise it will be 8192.
