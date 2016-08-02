# list std lib
go list std > /tmp/std.txt
# list dependencies
go list -f {{.Deps}} | xargs -n1 | sort -u > /tmp/deps.txt
cp /tmp/std.txt /tmp/all.txt
cat /tmp/deps.txt >> /tmp/all.txt
cat /tmp/all.txt | sort -u > /tmp/tmp.txt
cp /tmp/tmp.txt /tmp/temp2.txt
cat /tmp/deps.txt >> /tmp/temp2.txt
cat /tmp/temp2.txt | sort | uniq -d > /tmp/final.txt
cat /tmp/final.txt
