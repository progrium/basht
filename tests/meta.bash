
T_fails() {
	$SELF "$(dirname $BASH_SOURCE)/fails.bash" &> /dev/null
	[[ $? == 4 ]]
}

T_passes() {
	$SELF "$(dirname $BASH_SOURCE)/passes.bash" &> /dev/null
}

T_glob() {
	$SELF $(dirname $BASH_SOURCE)/glob*.bash \
		| grep "Ran 2 tests." > /dev/null
}

T_failMessage() {
	$SELF "$(dirname $BASH_SOURCE)/fails.bash" \
		| grep "This is a fail message." > /dev/null
}

T_example() {
	$SELF "$(dirname $BASH_SOURCE)/example.bash" &> /dev/null
}
