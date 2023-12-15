```
                                      __               
    ____  _________  _  ____  _______/ /__  ____ _____ 
   / __ \/ ___/ __ \| |/_/ / / / ___/ / _ \/ __ `/ __ \
  / /_/ / /  / /_/ />  </ /_/ / /__/ /  __/ /_/ / / / /
 / .___/_/   \____/_/|_|\__, /\___/_/\___/\__,_/_/ /_/ 
/_/                    /____/                          

by darkmage | @evildojo666 | https://www.evildojo.com

```

# Build

```
make
```

# Usage

```
./proxyclean -f [proxies.txt] -t [threadcount]
```

`proxyclean` defaults to using the file `proxies.txt` in the local directory, as well as a default threadcount of 10.

Running `proxyclean` will begin spawning `threadcount` goroutines that will each attempt to connect to a different IP address as if it were a SOCKS5 proxy, and then make a GET request to `www.google.com`.

`proxyclean` will output each IP address that returns a successful request.

