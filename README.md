# GoAutoYT
<p align="center">GoAutoYT makes it easy for you to automatically download videos from as many YouTube channels as you'd like.</p>

<p align="center"><img src="https://raw.githubusercontent.com/XiovV/go-auto-yt/master/demo.png" width=700 alt="Screenshot of Example Documentation created with Slate"></p>

Features
------------
* **Clean, very simple design** - The dashboard only contains an input form where you can add a channel and configure checking intervals and what to download, and a little list of all your channels where you can delete them or tell the server to check for new uploads immediately.

* **Everything is on a single page** - You can view and control everything from just one page. 

* **Makes downloading videos/audio automatically very easy** - Just paste a link of a channel you want to download, set a checking interval and that's it, the server will keep checking for new uploads and download if necessary.

Getting Started (without Docker)
------------
### Prerequisites
* **Windows, Mac or Linux** - Only tested on Linux, but should work on Mac and Windows
* **Go, version 1.13.4 or newer**
* **[YouTube API Key](https://www.slickremix.com/docs/get-api-key-for-youtube/)** - Step by step guide on how to get it (I am planning on making a crawler so the backend doesn't have to rely on the API)

### Setting Up (Tested on Linux, but should work on Mac. Windows - not sure)
```
git clone https://github.com/XiovV/go-auto-yt.git
export API_KEY=your_youtube_api_key
cd go-auto-yt
go build
./go-auto-yt
```

You can now go to https://localhost:8080 and start using GoAutoYT.

Getting Started (with Docker)
------------
### Prerequisites
* **Docker**

### Running The Container
```
TODO: Create image and write guide
```

## Roadmap
* Create a youtube-dl wrapper so the backend doesn't have to rely on YouTube API and to fix failed downloads
* Create a docker image

## Known Issues
* Sometimes the video just doesn't want to download and I have no idea why, but I'll probably fix it eventually. When a download fails it will be displayed in a list so you can manually download the video.
* Checking interval is hardcoded to 5 hours, at the moment you cannot change it.

## Built With
* [Go](https://golang.org/) - Go Language
* [Gorilla Mux](https://github.com/gorilla/mux) - Go Multiplexer
* [Bootstrap](https://getbootstrap.com/) - CSS Framework