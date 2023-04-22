output=`gofmt -l .`
if [ -n "$output" ]; then
  echo "Formatting check error. Format the following files:"
  echo $output
  exit 1
fi
