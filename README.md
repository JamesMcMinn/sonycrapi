sonycrapi
============
Sony Camera Remote API

Based off of documentaiton proived by Sony and using the Sony Camera Reomote API.

More information about the API from Sony can be found [here](https://developer.sony.com/develop/cameras/get-started/).

This has only been tested on my Sony NEX-6, and I cannot make any guarantees about other devices.

Install
===
    go get -u github.com/JamesMcMinn/sonycrapi

Usage 
=== 

Features 
===
- Device Discovery using uPnP
- Ability to control:
	- Shoot mode (Still, Movie, Audio, InvervalStill, Loop Recording)
	- Still quality (RAW+JPEG, Fine, Standard)
	- PostView size (2MP or Original)
- Take & Save still pictures
- Liveview via websocket (see examples folder for client) or raw JPEG bytestream

Todo
===
- Check camera is in correct mode before performing action (e.g. still mode before photo, movie mode before starting to record a movie)
