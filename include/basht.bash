
__contains() {
	local e; for e in "${@:2}"; do [[ "$e" == "$1" ]] && return 0; done; return 1
}

__failMacro() {
	local line="$1"; shift
	__test_status=1
	__test_message="$line: $*"
}
readonly T_fail='eval __failMacro $BASH_SOURCE:$LINENO'

main() {
	local run failed start stop duration
	declare -a processed
	run=0
	failed=0
	for file in "$@"; do
		source "$file"
		for t in $(declare -F | grep 'declare -f T_' | awk '{print $3}'); do
			if ! __contains "$t" "${processed[@]}"; then
				unset __test_status __test_message
				echo "=== RUN $t"
				start="$SECONDS"
				$t
				__test_status=${__test_status:-$?}
				stop="$SECONDS"
				duration=$((stop-start))
				processed+=("$t")
				run=$((run+1))
				if [[ "$__test_status" == 0 ]]; then
					echo "--- PASS $t (${duration}s)"
				else
					failed=$((failed+1))
					echo "--- FAIL $t (${duration}s)"
					if [[ "$__test_message" ]]; then
						echo "    $__test_message"
						echo
					fi
				fi
			fi
		done
	done
	echo
	if [[ "$failed" == "0" ]]; then
		echo "Ran $run tests."
		echo
		echo "PASS"
	else
		echo "Ran $run tests. $failed failed."
		echo
		echo "FAIL"
		exit $failed
	fi
}
