#!/bin/bash

java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/0.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/1.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/2.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/3.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/4.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/5.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/6.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/7.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/8.txt  2>&1
java -jar zed/target/zed-1.0-SNAPSHOT.jar >> dist/9.txt  2>&1

sleep 10

a=0
a=$((a+`awk 'END {print}' dist/0.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/1.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/2.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/3.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/4.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/5.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/6.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/7.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/8.txt | awk '{print int($0)}'`));
a=$((a+`awk 'END {print}' dist/9.txt | awk '{print int($0)}'`));

n=`echo "sclae=2; $a/10" | bc`
echo $n
