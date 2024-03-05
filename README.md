# Unifi Watcher

## What is it?
Long story short, I use unifi cameras and want a way to detect events and fire manual functions (while this is possible via home assistant, the lag between detect and notification has been getting worse for me) so after doing some digging into my UCK I discovered the entire system relies on and uses a postgres db. This simple stupid little thing just queries the event table on a regular basis for changes and fires whatever function you give it for each event.


## How do I use it?
`go get -u github.com/tehzwen/unifi-watcher`

- First you'll need to enable ssh on your unifi protect controller (https://help.ui.com/hc/en-us/articles/204909374-UniFi-Connect-with-SSH-Advanced)
- Once that is enabled, you'll need to build your go project using the following flags
    - GOARCH=arm64
    - GOOS=linux
- Once you have something built just sftp it to the controller and run it. I recommend using something like systemd or other services to keep it running in the background.

I personally use this for hitting a local notification service that I run at home for push notifications/emails but you can fire any type of go code you'd like inside the Watch loop.

## What's included?
- Simple wrapper for a watch func
- Examples on using it
- Docker compose setup for testing it against a psql db
