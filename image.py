#!/usr/bin/env python

import sys
from PIL import Image
from inky import InkyWHAT, InkyPHAT


imagePath = sys.argv[1]
img = Image.open(imagePath)
img = img.convert("1", dither = Image.NONE)

display = InkyPHAT("black")
display.set_image(img)
display.show()
