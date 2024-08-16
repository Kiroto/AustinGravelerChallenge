import random
import time
from itertools import repeat
import cProfile

def myFunction(atts):
    sessions = 0
    maxOnes = 0
    numbers = [0, 0, 0, 0]
    rollsPerSession = 231
    bitshifts = 16

    startTime = time.perf_counter()

    while numbers[0] < 177 and sessions < atts:
        rolls = 0
        while True:
            randomlyGeneratedNumber = random.getrandbits(32)
            for shift in range(bitshifts):
                roll = (randomlyGeneratedNumber >> (shift << 1)) & 0b11
                numbers[roll] += 1
                rolls += 1
                # This feels like it's holding me back
                if (rolls > rollsPerSession):
                    break
            else:
                continue
            # I heavily dislike this
            break

        sessions += 4
        for i in range(len(numbers)):
            if numbers[i] > maxOnes:
                maxOnes = numbers[i]
            numbers[i] = 0

    endTime = time.perf_counter()

    
    print("Highest Ones Roll: ", maxOnes)
    print("Number of Roll Sessions: ", sessions)
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

ATTEMPTS = 1_000_000

cProfile.run("myFunction(ATTEMPTS)")


# cProfile.run("theirFunction(ATTEMPTS)")