#### -- Print Header -- ####
RESET="\x1b[0m"
BRIGHT="\x1b[1m"
RED="\x1b[31m"
GREEN="\x1b[32m"

printf "\E[H\E[2J" ## Clear screen
echo $BRIGHT
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
	string='My string';
	if [[ $FILEPATH =~ "." ]]
	then
		Filename=$(echo $FILEPATH | cut -d "." -f 1 | cut -d "/" -f 2)
	else
		Filename="command_line"
	fi

	cmd="./Rubik $FILEPATH"
	output=$(eval "$cmd")
	incorrect=$(echo "$output" | tail -n 9 | head -n 1 )
	time=$(echo "$output" | tail -n 1)
	prefix=$(echo "$time" | rev | cut -c-1-2 | rev | cut -c-1-1)
	if [ "$prefix" = "m" ]
	then
		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
		time_cut=$(echo "scale = 9; ($time_cut / 1000)" | bc)
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
	elif [ "$prefix" = "µ" ]
	then
		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
		time_cut=$(echo "scale = 9; ($time_cut / 1000000)" | bc)
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
	elif [ "$prefix" = "n" ]
	then
		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
		time_cut=$(echo "scale = 9; ($time_cut / 1000000000)" | bc)
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
	else
		time_cut=$(echo "$time" | rev | cut -c2-42 | rev)
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
	fi
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
	if [ "$incorrect" == "Error: Solution Incorrect :(" ]
	then
		printf "$RED%-23s %-15s %-7s %s$RESET\n" $Filename "ERROR" $htm $time
	else
		printf "$GREEN%-23s %-15s %-7s %s$RESET\n" $Filename "OK" $htm $time
		((solved+=1))
	fi
	((count+=1))
}


# #### -- 10 Random unit tests -- ####
# echo "\x1b[1m#### ---- Random mix tests ---- ####\n\x1b[0m"
# printf "%-23s %-15s %-7s %s\n" "Mix" "Solved" "HTM" "Solve Time"
# random=0
# while [ $random -lt 10 ]
# do
# 	cmd="./Rubik mix/random$random.txt"
# 	output=$(eval "$cmd")
# 	incorrect=$(echo "$output" | tail -n 9 | head -n 1 )
# 	time=$(echo "$output" | tail -n 1)
# 	prefix=$(echo "$time" | rev | cut -c-1-2 | rev | cut -c-1-1)
# 	if [ "$prefix" = "m" ]
# 	then
# 		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
# 		time_cut=$(echo "scale = 9; ($time_cut / 1000)" | bc)
# 		time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
# 		time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
# 		worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
# 		best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
# 		if [ "$time_up" -gt "$worst_up" ]
# 		then
# 			time_worst=$time_cut
# 		fi
# 		if [ "$time_up" -lt "$best_up" ]
# 		then
# 			time_best=$time_cut
# 		fi
# 	elif [ "$prefix" = "µ" ]
# 	then
# 		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
# 		time_cut=$(echo "scale = 9; ($time_cut / 1000000)" | bc)
# 		time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
# 		time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
# 		worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
# 		best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
# 		if [ "$time_up" -gt "$worst_up" ]
# 		then
# 			time_worst=$time_cut
# 		fi
# 		if [ "$time_up" -lt "$best_up" ]
# 		then
# 			time_best=$time_cut
# 		fi
# 	elif [ "$prefix" = "n" ]
# 	then
# 		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
# 		time_cut=$(echo "scale = 9; ($time_cut / 1000000000)" | bc)
# 		time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
# 		time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
# 		worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
# 		best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
# 		if [ "$time_up" -gt "$worst_up" ]
# 		then
# 			time_worst=$time_cut
# 		fi
# 		if [ "$time_up" -lt "$best_up" ]
# 		then
# 			time_best=$time_cut
# 		fi
# 	else
# 		time_cut=$(echo "$time" | rev | cut -c2-42 | rev)
# 		time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
# 		time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
# 		worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
# 		best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
# 		if [ "$time_up" -gt "$worst_up" ]
# 		then
# 			time_worst=$time_cut
# 		fi
# 		if [ "$time_up" -lt "$best_up" ]
# 		then
# 			time_best=$time_cut
# 		fi
# 	fi
# 	htm=$(echo "$output" | tail -n 7 | head -n 1 | rev | cut -c1-2 | rev )
# 	if [ "$htm" -gt "$htm_worst" ]
# 	then
# 		htm_worst=$htm
# 	fi
# 	if [ "$htm" -lt "$htm_best" ]
# 	then
# 		htm_best=$htm
# 	fi
# 	htm_cumulative=$(echo "scale = 9; $htm_cumulative + $htm" | bc)
# 	if [ "$incorrect" == "Error: Solution Incorrect :(" ]
# 	then
# 		echo "\x1b[31mRandom unit $random:\t\tERROR\t	$htm\t$time\x1b[0m"
# 	else
# 		echo "\x1b[32mRandom unit $random:\t\tOK\t	$htm\t$time\x1b[0m"
# 		((solved+=1))
# 	fi
# 	((count+=1))
# 	random=$(($random + 1))
# done

# #### -- 10 Random -r -- ####
# echo
# random=0
# while [ $random -lt 10 ]
# do
# 	cmd="./Rubik -r"
# 	output=$(eval "$cmd")
# 	incorrect=$(echo "$output" | tail -n 9 | head -n 1 )
# 	time=$(echo "$output" | tail -n 1)
# 	prefix=$(echo "$time" | rev | cut -c-1-2 | rev | cut -c-1-1)
# 	if [ "$prefix" = "m" ]
# 	then
# 		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
# 		time_cut=$(echo "scale = 9; ($time_cut / 1000)" | bc)
# 		time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
# 		time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
# 		worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
# 		best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
# 		if [ "$time_up" -gt "$worst_up" ]
# 		then
# 			time_worst=$time_cut
# 		fi
# 		if [ "$time_up" -lt "$best_up" ]
# 		then
# 			time_best=$time_cut
# 		fi
# 	elif [ "$prefix" = "µ" ]
# 	then
# 		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
# 		time_cut=$(echo "scale = 9; ($time_cut / 1000000)" | bc)
# 		time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
# 		time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
# 		worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
# 		best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
# 		if [ "$time_up" -gt "$worst_up" ]
# 		then
# 			time_worst=$time_cut
# 		fi
# 		if [ "$time_up" -lt "$best_up" ]
# 		then
# 			time_best=$time_cut
# 		fi
# 	elif [ "$prefix" = "n" ]
# 	then
# 		time_cut=$(echo "$time" | rev | cut -c3-42 | rev)
# 		time_cut=$(echo "scale = 9; ($time_cut / 1000000000)" | bc)
# 		time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
# 		time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
# 		worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
# 		best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
# 		if [ "$time_up" -gt "$worst_up" ]
# 		then
# 			time_worst=$time_cut
# 		fi
# 		if [ "$time_up" -lt "$best_up" ]
# 		then
# 			time_best=$time_cut
# 		fi
# 	else
# 		time_cut=$(echo "$time" | rev | cut -c2-42 | rev)
# 		time_cumulative=$(echo "scale = 9; $time_cumulative + $time_cut" | bc)
# 		time_up=$(echo "scale = 0; $time_cut * 1000000000" | bc | cut -d "." -f 1)
# 		worst_up=$(echo "scale = 0; $time_worst * 1000000000" | bc | cut -d "." -f 1)
# 		best_up=$(echo "scale = 0; $time_best * 1000000000" | bc | cut -d "." -f 1)
# 		if [ "$time_up" -gt "$worst_up" ]
# 		then
# 			time_worst=$time_cut
# 		fi
# 		if [ "$time_up" -lt "$best_up" ]
# 		then
# 			time_best=$time_cut
# 		fi
# 	fi
# 	htm=$(echo "$output" | tail -n 7 | head -n 1 | rev | cut -c1-2 | rev )
# 	if [ "$htm" -gt "$htm_worst" ]
# 	then
# 		htm_worst=$htm
# 	fi
# 	if [ "$htm" -lt "$htm_best" ]
# 	then
# 		htm_best=$htm
# 	fi
# 	htm_cumulative=$(echo "scale = 9; $htm_cumulative + $htm" | bc)
# 	if [ "$incorrect" == "Error: Solution Incorrect :(" ]
# 	then
# 		echo "\x1b[31mRandom $random:\t\tERROR\t	$htm\t$time\x1b[0m"
# 	else
# 		echo "\x1b[32mRandom $random:\t\tOK\t	$htm\t$time\x1b[0m"
# 		((solved+=1))
# 	fi
# 	((count+=1))
# 	random=$(($random + 1))
# done

# # #### -- RANDOM STATS -- ####
# echo
# echo "\x1b[1mHalf-turn metric\x1b[0m"
# mean_int=$(echo "scale = 0; $htm_cumulative / $count" | bc)
# mean_float=$(echo "scale = 2; $htm_cumulative / $count" | bc)
# if [ "$htm_best" -lt 45 ]
# then
# 	echo "\x1b[32mBest:\t	$htm_best\x1b[0m"
# else
# 	echo "\x1b[31mBest:\t	$htm_best\x1b[0m"
# fi
# if [ "$htm_worst" -lt 45 ]
# then
# 	echo "\x1b[32mWorst:\t	$htm_worst\x1b[0m"
# else
# 	echo "\x1b[31mWorst:\t	$htm_worst\x1b[0m"
# fi
# if [ "$mean_int" -lt 45 ]
# then
# 	echo "\x1b[32mMean:\t	$mean_float\x1b[0m"
# else
# 	echo "\x1b[31mMean:\t	$mean_float\x1b[0m"
# fi

# echo "\n\x1b[1mSolve time\x1b[0m"
# time_mean=$(echo "scale = 9; $time_cumulative / $count" | bc)
# best_cut=$(echo $time_best | cut -d "." -f 1)
# if [ "$best_cut" == "" ]
# then
# 	echo "\x1b[32mBest:\t	$time_best\x1b[32ms\x1b[0m"
# else
# 	if [ "$best_cut" -lt 23 ]
# 	then
# 		echo "\x1b[32mBest:\t	$time_best\x1b[32ms\x1b[0m"
# 	else
# 		echo "\x1b[31mBest:\t	$time_best\x1b[31ms\x1b[0m"
# 	fi
# fi
# worst_cut=$(echo $time_worst | cut -d "." -f 1)
# if [ "$worst_cut" == "" ]
# then
# 	echo "\x1b[32mWorst:\t	$time_worst\x1b[32ms\x1b[0m"
# else
# 	if [ "$worst_cut" -lt 23 ]
# 	then
# 		echo "\x1b[32mWorst:\t	$time_worst\x1b[32ms\x1b[0m"
# 	else
# 		echo "\x1b[31mWorst:\t	$time_worst\x1b[31ms\x1b[0m"
# 	fi
# fi
# mean_cut=$(echo $time_mean | cut -d "." -f 1)
# if [ "$mean_cut" == "" ]
# then
# 	echo "\x1b[32mMean:\t	$time_mean\x1b[32ms\x1b[0m"
# else
# if [ "$mean_cut" -lt 23 ]
# 	then
# 		echo "\x1b[32mMean:\t	$time_mean\x1b[32ms\x1b[0m"
# 	else
# 		echo "\x1b[31mMean:\t	$time_mean\x1b[31ms\x1b[0m"
# 	fi
# fi
# echo

#### -- UNIT TESTS -- ####
echo "\x1b[1m#### ---- Unit tests ---- ####\n\x1b[0m"
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
	echo "\nSolved\x1b[32m $solved of $count \x1b[0mtotal cubes\n"
elif [ "$solved" == "0" ]
then
	echo "\nSolved\x1b[31m $solved of $count \x1b[0mtotal cubes\n"	
else
	echo "\nSolved\x1b[33m $solved of $count \x1b[0mtotal cubes\n"
fi

#### -- END -- ####
# echo "\n\n\x1b[1mAll Rubik tests finished\x1b[0m" ##\nTotal runtime $SECONDS seconds"
rm Rubik
