#!/bin/bash
# L07 Acceptance Tests

SRCDIR=$HOME/dev/go/src/p436/L07

#Standard Test
printf "Standard Test - Testing Input: [apple, 10, banana, 25, carrot, 100]\n\n"
printf "Expected Output:\n"
printf "apple\nbanana\ncarrot\n\n"
printf "Sum: 135\n\n"

printf "Actual Output:\n"
L07 apple 10 banana 25 carrot 100
printf "________________________________________________________________\n\n"

#Decimal Test
printf "Decimal Test - Testing Input: [banana, 10.5, apple, 23, carrot, daisy, 34.6, 45.3, 67.9]\n\n"
printf "Expected Output:\n"
printf "apple\nbanana\ncarrot\ndaisy\n\n"
printf "Sum: 181.3\n\n"

printf "Actual Output:\n"
L07 banana 10.5 apple 23 carrot daisy 34.6 45.3 67.9
printf "________________________________________________________________\n\n"

#Only Strings Test
printf "Only Strings Test - Testing Input: [banana, apple, carrot, daisy]\n\n"
printf "Expected Output:\n"
printf "apple\nbanana\ncarrot\ndaisy\n\n"
printf "Sum: 0\n\n"

printf "Actual Output:\n"
L07 banana apple carrot daisy
printf "________________________________________________________________\n\n"

#Only Numbers Test
printf "Only Numbers Test - Testing Input: [10, 23, 34.4, 87.9]\n\n"
printf "Expected Output:\n"
printf "\n"
printf "Sum: 155.3\n\n"

printf "Actual Output:\n"
L07 10 23 34.4 87.9
printf "________________________________________________________________\n\n"

#Negative Numbers Test
printf "Negative Numbers Test - Testing Input: [10, -23, 34.4, -87.9]\n\n"
printf "Expected Output:\n"
printf "\n"
printf "Sum: -66.5\n\n"

printf "Actual Output:\n"
L07 10 -23 34.4 -87.9
printf "________________________________________________________________\n\n"

#Empty Test
printf "Empty Test - Testing Input: []\n\n"
printf "Expected Output:\n"
printf "\n"
printf "Sum: 0\n\n"

printf "Actual Output:\n"
L07
printf "________________________________________________________________\n\n"
