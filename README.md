## Reggie Dahling

I'm making ReginaldBot for myself as a way to play with a personal bot.

Things I think might be fun to add long term:

1. Journaling. /jrnl today I went to the park and it was fun. This should be plain text-ish for eash future grepping and reformatting.
	Aside, super cool idea, track the phone location and log it to this file too
2. ToDod's. /todo learn how to make bots for telegram. This should integrate nicely with my existing taskpaper lists.
3. Alerting, have it tell me my RAID array is failing.
4. Controling IOT. /sprinklers off


## Task List

* factor jrnl responder out as separate goroutine
* make it work with minimal unique prefixes for commands, e.g. '/j' works for '/jrnl'
* factor jrnl out to a separate go package w/ cli
* figure out prometheus alerting webhook endpoint
* convert to binary build and properly install/run on system
* build taskpaper package to include as well


