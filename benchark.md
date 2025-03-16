# Mongo DB Stress Test

# Insert 850000 documents

## Go Insert

## Insert Many

2452 ms (2.45 s) -- no cpu or memory limit
9300 ms (9.3 s) -- 0.25 cpu, 512Mi memory limit
5054 ms (5.05 s) -- 0.5 cpu, 512Mi memory limit
2961 ms (2.96 s) -- 1 cpu, 512Mi memory limit
2578 ms (2.58 s) -- 2 cpu, 512Mi memory limit

### With Batches

100000= 939 ms (0.939 s)
10000 = 902 ms (0.902 s)
1000 = 1066 ms (1.066 s)

## Insert One (With Loop)

50918 ms (50.9 s) -- no cpu or memory limit (cpu constant 0.5 | mem = 500 mb)

## Find All

2831 ms (2.83 s) -- no cpu or memory limit

## Find By Name

10 ms (0.01 s) -- no cpu or memory limit

## Find by Id

9 ms (0.009 s) -- no cpu or memory limit

## Node JS Insert

## Insert Many

14169 ms (14.1 s) -- no cpu or memory limit

### With Batches

100000= 12087 ms (12.087 s)
10000 = 12388 ms (12.388 s)
1000 = 22093 ms (22.093 s)

## Insert One (With Loop)

165409 ms (165.4 s) -- no cpu or memory limit (cpu constant 0.5 | mem = 500 mb)

## Find All

7521ms (7.521s) -- no cpu or memory limit

## Find By Name

13.148ms (0.013148s) -- no cpu or memory limit

## Find by Id

12.217ms (0.012217s) -- no cpu or memory limit

## Mongo Import

6242 ms (6.24 s) -- no cpu or memory limit

## MongoDump

1399 ms (1.39 s) -- no cpu or memory limit
44405 ms (44.4 s) -- 0.25 cpu, 512Mi memory limit
20314 ms (20.3 s) -- 0.5 cpu, 512Mi memory limit
10248 ms (10.2 s) -- 1 cpu, 512Mi memory limit
1926 ms (1.92 s) -- 2 cpu, 512Mi memory limit
