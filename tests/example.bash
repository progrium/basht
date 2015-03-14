
T_additionUsingBC() {
  result="$(echo 2+2 | bc)"
  [[ "$result" -eq 4 ]]
}

T_additionUsingDC() {
  result="$(echo 2 2+p | dc)"
  [[ "$result" -eq 4 ]]
}
