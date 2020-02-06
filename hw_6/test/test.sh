go build ../main.go

./main --from=test.in --to=test.out --limit=33
diff test.out expected.out
if [ $? -ne 0 ]
then
  echo "test failed, files not equal"
  cat test.out
  exit 1
fi

./main --from=test.in --to=test.out --limit=33 --offset=10
diff test.out expected.out
if [ $? -ne 0 ]
then
  echo "test failed, files not equal"
  exit 2
fi
