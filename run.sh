#!/bin/bash

nohup ./user &
nohup ./comment &
nohup ./favorite &
nohup ./relation &
nohup ./video &
nohup ./apis &

echo "Running"