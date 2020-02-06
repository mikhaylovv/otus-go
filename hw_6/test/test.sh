go build ../main.go
./main.exe --from=test.in --to=test.out --limit=33 --offset=11
diff test.out expected.out >&2
if [ $? -ne 0 ]
then
  echo "test failed, files noe equal" >&2
  exit 1
fi

./main.exe --from=test.in --to=test.out --limit=33
if [ $? -ne 0 ]
then
  echo "test failed, files noe equal" >&2
  exit 2
fi
diff test.out expected.out
