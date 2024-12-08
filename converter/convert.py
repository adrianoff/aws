import imgkit
import sys
import os
from PIL import Image


def convert_to_monochrome(input_image_path, output_image_path):
    # Open the input image
    img = Image.open(input_image_path)

    # Convert the image to grayscale
    img = img.convert("L")

    # Flip the image
    img = img.transpose(Image.FLIP_LEFT_RIGHT)

    # Convert the grayscale image to monochrome (1-bit)
    img = img.point(lambda x: 0 if x < 210 else 255, '1')

    # Save the monochrome image as BMP format
    img.save(output_image_path)


try:
    # Path to save the output BMP image
    output_image_path = 'output_image_tmp.bmp'
    output_image_path_monochrome = sys.argv[2]

    # Custom CSS to set width and height
    custom_css = '''
    <style>
        body {
            width: 800px;
            height: 480px;
        }
    </style>
    '''

    # Options for imgkit to convert HTML to an image
    options = {
        'format': 'bmp',
        'quality': 100,
        'width': 800,
        'height': 480,
        'disable-smart-width': '',
        'quiet': '',
        "xvfb": ""
    }

    # Convert HTML to BMP image with custom CSS
    imgkit.from_string(sys.argv[1], output_image_path, options=options)
    convert_to_monochrome(output_image_path, output_image_path_monochrome)

    #os.remove(output_image_path)
except Exception as err:
    print(err)

print("Ok")
