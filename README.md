# bashT [![Circle CI](https://circleci.com/gh/progrium/basht.png?style=shield)](https://circleci.com/gh/progrium/basht)

Basht is a minimalist Bash testing utility. You can write tests in pure Bash, just pass it one or more Bash source files that define tests. Tests are functions that start with `T_`:

```
T_additionUsingBC() {
  result="$(echo 2+2 | bc)"
  [[ "$result" -eq 4 ]]
}

T_additionUsingDC() {
  result="$(echo 2 2+p | dc)"
  [[ "$result" -eq 4 ]]
}
```

Tests that return non-zero fail. Only the return value will fail a test, as `set -e` is not used. You can fail a test at any time by explicitly calling `return` with non-zero.

## Getting basht

Download and uncompress the latest binary tarball from [releases](https://github.com/progrium/basht/releases).

Alternatively, if you happen to also be using Go, you can install with `go get`:

	$ go get github.com/progrium/basht

## Running tests

Any filenames passed to basht are loaded and any tests found will be run. Take advantage of globbing for multiple files or directories of tests.

```
$ basht tests/example.bash
=== RUN T_additionUsingBC
--- PASS T_additionUsingBC (0s)
=== RUN T_additionUsingDC
--- PASS T_additionUsingDC (0s)

Ran 2 tests.

PASS
```

If tests pass, basht will exit zero. If any tests failed, basht exists non-zero with the number of failed tests.

```
$ basht tests/fails.bash
=== RUN T_failEquals
--- FAIL T_failEquals (0s)
=== RUN T_failMessage
--- FAIL T_failMessage (0s)
    tests/fails.bash:19: This is a fail message.

=== RUN T_failReturn
--- FAIL T_failReturn (0s)
=== RUN T_failSuccess
--- FAIL T_failSuccess (0s)

Ran 4 tests. 4 failed.

FAIL
```

## Macros

Basht provides no special assertions or helpers. However, there is one macro basht provides:

### $T_fail

Calling `$T_fail <message>` marks a test as failed and includes a failure message. There is an example above of how this is shown in the output. It includes the filename and line number with the message.

Keep in mind it does not return, so if used before the end of a test, you must return after.

```
T_failMessage() {
	false || $T_fail "This is a fail message."
}

T_failMultiple() {
	if ! something; then
		$T_fail "Something failed."
		return
	fi
	if ! another; then
		$T_fail "Another failed."
		return
	fi
}
```

## License

BSD
