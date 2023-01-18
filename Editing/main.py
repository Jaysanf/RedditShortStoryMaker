from moviepy.editor import *
from moviepy.video.VideoClip import 
from datetime import datetime
from const import *


def print_hi(name):
    # Use a breakpoint in the code line below to debug your script.
    print(f'Hi, {name}')  # Press Ctrl+F8 to toggle the breakpoint.

def combineAudioAndVideo(directory:str):
    mp3Files = os.listdir(BUNDLE_DIR + "/" + directory + "/" + MP3_DIR)
    txtFiles = os.listdir(BUNDLE_DIR + "/" + directory + "/" + TXT_DIR)
    clip = VideoFileClip(BUNDLE_DIR + "/" + directory + "/" + VIDEO_FILE)

    generator = lambda txt: TextClip(txt, \
                                     font='Berlin-Sans-FB-Demi-Bold', \
                                     size=(350, None), \
                                     fontsize=40, \
                                     color='white', \
                                     stroke_color='black', \
                                     stroke_width=2.5, \
                                     method='caption', \
                                     )
    subtitles = Subt
    return
def get_latest_dir() -> str:
    directories = os.listdir(BUNDLE_DIR)
    directories.remove('.gitkeep')
    if not directories:
        return ""

    latest_dir = datetime.strptime("01-01-2023_00-00-00", '%m-%d-%Y_%H-%M-%S')
    latest_dir_str = ""
    for _dir in directories:
        current_dir = datetime.strptime(_dir, '%m-%d-%Y_%H-%M-%S')
        if current_dir > latest_dir:
            latest_dir = current_dir
            latest_dir_str = _dir
    return latest_dir_str

# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    dir_name = get_latest_dir()
    combineAudioAndVideo(dir_name)
    print_hi('PyCharm')

