# txt-ink

This is a toy project for use with the [Pimoroni Inky pHat](https://shop.pimoroni.com/products/inky-phat).

We have two applications that work together. The Go app `txt` is a simple web server that writes text to an image file. The Python app `ink` takes the image and draws it on the pHat using the [Pimonori inky library](https://github.com/pimoroni/inky). They can be glued together with a watch script that automatically runs `txt` when `ink` updates the image file.
