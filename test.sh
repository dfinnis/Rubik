#### -- Print Header -- ####
printf "\E[H\E[2J"
echo "\x1b[1mLaunching Rubik Performance Test...\x1B[0m\n"
# SECONDS=0

count=0
solved=0
htm_best=420
htm_worst=0
htm_cumulative=0
# time_cumulative=0

go build
echo "Mix\t\t\tSolved\t\tHTM\tsolve time\x1b[0m"

#### -- test 0 -- ####
cmd="./Rubik \"U' F'\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
time=$(echo "$output" | tail -n 1)
htm=$(echo "$output" | tail -n 4 | head -n 1 | wc -w)
if [ "$htm" -gt "$htm_worst" ]
then
	htm_worst=$htm
fi
if [ "$htm" -lt "$htm_best" ]
then
	htm_best=$htm
fi
htm_cumulative=$(echo "scale = 9; $htm_cumulative + $htm" | bc)
if [ "$incorrect" == "Error: Solution incorrect :(" ]
then
	echo "\x1b[31mU' F':\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mU' F':\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 1 -- ####
cmd="./Rubik \"U' F2\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
time=$(echo "$output" | tail -n 1)
htm=$(echo "$output" | tail -n 4 | head -n 1 | wc -w)
if [ "$htm" -gt "$htm_worst" ]
then
	htm_worst=$htm
fi
if [ "$htm" -lt "$htm_best" ]
then
	htm_best=$htm
fi
htm_cumulative=$(echo "scale = 9; $htm_cumulative + $htm" | bc)
if [ "$incorrect" == "Error: Solution incorrect :(" ]
then
	echo "\x1b[31mU' F2:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mU' F2:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))


#### -- END -- ####
rm Rubik
echo "\n\n\x1b[1mAll Rubik tests finished\x1b[0m" ##\nTotal runtime $SECONDS seconds"

if [ "$solved" == "$count" ]
then
	echo "\n\x1b[32mPassed $solved of $count total tests\x1b[0m\n"
elif [ "$solved" == "0" ]
then
	echo "\n\x1b[31mPassed $solved of $count total tests\x1b[0m\n"	
else
	echo "\n\x1b[33mPassed $solved of $count total tests\x1b[0m\n"
fi


if [ "$count" != 0 ]
then
	mean_int=$(echo "scale = 0; $htm_cumulative / $count" | bc)
	mean_float=$(echo "scale = 2; $htm_cumulative / $count" | bc)
else
	mean_int="\x1b[31mFailed\x1b[0m"
	mean_float="\x1b[31mFailed\x1b[0m"
fi
echo "Half-turn metric"
if [ "$htm_best" -lt 30 ]
then
	echo "\x1b[32mBest:\t$htm_best\x1b[0m"
else
	echo "\x1b[31mBest:\t$htm_best\x1b[0m"
fi

if [ "$mean_int" -lt 30 ]
then
	echo "\x1b[32mMean:\t      $mean_float\x1b[0m"
else
	echo "\x1b[31mMean:\t      $mean_float\x1b[0m"
fi

if [ "$htm_worst" -lt 30 ]
then
	echo "\x1b[32mWorst:\t$htm_worst\x1b[0m"
else
	echo "\x1b[31mWorst:\t$htm_worst\x1b[0m"
fi
# echo "Best: $htm_best"
# echo "Mean:$mean"
# echo "Worst:$htm_worst\n"
