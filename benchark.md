# Mongo DB Stress Test

# Insert 850000 documents

## Go Insert

## Insert Many

2452 ms (2.45 s) -- no cpu or memory limit
9300 ms (9.3 s) -- 0.25 cpu, 512Mi memory limit
5054 ms (5.05 s) -- 0.5 cpu, 512Mi memory limit
2961 ms (2.96 s) -- 1 cpu, 512Mi memory limit
2578 ms (2.58 s) -- 2 cpu, 512Mi memory limit

## Insert One (With Loop)

50918 ms (50.9 s) -- no cpu or memory limit (cpu constant 0.5 | mem = 500 mb)

## Node JS Insert

## Insert Many

14169 ms (14.1 s) -- no cpu or memory limit

## Insert One (With Loop)

165409 ms (165.4 s) -- no cpu or memory limit (cpu constant 0.5 | mem = 500 mb)

## Mongo Import

6242 ms (6.24 s) -- no cpu or memory limit

## MongoDump

1399 ms (1.39 s) -- no cpu or memory limit
44405 ms (44.4 s) -- 0.25 cpu, 512Mi memory limit
20314 ms (20.3 s) -- 0.5 cpu, 512Mi memory limit
10248 ms (10.2 s) -- 1 cpu, 512Mi memory limit
1926 ms (1.92 s) -- 2 cpu, 512Mi memory limit
