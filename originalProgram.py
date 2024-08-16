import random
import time
from itertools import repeat
import cProfile

# This is where the results of the rolls are stored.
# It is reset on each iteration (don't need to declare it here, even)
# If the random "item" selected is "3", then the third index is incremented by 1. Similar with the other 3 indices.
# In this program, we consider a paralysis proc to be if the selected item is "1" (the first index).

# The amount of roll sessions that have been performed until now.

# The greatest amount of paralysis procs in a single session.


# Do roll sessions
# Until we hit 177 procs
# or until we hit a billion rolls.
def myFunction(atts):
    rolls = 0
    maxOnes = 0
    numbers = [0, 0, 0, 0]

    bitshifts = 16

    startTime = time.perf_counter()

    while numbers[0] < 177 and rolls < atts:
        for _ in repeat(None, 231):
            roll = random.randint(0, 3)
            numbers[roll] += 1
        rolls += 4
        for i in range(len(numbers)):
            if numbers[i] > maxOnes:
                maxOnes = numbers[i]
            numbers[i] = 0

    endTime = time.perf_counter()

    
    print("Highest Ones Roll: ", maxOnes)
    print("Number of Roll Sessions: ", rolls)
    print(f"Time elapsed: {(endTime-startTime):0.4f} seconds")


def theirFunction(atts):
    rolls = 0
    maxOnes = 0
    items = [1, 2, 3, 4]

    startTime = time.perf_counter()

    numbers = [0, 0, 0, 0]

    while numbers[0] < 177 and rolls < atts:
        numbers = [0, 0, 0, 0]
        for i in repeat(None, 231):
            roll = random.choice(items)
            numbers[roll - 1] = numbers[roll - 1] + 1
        rolls = rolls + 1
        if numbers[0] > maxOnes:
            maxOnes = numbers[0]

    endTime = time.perf_counter()

    print("Highest Ones Roll: ", maxOnes)
    print("Number of Roll Sessions: ", rolls)
    print(f"Time elapsed: {(endTime-startTime):0.4f} seconds")

ATTEMPTS = 1_000_0

cProfile.run("myFunction(ATTEMPTS)")


# cProfile.run("theirFunction(ATTEMPTS)")