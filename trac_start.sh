#!/bin/sh
tracd -d -p 8000 --auth=trac,./trac/passwd.digest,localhost ./trac/
