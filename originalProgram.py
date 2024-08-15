import random
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
