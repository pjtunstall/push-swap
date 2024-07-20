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

We're given a list of numbers on a circular stack A, together with an empty circular stack B, and a set of eleven operations: rotation (`ra`, `rb`) and reverse rotation (`rra`, `rrb`) of a stack, swapping the top two elements of a stack (`sa`, `sb`), pushing numbers from A to B or vice-versa (`pa`, `pb`), and performing a combined rotation (`rr`), reverse rotation (`rrr`), or swap (`ss`) of both stacks.

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

If you're running tests, you might want to use the `-count=1` flag to make sure the test is actually performed! Sometimes cached results are given. This can happen if you repeat a test without changing your code (perhaps hoping the test will generate different random numbers), or if you make a change to code in a different package from the test (which can be a problem if your code depends on the code you changed in that external package). Cached results can also be given if you run your code again after changing environment variables, command-line flags, or some other file or database. The `-count=1` flag ensures that the test is definitely run once. For example, to test the general function:

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

We've provided a Zsh script to run the audit questions. Make sure the executables are built in the correct folders before running it, as explained in the previous section. If you want to take this shortcut, cd to the `cmd/push-swap` and type `chmod +x audit.zsh`, then execute the script with `./audit.zsh`, assuming you have Zsh. If you have Bash, change the shebang at the beginning of the file to `#!/usr/bin/env bash`.

If auditing this way, be sure to verify that the script does actually do what the audit questions ask, and to consider the subjective questions at the end.

Thanks!

## 3. Research

### a. Our peers at affiliated schools

Our general technique, for stack size 93 or more, is essentially that of [Fred 1000orion](https://www.youtube.com/watch?v=2aMrmWOgLvU), who triages the numbers into three buckets on stack B, then [insertion sorts](https://en.wikipedia.org/wiki/Insertion_sort) them back to A.

In more detail, he splits the numbers into three more-or-less equal-sized chunks, consisting of the smallest, middle, and largest values. While more than one third of the numbers are left on A, he checks the element at the top of A to see which third it belongs to. If it's among the biggest numbers, he uses `ra` to move it out of the way to the bottom of the stack. If it's in the third consisting of the middle-sized numbers, he just pushes it to B with `pb`. If it's among the smallest numbers, he pushes it to B and rotates it to the bottom with `pb rb`.

When only one third of the numbers are left on A, he keeps pushing whatever is at the top of A with `pb` till there are only 3 numbers left on A. He sorts those three in place on A, then moves the rest back to A. Before each `pa`, he performs a cost check for each element of B to see which will need the fewest rotations of A and B to put it in the right place on A. He then pushes the cheapest number.

This is currently the best way we've found for large stacks. It does assume that the there are at least 8 numbers (8 specifically, due to how we chose round for stacks not divisible by 3), but this turned out not to be relevant in our program, since [Ali Yigit Ogun's self-titled "Turk" algorithm](https://medium.com/@ayogun/push-swap-c1f5d2d41e97) proved better at sorting stacks smaller than 93. We recommend this [YouTube video by Thuy Quematon (Thuggonaut)](https://www.youtube.com/watch?v=wRvipSG4Mmk) for more detail.

Ali used an [insertion sort](https://en.wikipedia.org/wiki/Insertion_sort) of all but three numbers to stack B in descending order, then insertion sort again back to A in ascending order, with a cost check in both directions as per Fred. That is to say, at each iteration, he pushes the number for which the least amount of rotations is needed to bring it to the top of its own stack and its target (i.e. the number that should be beneath it) to the top of the other stack. Ali explains the cost check well with an example.

Curiously, Ali's method performs better and better, compared to Fred's, as stack size increases from 8 to 42, after which Fred's starts to improve, eventually overtaking at it 93 (judging both by mean number of instructions and percent of cases where Ali takes fewer instructions than Fred). For example, to sort 8 random numbers, Ali takes fewer instructions 60% of the time. To sort 42 random numbers, Ali wins 91% of the time. It's only at 93 that Ali is finally bested, winning just 49% of the time. And to sort 100 numbers, he's down to 43%.

We've treated the very smallest stacks as special cases. We hardcoded 3 and 4. For 5, we followed [Jamie Dawson](https://medium.com/@jamierobertdawson/push-swap-the-least-amount-of-moves-with-two-stacks-d1e76a71789a), who also treated it as a special case: Push the top two numbers from A to B, sort the three on A, swap (or equivalently rotate) the two on B, insert back at the correct places, and rotate till the smallest is on top.

For stack sizes 5, 6, 7, and 8, after finding our provisional result, we run a cheeky BFS to see if there's a shorter sequences of instructions that avoids pushes. Perhaps surprisingly, there often is. BFS beats Turk on 74 out of the 120 possible permutations of 5 numbers (62%), 567 out of the 720 permutations of 6 numbers (79%), and 3683 out of 5040 for 7 (73%). BFS becomes prohibitively time-consuming for larger stacks, so we make do with some simple checks for low-hanging fruit: stacks that are already sorted, and those that can be sorted with rotations alone, or with a swap, possibly preceded and possibly succeeded by rotations.

With the methods above, our program sorts all but two permutations of the numbers from 1 to 6 (inclusive) in under 13 instructions. The exceptions are: "2 6 5 4 3 1" and "3 1 2 6 5 4". Both take exactly 13. For the sake of neatness, we've hardcoded them to take less.

There are several other Medium articles on the subject. We've read ones by [Leo Fu](https://medium.com/nerd-for-tech/push-swap-tutorial-fa746e6aba1e), [Julien Caucheteux](https://medium.com/@julien-ctx/push-swap-an-easy-and-efficient-algorithm-to-sort-numbers-4b7049c2639a), [Dan Sylvain](https://medium.com/@dansylvain84/my-implementation-of-the-42-push-swap-project-2706fd8c2e9f), [YYBer](https://medium.com/@YYBer/my-one-month-push-swap-journey-explore-an-easily-understand-and-efficient-algorithm-11449eb17752), and [Luca Fischer](https://medium.com/@lucafischer_11396/two-stacks-one-goal-understanding-the-push-swap-algorithm-e08e5986f657).

Dan's distinct feature is that he starts by finding the longest increasing subsequence (LIS) in the initial arrangement of A (ignoring its circularity, that is, and not wrapping around). He pushes everthing else to B in two buckets: numbers greater than the median go to the top half of B, the rest to the bottom. They're then insertion sorted back using a cost check like Ali and Fred.

Julien does something similar to Ali. His method is a bit simpler and a bit less effective at sorting 100 numbers, at least in our implementation. He starts by pushing everything to stack B. He then sorts them back to A via insertion sort with cost check, choosing, at each iteration, to push back the number that can be correctly placed on A with the least rotations.

As mentioned, aside from optimizations to combine rotations and eliminate needless pushes, we use Jamie Dawson's special-case technique for 5 numbers. He pushes the top two numbers indiscriminately. Now, Jamie says, "We’ll bring those numbers back once the three numbers in Stack A are sorted from smallest to largest." In his example, this happens to rotate stack A into the right position to receive the number at the top of stack B. In other cases, though, it might be counterproductive to rotate stack A all the way till the smallest number is on top. So we omit this step and just make a swap if needed to make A rotatable into the correct order. Then, before each push back, we rotate stack A to where it needs to be to receive the returning number.

For larger stacks, Jamie first tried what he calls [insertion sort](https://en.wikipedia.org/wiki/Insertion_sort), even linking to this Wikipedia article, but which sounds more like [selection sort](https://en.wikipedia.org/wiki/Selection_sort). At each step, he pushed the minumum number from A and, in this way, moved all the numbers to B in descending order. He then pushed everything back to A. This, he says, sorted the numbers, but was over the maximum length requirement to pass. (He doesn't say what that requirement was, but "under 1500" is the lowest scoring result in the table quoted by Leo Fu.)

The solution he eventually settled on is a variant of [bucket sort](https://en.wikipedia.org/wiki/Bucket_sort) except that, unconventionally, he inserts each bucket's contents at the correct location inside the bucket as he fills it rather than, as others do, first just dealing the numbers into buckets on stack B, and only after all the buckets are filled, inserting their contents into the correct locations on stack A.

(According to Wikipedia, the classic bucket orders an array as follows: It creates an auxiliary array, "scatters" the items into buckets in this auxiliary array, then sorts them in their buckets, and only then "gathers" the items, returning the buckets in the right order to the original array. Neither Jamie nor Fred do exactly this, but something more like what Fred does is described as a common [optimization](https://en.wikipedia.org/wiki/Bucket_sort#Optimizations). Note, however, that what is being optimized, in the sense of the Wikipedia article, is time complexity, not amount of push-swap instructions.)

For stacks of 100 numbers, Jamie uses 5 buckets, and for stacks of 500 numbers, 11 buckets. Thus, for 100, he pushes the lowest 20 numbers to stack B first, then the next lowest, and so on till everything is sorted on B. It's apparent from his description and example that he does this in ascending order. Descending would be better as it would allow the numbers to fall naturally into ascending order when pushed back to A. (It's strange that Jamie misses this trick, especially as he used it on his initial, naive attempt, using selection sort.) Anyway, at each iteration, he pushes the number in the current range that takes least rotations to reach the top of stack A, after first rotating B so that the number being pushed will land in the desired place. Finally he pushes them all back, rotating B between pushes to account for that fact that it's in ascending order. (If B was in descending order, we could just push everything back to A, saving ourselves 99 rotations.)

As with Julien's method, all but trivial sorting happens in one direction. For Jamie, the sorting happens from A to B; for Julien, from B to A. We'll see in [Results](#c-results) that a conventional bucket sort first, to triage the numbers onto B before sorting them fully onto A, would be better, as would fewer buckets. In fact, the best number of buckets with this technique is one! If we just push everything from A to B, then insertion sort back, Jamie's technique reduces to Julien's.

YYBer, like many push-swappers, including ourselves, follows the convenience of ranking the numbers, thus "-3 -100 0 360 108" becomes "2 1 3 5 4". This can simplify calculations, and the instructions needed to sort them are the same. (Just keep in mind that she uses the word index to mean rank in this sense.) For stack size 100, she pushes them from A to B as 8 buckets of 12 numbers each. To sort 500 numbers, she uses 12 buckets of 40 numbers each. In either case, she pushes whatever is left to the top of B as one last bucket till A is empty.

How YY decides on the next number in the current range to push to B is not altogether clear. Her method has much in common with Luca Fischer's, who uses the same procedure as Jamie, so it's likely she does the same.

It's no accident that both bucket counts are even. This lets her deal with two buckets at once. For example, with 100 numbers, she puts anything from 13-24 on top of B with a simple `pb`, and anything from 1-12 on the bottom with `pb rb`. When this is done, she moves on to the next two buckets, putting 37-48 on top of B and 25-36 at the bottom, and so on.

When the remaining numbers have been pushed to B, so that A is empty, YY uses a cost checking procedure to push everything back. First the maximum number is pushed to A. Then the new number at the top of B is pushed and rotated to the bottom: `pa ra`.

Then, at each iteration, she checks if the value at the top of B is one less than the current top of A. If so, she places it there with `pa`. Otherwise, if the value at the top of B is greater than the one at the bottom of A, she runs `pa ra` to push it to A and rotate it to the bottom. But if the number at the top of B can be placed neither at the top nor at the bottom of A, she finds the maximum remaining number on B and also locates the number greater than the one at the bottom of A which is least rotations (or reverse rotations) away from the top. She checks which of these two takes the smallest amount of rotations (or reverse rotations) to reach the top of B, then rotates it there and pushes it. If the number pushed was the maximum, it will be correctly positioned as one less than the previous top of A. If not, she rotates it to the bottom of A.

Every time a number is pushed to the top of A and left there, a check is made to see if the number at the bottom can now be reverse rotated into its final position on top. This is repeated for as long as possible.

In this way, she sorts everything back to A.

Luca Fischer follows much the same approach. His and YY's articles are best read together. Obscure points in one may be clarified by the other. He also pushes pairs of buckets at a time to B. Both he and YY use the word "ratio" for the size of each bucket (of the pair) or, to put it another way, for the midpoint between the buckets of the first pair. (This value is presumably calculated as a ratio of the original stack size, although YY uses different ratios for 100 and 500.) YY and Luca each have a video showing how this puts the two buckets with the smallest values in the middle of B, and how, at each stage, the next pair of buckets arrive with the smaller values at the bottom of B, and the larger values at the top.

Luca doesn't say that he pushes the first value to go back to A indiscriminately from the top, but we've assumed so by analogy with YY's indiscriminate first push after the maximum of all the numbers has gone back. His description doesn't make sense without this assumtion (as otherwise no number on B would ever be greater than the number at the bottom of A), and his video clearly shows values less than the maximum being stored temporarily on the bottom of A.

There are just a couple of differences: in Luca's case, for stack size 100, each bucket of the pair holds 14 numbers (versus 12 for YY); and, whereas YY empties stack A, Luca only pushes till the largest three numbers are left on A, and sorts them there. On the way back, Luca checks what's at the top of B. If its correct final position is on top of the current top of A, he pushes it there. If it's greater than the number at the bottom of A, he pushes and rotates to the bottom of A. (He characterizes this as "temporarily storing values in the bottom portion of stack A.") If neither condition is met, he pushes the current maximum from B.

We can now think of Fred's style of bucket sort as a special case of the method used by YY and Luca. Rather than a division into three equal parts, we can see it as sorting into a pair of buckets, each of size 33 (or 34, depending on one's rounding strategy), then pushing the remainder to the top of B. In Fred's case, the first pair of buckets is too big, relative to the stack size, to allow for more pairs (or indeed single buckets, given his rule to stop when there are three numbers left on A). Hence he only performs one iteration of dealing out a pair of buckets before pushing the remainder.

YY and Luca just continue the process, placing buckets in layers as they wind out from the center of B. This is great positioning to take advantage of the circularity of the stack, since the bottom is closer to the top, in terms of rotations, than the center.

Leo Fu takes a quite different approach. He uses base 2 [radix sort](https://en.wikipedia.org/wiki/Radix_sort), padding smaller numbers with leading zeros as needed. For each bit, starting with the least significant (rightmost), he checks the number at the top of stack A. If the relevant bit is 0, he pushes the top number from A to B with `pb`; otherwise he applies `ra` to rotate it out of the way to the bottom of A. In this way, he goes through all the numbers in A. Then he pushes back everything from B with `pa`, and procedes in this way through all the bits. He says it didn't get him the highest score. Presumably the cost of having to push all those numbers back and forth on multiple passes was too much.

No doubt [btilly's idea](https://stackoverflow.com/questions/75100698/push-swap-what-is-the-most-efficient-way-to-sort-given-values-using-a-limited-s) of sharing out runs between the two stacks and merging them back and forth would suffer from the same issue: too much pushing.

A few others have made their solutions public on GitHub, such as [Adrian Roque](https://github.com/AdrianWR/push_swap), who credits [Anya Schukin](https://github.com/anyaschukin/Push_Swap) with this idea. Anya describes how, for stacks of no more than 100 numbers, she uses 2 buckets. Everything less than the median is pushed to B. Then, while the larger half is still on A, she sorts the smaller half back. At each iteration, she finds the minumum and the maximum number remaining on B and pushes whichever is cheaper (takes fewest rotations). For stacks larger than 100 and no larger than 500, she says, "I executed the same process but divided stack `a` by quarters instead of median."

### b. Grading systems

It seems the push-swap rules have varied slightly over time and space. We had two write a checker and a push-swap program, as did Ali Yigit Ogun at 42-Heilbronn; others, such as YYBer, only had to write push-swap while a checker was provided. The projects I've seen discussed online were written in C (although the articles focus on strategy rather than implementation). Ours had to be in Go. But the basic idea of sorting numbers using two stacks is the same, as are the permitted operations.

Different scoring systems are used by the various schools, which can sometimes offer clues about the performance of folks' solutions when they don't go into specifics. At Ecole 42, Lyon, in 2021, [Leo Fu](https://medium.com/nerd-for-tech/push-swap-tutorial-fa746e6aba1e) passed by sorting 100 numbers in "about 1084 instructions". He quotes a scoring system in which top marks are gained by sorting 100 numbers in less than 700 instructions, and 500 in less than 5500.

Ali Yigit Ogun says he scored 125/125. As for what this means, he links to a PDF of his school's instructions, but all they say on scoring is that if your list of instructions is "too big" it will fail. (It refers to a "maximum number tolerated" without specifying.) Similarly, at 42 Silicon Valley in 2019, JD needed to pass some requirement for 100 and 500, although he doesn't say how many instructions he was allowed. Of course, a dedicated push-swappist could persuse the commit histories of these various schools' public repos.

[Dan Sylvain](https://medium.com/@dansylvain84/my-implementation-of-the-42-push-swap-project-2706fd8c2e9f) presents a similar grading system to LF, but talks about a "base project" having to be perfect before you can conplete the "bonus part", as if either the 100 or the 500 test are bonuses.

At any rate, by December 2023, at 01 Founders in London, we'd get an unspecified bonus if we could sort 100 mumbers in less than 700 of the specified operations. No mention is made of 500 numbers in our [audit questions](https://github.com/01-edu/public/tree/master/subjects/push-swap/audit).

### c. Results

In what follows, we'll compare the performance of various algorithms on 10,000 trials at sorting stacks of 100 numbers. Since doing these tests, we've realized that performance on 100 numbers may not be indicative of how they fare on smaller stacks, so please bear in mind that this is not an absolute verdict.

Our implementation of Fred Orion's algorithm took an average of 555 instructions to sort 100 numbers, with a standard deviation of 25, and never more than 699. The worse cases were in the low 600s, with 498 (4.98%) over 600. Of all the ways we've tried so far, this is the winner, and it sounds like Fred may have achieved some further optimization that wasn't revealed in his summary, unless I've overlooked it. He says he sorted 100 with a mean of 510 instructions. But I don't know how many tests he did. To sort 500 numbers, he reports a mean of 3750 instructions; our version scored 4216 on 100 trials, with a standard deviation of 121.

We tried varying the number of buckets used in this technique: four buckets performed somewhat worse at 569 instructions, and two buckets even worse at 573.

Ali Yigit Ogun's "Turk" algorithm achieved a mean of 561, with a standard deviation of 23. Again, the the worst cases were in the low 600s, and 549 (5.49%) of them were over 600. (Without Ali's cost calculation, the mean was 1387.)

Our initial checks to see if the stack can be simply swapped and rotated into order make no appreciable difference to the score when sorting 100 numbers on so many trials.

Dan Sylvain did pretty well with his strategy of sorting all but the longest increasing subsequence into two buckets on B, then insertion sorting back to A with a cost check. Our implementation scored a mean of 566.

Inspired by Dan, we also tried leaving the longest run (i.e. the longest sequence of numbers adjacent in the initial stack, such that their ranks are consecutive integers) on A and pushing everything else into two buckets on B, after which we insertion sorted them back with a cost check. This resulted in a mean of 577 instructions. But then, the length of the longest run, for 100 uniformly distributed random numbers, is mostly 1 or 2.

Julien Caucheteux's approach of pushing everything, then insertion sorting with cost checking like Ali on the way back took 584 instructions, in our interpretation.

YYBer reports first scoring 750 for 100 numbers by pushing the smallest half to A, then the smallest half of the rest, and so on till A is empty, and then selection sorting them back to A. After this attempt, she switched to a new strategy: push into 8 equal-sized buckets on B in the case of 100 numbers, or 12 in the case of 500, then push whatever is left till A is empty. Then selection sort back to A. This, she says, required less than 700 instructions for 100 numbers, and about 6500 for 500. As a final optimization, she replaced selection sort with the cost-checking procedure described above. With this, she reports a mean score for 500 numbers to "5300-5400 steps".

Our implementation of YY's algorithm scored 631 on 100 numbers (with just 9 over 699), and 4814 on 500.

Anya Schukin's algorithm, in our implementation, sorted 100 numbers in 811 instructions.

Our version of Jamie Dawson's algorithm took 867 instructions. It can be cut to 768 by the simple expedient of switching the intial sort onto B from ascending to descending. Shared rotations bring it down further to 713. A natural question is: what would happen if we just triage the numbers into buckets on B (as in a traditional bucket sort), then only apply the insertion sort as we push them back to A? It turns out we only need 674 instructions. How about adjusting the number of buckets?

5: 674  
4: 643  
3: 632  
2: 595  
1: 584

So, here, the buckets are actually making it worse! Notice that, with one bucket, i.e. just pushing everything to B and insertion sorting back, this incremental optimization of Jamie's algorithm becomes exactly Julien's method.

How is Fred doing so well with 3 buckets, then? He doesn't search for numbers in a current range and rotate them into position to push. Instead, he just takes whatever is on top of A and either rotates it out of the way to the bottom of A (to be pushed later), or simply pushes to B, or pushes to B and rotates: just one or two moves to place each item.

Leo Fu reports "about 1084" instructions for 100 numbers, and "about 6756" for 500, then remarks that he actually always got exactly 6756, no matter how many times he tested it on different random numbers, and poses the question: why? We'll return to this [shortly](#c-why-does-leo-fus-radix-sort-always-take-the-same-amount-of-instructions-for-a-given-stack-size).

In the table below, the number of buckets is shown in paretheses where we tried varying it. Fred(3) is Fred's original algorithm, and Dan(2) is Dan's original.

Fred(3) 555  
Ali 561  
Dan(2) 566  
Fred(4) 569  
Fred(2) 573  
Longest Run 577  
Dan(3) 578  
Julien 584  
YYBer 631  
Luca 655  
Jamie (+desc, +shared, +triage) 674  
Jamie (+desc, +shared) 713  
Jamie (+desc) 768  
Anya 811  
Jamie 867  
Leo 1084  
Ali(-cost) 1387

As for 500 numbers, Jamie says he used the same logic as for 100, "But instead of splitting it into 5 chunks, \[I\] just split it into 11 chunks. Why 11? 11 chunks are what I decided to use after running several tests on it. The range of action points I got was way less than other numbers I tested it on."

We tested a version of his algorithm--optimized with shared rotations and descending order on B, and making the initial bucket sort a simple triage rather than combining it with instertion sort--for different amounts of buckets on 500 numbers. Each test was performed 100 times. The sweet spot seems to be around 4 buckets.

1: 5208  
2: 4677  
3: 4532  
4: 4505  
5: 4591  
6: 4735  
7: 4890  
8: 4997  
9: 5166  
10: 5326  
11: 5543  
12: 5726  
...  
20: 7280

turk: 5105  
orion: 4216

On the other hand, Jamie's algorithm with full insertion sort into the buckets on B (more like what he actuall used, just optimized with descending order on B and shared rotations) does seem to be optimal around the low teens, so his choice of 11 buckets is not unreasonable.

1: 32199  
2: 19493  
3: 10804  
4: 11552  
5: 9806  
6: 7963  
7: 8677  
8: 7345  
9: 7109  
10: 6978  
11: 6821  
12: 6743  
13: 6684  
14: 6701  
15: 6776  
14: 6701  
15: 6776  
16: 6834  
17: 6914  
18: 7057  
19: 7114  
20: 7205  
..  
30: 8555

But to really get a sense of how these algorithms compare, we'd also need to test them over a range of numbers, including smaller stacks. As mentioned, Ali performed better for stacks of less than 93 numbers. What other surprises are out there?

## 4. Structure and strategy

You'll find the source code in several folders: `push-swap`, for the program that generates instructions to sort a given list of numbers, and `checker` for the program to check whether a given sequence of instructions sorts a given list of numbers, both in `cmd`.

The `ps` package is for functions and structs these programs share.

`cmd` also contains a folder called `explorer`, which is a playground for exploring new ideas. Here we looked at the advantages of BFS. The JSON files contain lists of the shortest push-free solutions for stack sizes 5, 6, and 7. `explorer` is also the messy attic where we stow the remains of old experiments, including our implementations of the various algorithms, and some tests to compare their performance.

Like most push-swappers, we dealt with initial stacks of less than six numbers as special cases. This was partly because they lend themselves to optimizations that would be prohibitively complex or time-consuming for longer lists, and partly so we could treat these smaller problems as warm up exercises. The six permutations of three elements are easily checked by hand. With a ranking function to simplify the task, it didn't take much longer to find optimal solutions for the 24 permutations of four elements and hardcode them.

At 120 permutation, the five-number challenge seemed like a good place to start automating, particularly as Jamie Dawson uses it as an example to develop intuition for how one might proceed to the general case. Indeed, we followed his method.

For 6 through 92, we used Ali Yigit Ogun's "Turk" algorithm, which resembles Jamie Dawson's for 5 numbers, except for Ali's more general cost check. The first two numbers on the stack are pushed indiscriminately from A to B. However, Ali sorts the rest of the numbers from A to B so that they land in descending order. This means it takes fewer instructions to push them back into ascending order on A. Before deciding which number to push from A, he checks each one to find the cheapest.

The cheapest number to move is the one that miminizes

```
A + B - C
```

where `A` is number of rotations needed to bring this number to the top of stack A, `B` is the number of rotations needed to put its target number at the top of stack B, and `C` is any combined rotations. At least in our program, this calculation makes a big saving on instructions for sorting 100 numbers, as detailed in the previous section.

When pushing the cheapest number from A to B, its target is the biggest smaller number or, if there is no smaller number, the maximum of B. When pushing back, the target is the smallest bigger number or, if there is no bigger number, the minimum. For more detail, examples, and illustrations, see Ali's article and the [video by Thuy Quematon](https://www.youtube.com/watch?v=wRvipSG4Mmk).

For the general case with more than 92 numbers, we use Fred Orion's initial triage of the lower two thirds into two buckets on B, the greater numbers on top, the lesser at the bottom, then push all but three of the rest to the top of B, sort those three in place, and then insertion sort everything back with the same cost check that Ali uses.

As mentioned, we check for some cases where there is a shorter push-free sort. For stack sizes 5 though 8, we find the shortest push-free solution with BFS and return this if it's shorter than our current result. For larger stacks, we make do with a few simple checks to see if the stack is already sorted or can be just rotated into the correct order, or sorted with a swap, or rotations and a swap, or rotations and a swap and more rotations.

## 5. Mathematical curios

### a. Swaps and rotations are enough to sort

Any permutation of `n` elements, `{1, 2, 3, ..., n}`, can be expressed as a sequence of swaps and rotations, so, if we didn't care about how many instructions it takes to sort a stack, we could just use these two operations.

In the language of group theory, an `n`-cycle, such as `(1 2 3 ... n)` (i.e. the permutation that sends `1` to where `2` was, and `2` to where `3` was, and ..., and `n` to where `1` was), and a transposition of elements that are adjacent in this cycle, such as `(1 2)`, together generate the whole symmetric group on `n` elements. These statements are equivalent because the inverse of `(1 2 3 ... n)` (a reverse rotation) is a composition of `n - 1` instances of `(1 2 3 ... n)`, while `(1 2)` is its own inverse.

To see that they generate the whole symmetric group, we can use the fact every permutation can be expressed as a composition of transpositions. A transposition of neighboring elements can be achieved by rotating them into the top two spots on the stack, then performing the swap operation, then rotating them back to their original positions. To transpose elements that are a positive number `k` steps apart, place one of them at the top of the stack so that the other is no more than half `n` steps below it. (This is the maximum distance apart around the circlular stack that any pair of elements can be.) Then perform a swap, then `k - 1` rotations, each followed by a swap, and then `k - 1` reverse rotations, each also followed by a swap. The result will be that the elements have changed places.

Since any permutation can be sorted on one stack, it follows that it can be sorted when two are available, an unspoken assumption of this project!

### b. Antipodeal elements: an optimization for stacks of even size

Due to the circular nature of the stacks, the cheapest numbers to push will tend to be those near the top or the bottom. In other words, a number is actually furthest from the top when it's near the middle of the stack. If the stack has `n` elements indexed from `0` at the top, then those whose index is less than or equal to the floor of `n/2` will reach the top sooner when rotated upwards, while, for those whose index is greater than the floor of `n/2`, the top is reached soonest when they're rotated downwards. This means that, when `n` is an even number, there will be a middle element (we could call it opposite or antipodeal) which takes either `n/2` upwards or `n/2` downwards rotations to reach the top. (Think how a clock, where even-numbered 12 is also 0, has such an opposite element: 6.)

One consequence of this is that, if we need to rotate one stack, say, `r` times upwards, and the other stack has `2 * r` elements, then if we need to rotate the second stack `r` times, we can choose to rotate it upwards too, to take advantage of the combined rotation operation.

While, interesting to think about, and certainly an improvement, this didn't make a noticeable difference to the mean number of instructions needed for large stacks.

### c. Why Leo Fu always got the same amount of instructions for a given stack size

As mentioned above, Leo Fu reports that his implementation of base 2 radix sort took "about 1084" instructions for 100 numbers, and "about 6756" for 500. He imediately corrects himself, saying that he actually always got exactly 6756.

To see how this can be, and why he must also have always got exactly 1084 for 100 numbers, note first that Leo takes the convenient step of converting the original values to their rank: 0, 1, 2, 3, 4, ..., 98, 99.

Ceil(log2(99)) = 7, so it takes 7 bits to express these numbers. Picture them listed in binary form, padding the smaller ones with leading zeros:

0000000  
0000001  
0000010  
0000011  
0000100  
...  
1100010  
1100011

In sorting, there's one pass (iteration) for every bit, starting with the least significant, i.e. at the right. At every pass, all the numbers with a 0 at that bit will be pushed to B with `pb` and eventually back with `pa`, while all the numbers with a 1 at that bit will be rotated out of the way with `ra`.

If the size of the stack was a power of 2, there would be an equal number of ones and zeros in the list of numbers (in binary form) to be sorted. But the bits of 99 are 1100011, so not every 7-bit sequenece of 0s and 1s is represented among the numbers to be sorted. Since it's the greatest 28 numbers that are missing between 99 and 127 (one less than the next power of 2), i.e. from the full total of all possible 7-bit sequences, it's the 0s that will be overrepresented in the total, shifting the balance in favor of pushes.

We can calculate the exact amount of instructions as the number of ones in the list, plus twice the number of zeros, `ones + 2 * zeros`, or alternatively as `bits * n + zeros` (with n = 100 in this case; Leo gives the latter formula in a comment on his article), since every number requires at least one instruction per bit (whether a rotation or a push), and then those with zero at that bit need one extra instruction to push them back. Sure enough, both formulas give 1084 for n = 100.

Similarly, ceil(log2(499)) = 9, so it takes 9 bits to represent the numbers from 0 to 499. It's the highest 12 numbers that are missing from the full total of 2^9 = 512 possible 9-bit sequences, so, as always when the stack size is not a power of 2, the zeros will be overrepresented, and there will be more pushes than rotations.

At this point, though, we have a a bit of a mystery. By my calculation, `9 * 500 + zeros` and `ones + 2 * zeros` both give 6784 rather than Leo's 6756. I'm quite puzzled by this. In his article, Leo actually shows a screenshot of a checker program showing the result 6756 for every random sequence of 500 numbers it tested, so it's not typo. I wonder if he introduced some optimization for stack sizes greater than 100 that I overlooked or that he didn't mention. I've written to him, asking if he has any insight. I'll update this if I learn more. Here's how I'm counting:

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println(countTotal(500))
}

func countZeros(n, bits int) int {
	zeros := 0
	for i := 0; i < bits; i++ {
		if (1<<i)&n == 0 {
			zeros++
		}
	}
	return zeros
}

func countTotal(n int) int {
	if n == 0 {
		return 0
	}
	zeros := 0
	bits := 1
	for m := 1; m<<bits < n; bits++ {
	}
	for i := 0; i < n; i++ {
		newZeros := countZeros(i, bits)
		zeros += newZeros
	}
	// Equivalently, since ones == bits*n - zeros, we could return
	// ones + 2*zeros.
	return bits*n + zeros
}

```

## 6. Detour: bitmasks

The code in the previous section uses a bitmask to identify individual bits of a number. Here we illustrate the concept with another example.

In the `getInstructions` function of the `checker` program (located in `get-instructions.go`), we wanted to move the cursor up a line to eliminate the blank line that results when the user indicates that they've finished typing instructions by pressing enter on a line with no instructions.

However, we can't unconditionally move up a line because, when the intructions are piped to the program, there is no blank line. Hence we check whether the input is from a character device (i.e. the terminal) before moving up a line.

To this end, the `main` function in `checker` passes a boolean, `(fi.Mode() & os.ModeCharDevice) != 0`, to `getInstructions`.

The `&` is the bitwise AND operator. It returns a number which, when expressed in binary, has a 1 only at positions where both of its operands have 1, and 0 everywhere else.

`fi.Mode()` returns a number, the 0s and 1s of whose binary representation indicate file mode and permissions. In particular, its 14th bit shows whether the input is from the command line.

`os.ModeCharDevice` is a constant, equal to the 14th power of 2, namely 8192. This will be the bitmask we use to query that 14th bit.

Thus, `fi.Mode() & os.ModeCharDevice` will only be zero when the 14th bit of `fi.Mode` is 0. (Otherwise it will be 8192.)
