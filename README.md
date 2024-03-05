# Unifi Watcher

## What is it?
Long story short, I use unifi cameras and want a way to detect events and fire manual functions (while this is possible via home assistant, the lag between detect and notification has been getting worse for me) so after doing some digging into my UCK I discovered the entire system relies on and uses a postgres db. This simple stupid little thing just queries the event table on a regular basis for changes and fires whatever function you give it for each event.
