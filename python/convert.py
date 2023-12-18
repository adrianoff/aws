import imgkit
from PIL import Image


def convert_to_monochrome(input_image_path, output_image_path):
    # Open the input image
    img = Image.open(input_image_path)

    # Convert the image to grayscale
    img = img.convert("L")

    # Flip the image
    img = img.transpose(Image.FLIP_LEFT_RIGHT)

    # Convert the grayscale image to monochrome (1-bit)
    img = img.point(lambda x: 0 if x < 128 else 255, '1')

    # Save the monochrome image as BMP format
    img.save(output_image_path)


# Path to your HTML file
html_file_path = 'index.html'

# Path to save the output BMP image
output_image_path = 'output_image.bmp'

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
    'disable-smart-width': ''
}

# Convert HTML to BMP image with custom CSS
imgkit.from_file(html_file_path, output_image_path, options=options)
convert_to_monochrome(output_image_path, 'output_image_monochrome.bmp')

print(f"HTML page converted to monochrome BMP image (800x480): {output_image_path}")
