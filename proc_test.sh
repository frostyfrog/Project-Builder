#!/bin/bash
for i in {1..5}; do
	echo -n "${i}."
	sleep 1
done
echo
exec 1>&2
for i in {1..2}; do
	echo -n "${i}."
	sleep 1
done
echo
