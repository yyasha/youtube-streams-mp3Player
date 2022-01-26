Youtube-streams-mp3Player
-----------------
Playing music or other streams in your console

Go installing
-----------
Official guide: https://golang.org/doc/install

Building
---------------
```
go build player.go
```

Usage
----------------
```
./player -url "YouTube url"
```

Arguments
---------------
```
-display
```
Show display (default off)

```
-url "..."
```
Link to the video on YouTube

```
-volume 76
```
Set volume (default 100)

Example
----------------
Play lofi hip hop at volume 50 without display
```
./player -url "https://www.youtube.com/watch?v=5qap5aO4i9A" -volume 50
```
