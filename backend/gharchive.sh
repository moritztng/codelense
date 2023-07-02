#!/bin/bash
for i in {0..699}
do
  wget -qO- "https://storage.googleapis.com/codelens_gharchive/events$(printf %012d $i).csv.gz" | gzip -dc  | tail -n +2
done
