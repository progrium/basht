
fails() {
	false
}

T_failEquals() {
	[[ "yes" == "no" ]]
}

T_failReturn() {
	return 1
}

T_failSuccess() {
	fails
}

T_failMessage() {
	false || $T_fail "This is a fail message."
}
