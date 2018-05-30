#!/bin/bash
# L11 Acceptance Tests

SRCDIR=$HOME/dev/go/src/p436/L11

#Standard Test
printf "Standard Test - Testing Input: [apple, 10, banana, 25, carrot, 100]\n\n"
printf "Expected Output:\n"
printf "Start Time:  Mon Mar 26 01:47:14 EDT 2018
Thread: carrot - Count 1
Thread: apple - Count 1
Thread: banana - Count 1
Thread: apple - Count 2
Thread: apple - Count 3
Thread: banana - Count 2
Thread: apple - Count 4
Thread: apple - Count 5
Thread: banana - Count 3
Thread: apple - Count 6
Thread: apple - Count 7
Thread: banana - Count 4
Thread: apple - Count 8
Thread: apple - Count 9
Thread: apple - Count 10
Thread: carrot - Count 2
Thread: banana - Count 5
Thread: banana - Count 6
Thread: banana - Count 7
Thread: banana - Count 8
Thread: carrot - Count 3
Thread: banana - Count 9
Thread: banana - Count 10
Thread: carrot - Count 4
Thread: carrot - Count 5
Thread: carrot - Count 6
Thread: carrot - Count 7
Thread: carrot - Count 8
Thread: carrot - Count 9
Thread: carrot - Count 10
All Threads Completed.
End Time:  Mon Mar 26 01:47:15 EDT 2018\n\n"

printf "Actual Output:\n"
L11 apple 10 banana 25 carrot 100
printf "________________________________________________________________\n\n"

#Single Thread Test
printf "Single Thread Test - Testing Input: [apple 10]\n\n"
printf "Expected Output:\n"
printf "Start Time:  Mon Mar 26 01:53:27 EDT 2018
Thread: apple - Count 1
Thread: apple - Count 2
Thread: apple - Count 3
Thread: apple - Count 4
Thread: apple - Count 5
Thread: apple - Count 6
Thread: apple - Count 7
Thread: apple - Count 8
Thread: apple - Count 9
Thread: apple - Count 10
All Threads Completed.
End Time:  Mon Mar 26 01:53:27 EDT 2018\n\n"

printf "Actual Output:\n"
L11 apple 10
printf "________________________________________________________________\n\n"

#Long Delay Test
printf "Long Delay Test - Testing Input: [apple, 500, banana, 1000, carrot, 2000]\n\n"
printf "Expected Output:\n"
printf "Start Time:  Mon Mar 26 01:59:17 EDT 2018
Thread: carrot - Count 1
Thread: banana - Count 1
Thread: apple - Count 1
Thread: apple - Count 2
Thread: banana - Count 2
Thread: apple - Count 3
Thread: apple - Count 4
Thread: carrot - Count 2
Thread: banana - Count 3
Thread: apple - Count 5
Thread: apple - Count 6
Thread: banana - Count 4
Thread: apple - Count 7
Thread: apple - Count 8
Thread: carrot - Count 3
Thread: banana - Count 5
Thread: apple - Count 9
Thread: apple - Count 10
Thread: banana - Count 6
Thread: carrot - Count 4
Thread: banana - Count 7
Thread: banana - Count 8
Thread: carrot - Count 5
Thread: banana - Count 9
Thread: banana - Count 10
Thread: carrot - Count 6
Thread: carrot - Count 7
Thread: carrot - Count 8
Thread: carrot - Count 9
Thread: carrot - Count 10
All Threads Completed.
End Time:  Mon Mar 26 01:59:37 EDT 2018\n\n"

printf "Actual Output:\n"
L11 apple 500 banana 1000 carrot 2000
printf "________________________________________________________________\n\n"

#Empty Test
printf "Empty Test - Testing Input: []\n\n"
printf "Expected Output:\n"
printf "Start Time:  Mon Mar 26 02:00:36 EDT 2018
All Threads Completed.
End Time:  Mon Mar 26 02:00:36 EDT 2018\n\n"

printf "Actual Output:\n"
L11
printf "________________________________________________________________\n\n"

#Non-Integer Error Test
printf "Non-Integer Error Test - Testing Input: [apple, 10, banana, 25.5]\n\n"
printf "Expected Output:\n"
printf "Start Time:  Mon Mar 26 02:04:28 EDT 2018
The delay for the 2nd thread is not an integer

All Threads Completed.
End Time:  Mon Mar 26 02:04:28 EDT 2018\n\n"

printf "Actual Output:\n"
L11 apple 10 banana 25.5
printf "________________________________________________________________\n\n"

#Missing Argument Error Test
printf "Missing Argument Error Test - Testing Input: [apple, 10, banana]\n\n"
printf "Expected Output:\n"
printf "Start Time:  Mon Mar 26 02:05:41 EDT 2018
Missing Argument: Each thread must have a name and delay.\n\n"

printf "Actual Output:\n"
L11 apple 10 banana
printf "________________________________________________________________\n\n"
