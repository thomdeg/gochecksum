# gochecksum
A tiny multiplatform tool to calculate file checksums in Go (golang). Can help to verify downloads.


### Build

    $ go build .
    
### Usage

We can test the compiled program on its source code:

    $ ./gochecksum gocheck_main.go 
    
    file: gocheck_main.go
    md5   : ab2a7004765b77d4a09bf09985e84864
    sha1  : 269baabacba334e204559cfb37ad9ab1430f95c3
    sha224: 64b19a478afd7589144f8393778a8e670488ff5381a130dd9d429cb9
    sha256: 2e12eeb446b6ccd835149f5e9372d7cec33f3b7de20401ad4a51bca84e4a6243
    sha384: 91343729fc9b3291d0086ed25616bd1c1f9479f4f7f3b83b9229ac565c6dad16c5e9a0ab84eb342503a6b0327ae9f956
    sha512: 7b3877c04bf6cceae3ff0055e02877c041612930095f888b0389e777c25db2e3409cebd6010a0fedf1ed284ec5fae4697bacf8b2e1b187c9a856f90a33b6179f
    -------------

### go version

compiled with  go 1.15.x 