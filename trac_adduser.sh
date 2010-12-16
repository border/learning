#!/bin/sh
USER=border

htdigest -c ./trac/passwd.digest localhost $USER

trac-admin ./trac/ permission add admin TRAC_ADMIN

trac-admin ./trac/ permission add $USER admin

trac-admin ./trac/ permission remove anonymous TICKET_CREATE
trac-admin ./trac/ permission remove anonymous TICKET_MODIFY
