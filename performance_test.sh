#### -- Print Header -- ####
printf "\E[H\E[2J"
echo "\x1b[1mLaunching Rubik Performance Test...\x1B[0m\n"

count=0
solved=0
# best=420
# worst=0
# mcumulative=0
# tcumulative=0
SECONDS=0

go build
echo "Mix\t\t\tSolved\t\tHTM\tsolve time\x1b[0m"

cmd="./Rubik \"U' F'\""
output=$(eval "$cmd")
incorrect=$(echo "$output" | head -n 2 | tail -n 1 )
time=$(echo "$output" | tail -n 1)
len=$(echo "$output" | tail -n 4 | head -n 1 | wc -w)
if [ "$incorrect" == "Error: Solution incorrect :(" ]
then
	echo "\x1b[31mU' F':\t\t\tERROR\t $len\t$time\x1b[0m"
else
	echo "\x1b[32mU' F':\t\t\tOK\t $len\t$time\x1b[0m"
	((solved+=1))
fi
((count+=1))

rm Rubik

echo "\n\n\x1b[1mAll Rubik tests finished\x1b[0m\nTotal runtime $SECONDS seconds"

if [ "$solved" == "$count" ]
then
	echo "\n\x1b[32mPassed $solved of $count total tests\x1b[0m\n"
elif [ "$solved" == "0" ]
then
	echo "\n\x1b[31mPassed $solved of $count total tests\x1b[0m\n"	
else
	echo "\n\x1b[33mPassed $solved of $count total tests\x1b[0m\n"
fi
