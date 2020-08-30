#### -- Print Header -- ####
printf "\E[H\E[2J"
echo "\x1b[1mLaunching Rubik Performance Test...\x1B[0m\n"
# SECONDS=0

count=0
solved=0
htm_best=420
htm_worst=0
htm_cumulative=0
time_best=420
time_worst=0
time_cumulative=0

go build

#### -- 10 Random unit tests -- ####
echo "\x1b[1m#### ---- Random mix tests ---- ####\n\x1b[0m"
echo "Mix\t\t\tSolved\t\tHTM\tsolve time\x1b[0m"
random=0
while [ $random -lt 10 ]
do
	cmd="./Rubik \"$(< mix/random$random.txt)\""
	output=$(eval "$cmd")
	incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
		echo "\x1b[31mRandom unit $random:\t\tERROR\t  $htm\t$time\x1b[0m"
	else
		echo "\x1b[32mRandom unit $random:\t\tOK\t  $htm\t$time\x1b[0m"
		((solved+=1))
	fi
	((count+=1))
	random=$(($random + 1))
done

#### -- 10 Random -r -- ####
echo
random=0
while [ $random -lt 10 ]
do
	cmd="./Rubik -r"
	output=$(eval "$cmd")
	incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
		echo "\x1b[31mRandom $random:\t\tERROR\t  $htm\t$time\x1b[0m"
	else
		echo "\x1b[32mRandom $random:\t\tOK\t  $htm\t$time\x1b[0m"
		((solved+=1))
	fi
	((count+=1))
	random=$(($random + 1))
done

# #### -- RANDOM STATS -- ####
echo
echo "\x1b[1mHalf-turn metric\x1b[0m"
mean_int=$(echo "scale = 0; $htm_cumulative / $count" | bc)
mean_float=$(echo "scale = 2; $htm_cumulative / $count" | bc)
if [ "$htm_best" -lt 45 ]
then
	echo "\x1b[32mBest:\t$htm_best\x1b[0m"
else
	echo "\x1b[31mBest:\t$htm_best\x1b[0m"
fi
if [ "$htm_worst" -lt 45 ]
then
	echo "\x1b[32mWorst:\t$htm_worst\x1b[0m"
else
	echo "\x1b[31mWorst:\t$htm_worst\x1b[0m"
fi
if [ "$mean_int" -lt 45 ]
then
	echo "\x1b[32mMean:\t      $mean_float\x1b[0m"
else
	echo "\x1b[31mMean:\t      $mean_float\x1b[0m"
fi

echo "\n\x1b[1mSolve time\x1b[0m"
time_mean=$(echo "scale = 9; $time_cumulative / $count" | bc)
best_cut=$(echo $time_best | cut -d "." -f 1)
if [ "$best_cut" == "" ]
then
	echo "\x1b[32mBest:\t      $time_best\x1b[32ms\x1b[0m"
else
	if [ "$best_cut" -lt 23 ]
	then
		echo "\x1b[32mBest:\t      $time_best\x1b[32ms\x1b[0m"
	else
		echo "\x1b[31mBest:\t      $time_best\x1b[31ms\x1b[0m"
	fi
fi
worst_cut=$(echo $time_worst | cut -d "." -f 1)
if [ "$worst_cut" == "" ]
then
	echo "\x1b[32mWorst:\t      $time_worst\x1b[32ms\x1b[0m"
else
	if [ "$worst_cut" -lt 23 ]
	then
		echo "\x1b[32mWorst:\t      $time_worst\x1b[32ms\x1b[0m"
	else
		echo "\x1b[31mWorst:\t      $time_worst\x1b[31ms\x1b[0m"
	fi
fi
mean_cut=$(echo $time_mean | cut -d "." -f 1)
if [ "$mean_cut" == "" ]
then
	echo "\x1b[32mMean:\t      $time_mean\x1b[32ms\x1b[0m"
else
if [ "$mean_cut" -lt 23 ]
	then
		echo "\x1b[32mMean:\t      $time_mean\x1b[32ms\x1b[0m"
	else
		echo "\x1b[31mMean:\t      $time_mean\x1b[31ms\x1b[0m"
	fi
fi
echo

#### -- UNIT TESTS -- ####
echo "\x1b[1m#### ---- Unit tests ---- ####\n\x1b[0m"
#### -- test 1 -- ####
cmd="./Rubik \"$(< mix/subject.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31msubject:\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32msubject:\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 2 -- ####
cmd="./Rubik \"$(< mix/subject2.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31msubject2:\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32msubject2:\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 3 -- ####
cmd="./Rubik \"$(< mix/all.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mall:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mall:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 4 -- ####
cmd="./Rubik \"$(< mix/only0.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31monly0:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32monly0:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 5 -- ####
cmd="./Rubik \"$(< mix/only1.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31monly1:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32monly1:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 6 -- ####
cmd="./Rubik \"$(< mix/only2.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31monly2:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32monly2:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 7 -- ####
cmd="./Rubik \"$(< mix/spacing.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mspacing:\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mspacing:\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 8 -- ####
cmd="./Rubik \"$(< mix/F1.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mF1:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mF1:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 9 -- ####
cmd="./Rubik \"$(< mix/F2.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mF2:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mF2:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))


#### -- test 10 -- ####
cmd="./Rubik \"$(< mix/F3.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mF3:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mF3:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))


#### -- test 11 -- ####
cmd="./Rubik \"$(< mix/F4.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mF4:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mF4:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))


#### -- test 12 -- ####
cmd="./Rubik \"$(< mix/F5.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mF5:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mF5:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 13 -- ####
cmd="./Rubik \"$(< mix/U_F.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mU_F:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mU_F:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 14 -- ####
cmd="./Rubik \"$(< mix/empty.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mempty:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mempty:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 15 -- ####
cmd="./Rubik \"$(< mix/matsValk555WR.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mmatsValk555WR:\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mmatsValk555WR:\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 16 -- ####
cmd="./Rubik \"$(< mix/hard.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31mhard:\t\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32mhard:\t\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

#### -- test 17 -- ####
cmd="./Rubik \"$(< mix/superflip.txt)\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | tail -n 7 | head -n 1 )
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
	echo "\x1b[31msuperflip:\t\tERROR\t  $htm\t$time\x1b[0m"
else
	echo "\x1b[32msuperflip:\t\tOK\t  $htm\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

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