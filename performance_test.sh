## echo "Usage: '''go build''' to build the binary 'Rubik'. then ./performance_test.sh"

#### -- Print Header -- ####
start=`date +%s`
printf "\E[H\E[2J"
echo "\x1b[1mLaunching Rubik Performance Test\x1B[0m\n"

count=0
solved=0
# best=420
# worst=0
# mcumulative=0
# tcumulative=0
SECONDS=0

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
