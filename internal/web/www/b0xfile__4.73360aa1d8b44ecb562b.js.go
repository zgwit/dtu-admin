// Code generaTed by fileb0x at "2020-09-29 16:10:37.5586576 +0800 CST m=+0.163969401" from config file "b0x.yaml" DO NOT EDIT.
// modified(2020-09-29 16:08:48.3461282 +0800 CST)
// original path: portal\dist\4.73360aa1d8b44ecb562b.js

package wwwFiles

import (
	"os"
)

// File473360aa1d8b44ecb562bJs is "/4.73360aa1d8b44ecb562b.js"

func init() {

	f, err := FS.OpenFile(CTX, "/4.73360aa1d8b44ecb562b.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(File473360aa1d8b44ecb562bJs)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}