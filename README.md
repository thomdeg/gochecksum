# gochecksum
A tiny multiplatform tool to calculate file checksums in Go (golang). It can help to verify downloaded files.


### Build

    $ go build .
    
### Usage

We can test the compiled program on its source code:

    $ ./gochecksum gocheck_main.go
    file: gocheck_main.go
    md5   : 204c4050183e1e53608267cd33658122
    sha1  : f0772d4361cea923973035068df7216648308d92
    sha224: 0a681803858061b3acd8fd730e67ccb4d125e89207dd7005da3d8d1e
    sha256: 174ea795e620d9b0a6474de12acea18bc9171ba47bfaf95aadee3dfc646bc72a
    sha384: 70fcc357a25e2d7d3a088f8ea1cf390780b8aa31255522ba9f952431a45cf8c05ddd0ee274e1dbe05969336fcd00791a
    sha512: b0f7773476044fa9c27ddba5db2eb400a25286e008cb43b3f2e48ec81c5b2d3b13e1cfc82ef1dfc99a08f0ef33748582498642fc387527ea25d929dcbdf83701
    -------------

or the previous version from commit f98c138:

    $ ./gochecksum <(git show f98c138:./gocheck_main.go)
    file: /dev/fd/63
    md5   : ab2a7004765b77d4a09bf09985e84864
    sha1  : 269baabacba334e204559cfb37ad9ab1430f95c3
    sha224: 64b19a478afd7589144f8393778a8e670488ff5381a130dd9d429cb9
    sha256: 2e12eeb446b6ccd835149f5e9372d7cec33f3b7de20401ad4a51bca84e4a6243
    sha384: 91343729fc9b3291d0086ed25616bd1c1f9479f4f7f3b83b9229ac565c6dad16c5e9a0ab84eb342503a6b0327ae9f956
    sha512: 7b3877c04bf6cceae3ff0055e02877c041612930095f888b0389e777c25db2e3409cebd6010a0fedf1ed284ec5fae4697bacf8b2e1b187c9a856f90a33b6179f
    -------------


### go version

compiled with  go 1.15.x 
