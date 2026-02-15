# PUSH SWAP

This is Golang implementation of PushSwap project (42 Subject).

## Basic description

Given two stacks: A and B. Both stacks can contain `only` unique numbers. In the beginning all numbers are stored in stack A. Order of numbers is random, can be sorted or unsorted.

Student should implement two programs:
1) push_swap: Program that sorts numbers in stack A using given commands and stack B as buffer.
2) checker: Program that checks whether after executing given commands stack is sorted or not.


## Usage

Firstly, build push_swap program.

```bash
$> cd push_swap && go build .
```

Build checker program.

```bash
$> cd ../checker && go build . && mv checker ../push_swap && cd ../push_swap
```

Execution

```bash
$> ARG="4 67 3 89"; ./push_swap $ARG | ./checker $ARG
OK
```

## Project Structure

There are 6 different directories in current project:
* checker - checker program
* node - list data structure
* push_swap - push_swap program
* pushswap - main data structure that contains two stacks, command methods (swapA, swapB, ...) etc.
* stack - stack data structure with methods as: pushBack, pushFront, pop, etc.
* tools - helper functions to work with numbers and arrays.
