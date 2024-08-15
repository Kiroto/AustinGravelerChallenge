# Fast graveler lockpick checker

Note: As of the time of writing this, I'm very well aware my go code is very messy and ugly to look at, but it is performant.

## TL;DR
Python slow.
Golang Fast.
Getting more value out of your generated numbers.

### To execute:
> Note: Remember that executing files you don't know is dangerous. Make sure you trust whatever you're executing, or otherwise execute programs in a volatile environment disconnected from the internet, or otherwise stuff that you don't care losing.

#### Build it yourself with Go
1. Set up **Go** in your computer ([here's an official link](https://go.dev/doc/install))
2. Download the source code.
3. Build `fast.go` (via running `go build fast.go` on a terminal whilst the terminal is open at that folder)
4. Run it, just execute `./fast.exe` on the terminal.

#### Trust a random guy in the internet and download it and run it yourself.

You can find the executable on the Releases tab. Otherwise [this](/releases/latest) link should take you there.

## About the original Python Code

Obtained from [Austin's github repo](https://github.com/arhourigan/graveler/blob/main/graveler.py). (+ some comments and formatting)


<details>
<summary>Original Code</summary>

```py
import random
import math
from itertools import repeat

# Indices 1 to 4, out of which one (pseudo random) result will be selected.
# This is the source of the 1 in 4 chance.
items = [1,2,3,4]

# This is where the results of the rolls are stored.
# It is reset on each iteration (don't need to declare it here, even)
# If the random "item" selected is "3", then the third index is incremented by 1. Similar with the other 3 indices.
# In this program, we consider a paralysis proc to be if the selected item is "1" (the first index).
numbers = [0,0,0,0]

# The amount of roll sessions that have been performed until now.
rolls = 0

# The greatest amount of paralysis procs in a single session.
maxOnes = 0

# Do roll sessions
# Until we hit 177 procs
# or until we hit a billion rolls.
while numbers[0] < 177 and rolls < 1_000_000_000:
    numbers = [0,0,0,0]
    for i in repeat(None, 231):
        roll = random.choice(items)
        numbers[roll-1] = numbers[roll-1] + 1
    rolls = rolls + 1
    if numbers[0] > maxOnes:
        maxOnes = numbers[0]

print("Highest Ones Roll:",maxOnes)

print("Number of Roll Sessions: ",rolls)
```
</details>


<details>
<summary>Comments by me + formatting</summary>

```py
import random
# Not used
import math
from itertools import repeat

# Indices 1 to 4, out of which one (pseudo random) result will be selected.
# This is the source of the 1 in 4 chance.
items = [1, 2, 3, 4]

# This is where the results of the rolls are stored.
# It is reset on each iteration (don't need to declare it here, even)
# If the random "item" selected is "3", then the third index is incremented by 1. Similar with the other 3 indices.
# In this program, we consider a paralysis proc to be if the selected item is "1" (the first index).
numbers = [0, 0, 0, 0]

# The amount of roll sessions that have been performed until now.
rolls = 0

# The greatest amount of paralysis procs in a single session.
maxOnes = 0

# Do roll sessions
# Until we hit 177 procs
# or until we hit a billion rolls.
while numbers[0] < 177 and rolls < 1_000_000_000:
    numbers = [0, 0, 0, 0]
    for i in repeat(None, 231):
        roll = random.choice(items)
        numbers[roll - 1] = numbers[roll - 1] + 1
    rolls = rolls + 1
    if numbers[0] > maxOnes:
        maxOnes = numbers[0]

print("Highest Ones Roll: ", maxOnes)

print("Number of Roll Sessions: ", rolls)
```
</details>

### Why is it oddly cool?
1. It only uses one required if statement per roll session.
2. You're using `itertool`'s `repeat` when you could have as well used `range` for the same effect, yet `repeat` is apparently faster.

### Why is it slow and/or ugly? + nitpicks
1. It's written in Python and most probably ran interpreted (instead of compiled), which isn't particularily performant.
2. Three fourths of the random numbers generated are ignored.
3. Don't need the math library
4. You're disposing of and creating a new value for `numbers` every roll session. (TO BE TIMED.)

### How can we make it faster?
1. Write it in a more performant language.
2. Use all the random numbers generated, as possible.

### How did I do it?
1. Used Golang
2. Saved all 4 "possible outcomes" of the random number generation as their separate battle session.
3. Avoided if statements where possible (although I did add one that doesn't affect perfonmance much)
4. Didn't use many print statements (those are slow)