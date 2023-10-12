#!/usr/bin/bash
echo "Hello Goofer!"
while true; do
    inotifywait -e modify,create,delete,attrib -rqq .
    make -s build
    echo "[$(date +"%Y-%m-%d %T")] restarted"
    sleep 1
done
