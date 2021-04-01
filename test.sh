#### -- Print Header -- ####
RESET="\x1b[0m"
BRIGHT="\x1b[1m"
RED="\x1b[31m"
GREEN="\x1b[32m"

printf "\E[H\E[2J" ## Clear screen
printf $BRIGHT
echo "Launching Rubik Performance Test...$RESET\n"

## Initialize
count=0
solved=0
htm_best=420
htm_worst=0
htm_cumulative=0
time_best=420
time_worst=0
time_cumulative=0

go build


#### -- Test Function -- ####
unit_test()
{
	## Initialize
	FILEPATH=$1
	NUM=$2
	if [[ $FILEPATH =~ "." ]]
	then
		if [[ $FILEPATH =~ "random" ]]
		then
			Filename="static_$NUM"
		else
			Filename=$(echo $FILEPATH | cut -d "." -f 1 | cut -d "/" -f 2)
		fi
	elif [[ $NUM ]]
	then
		Filename="dynamic_$NUM"
	else
		Filename="command_line"
	fi

	## Run
	cmd="./Rubik $FILEPATH"
	output=$(eval "$cmd")
	incorrect=$(echo "$output" | tail -n 9 | head -n 1 )
	
	## Time
	time=$(echo "$output" | tail -n 1)
	prefix=$(echo "$time" | rev | cut -c-1-2 | rev | cut -c-1-1)
	if [ "$prefix" = "m" ]
	then
		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
		time_cut=$(echo "scale = 9; ($time_cut / 1000)" | bc)
	elif [ "$prefix" = "µ" ]
	then
		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
		time_cut=$(echo "scale = 9; ($time_cut / 1000000)" | bc)
	elif [ "$prefix" = "n" ]
	then
		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
		time_cut=$(echo "scale = 9; ($time_cut / 1000000000)" | bc)
	else
		time_cut=$(echo "$time" | rev | cut -c2-42 | rev)
	fi

	time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
	time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
	worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
	best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
	if [ "$time_up" -gt "$worst_up" ]
	then
		time_worst=$time_cut
	fi
	if [ "$time_up" -lt "$best_up" ]
	then
		time_best=$time_cut
	fi

	## HTM
	htm=$(echo "$output" | tail -n 7 | head -n 1 | rev | cut -c1-2 | rev )
	if [ "$htm" -gt "$htm_worst" ]
	then
		htm_worst=$htm
	fi
	if [ "$htm" -lt "$htm_best" ]
	then
		htm_best=$htm
	fi
	htm_cumulative=$(echo "scale = 9; $htm_cumulative + $htm" | bc)

	## Print Result
	if [ "$incorrect" == "Error: Solution Incorrect :(" ]
	then
		printf "$RED%-23s %-15s %-7s %s$RESET\n" $Filename "ERROR" $htm $time
	else
		printf "$GREEN%-23s %-15s %-7s %s$RESET\n" $Filename "OK" $htm $time
		((solved+=1))
	fi
	((count+=1))
}


#### -- 10 Random static unit tests -- ####
echo "$BRIGHT#### ---- Random mix tests ---- ####\n$RESET"
printf "%-23s %-15s %-7s %s\n" "Mix" "Solved" "HTM" "Solve Time"
random=0
while [ $random -lt 10 ]
do
	unit_test mix/random$random.txt $random
	random=$(($random + 1))
done

#### -- 10 --random dynamic tests -- ####
echo
random=0
while [ $random -lt 10 ]
do
	unit_test -r $random
	random=$(($random + 1))
done

# #### -- RANDOM STATS -- ####
echo $BRIGHT
echo "Half-turn metric$RESET"
mean_int=$(echo "scale = 0; $htm_cumulative / $count" | bc)
mean_float=$(echo "scale = 2; $htm_cumulative / $count" | bc)
if [ "$htm_best" -lt 45 ]
then
	printf $GREEN
	echo "Best:\t	$htm_best $RESET"
else
	printf $RED
	echo "Best:\t	$htm_best $RESET"
fi
if [ "$htm_worst" -lt 45 ]
then
	printf $GREEN
	echo "Worst:\t	$htm_worst $RESET"
else
	printf $RED
	echo "Worst:\t	$htm_worst $RESET"
fi
if [ "$mean_int" -lt 45 ]
then
	printf $GREEN
	echo "Mean:\t	$mean_float $RESET"
else
	printf $RED
	echo "Mean:\t	$mean_float $RESET"
fi

echo $BRIGHT
echo "Solve time $RESET"
time_mean=$(echo "scale = 9; $time_cumulative / $count" | bc)
best_cut=$(echo $time_best | cut -d "." -f 1)
if [ "$best_cut" == "" ]
then
	printf $GREEN
	echo "Best:\t	$time_best\x1b[32ms $RESET"
else
	if [ "$best_cut" -lt 23 ]
	then
		printf $GREEN
		echo "Best:\t	$time_best\x1b[32ms $RESET"
	else
		printf $RED
		echo "Best:\t	$time_best\x1b[31ms $RESET"
	fi
fi
worst_cut=$(echo $time_worst | cut -d "." -f 1)
if [ "$worst_cut" == "" ]
then
	printf $GREEN
	echo "Worst:\t	$time_worst\x1b[32ms $RESET"
else
	if [ "$worst_cut" -lt 23 ]
	then
		printf $GREEN
		echo "Worst:\t	$time_worst\x1b[32ms $RESET"
	else
		printf $RED
		echo "Worst:\t	$time_worst\x1b[31ms $RESET"
	fi
fi
mean_cut=$(echo $time_mean | cut -d "." -f 1)
if [ "$mean_cut" == "" ]
then
	printf $GREEN
	echo "Mean:\t	$time_mean\x1b[32ms $RESET"
else
if [ "$mean_cut" -lt 23 ]
	then
		printf $GREEN
		echo "Mean:\t	$time_mean\x1b[32ms $RESET"
	else
		printf $RED
		echo "Mean:\t	$time_mean\x1b[31ms $RESET"
	fi
fi

#### -- UNIT TESTS -- ####
echo $BRIGHT
echo "#### ---- Unit tests ---- ####\n $RESET"
unit_test mix/subject.txt
unit_test mix/subject2.txt
unit_test mix/all.txt
unit_test mix/only0.txt
unit_test mix/only1.txt
unit_test mix/only2.txt
unit_test mix/spacing.txt
unit_test mix/F1.txt
unit_test mix/F2.txt
unit_test mix/F3.txt
unit_test mix/F4.txt
unit_test mix/F5.txt
unit_test mix/U_F.txt
unit_test mix/empty.txt
unit_test mix/matsValk555WR.txt
unit_test mix/hard.txt
unit_test mix/superflip.txt
unit_test "\"U2 R2 B2 D2 R2 L2 U2 B2 D2 R2 L R U2 F2 F' B R2 F D2 U D\""

#### -- STATS -- ####
echo
if [ "$solved" == "$count" ]
then
	echo "\nSolved\x1b[32m $solved of $count$RESET total cubes\n"
elif [ "$solved" == "0" ]
then
	echo "\nSolved\x1b[31m $solved of $count$RESET total cubes\n"	
else
	echo "\nSolved\x1b[33m $solved of $count$RESET total cubes\n"
fi

rm Rubik
