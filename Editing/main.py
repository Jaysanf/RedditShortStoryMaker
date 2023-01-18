from moviepy.editor import *
from moviepy.video.tools.subtitles import SubtitlesClip
from datetime import datetime
from const import *


def combineAudioAndVideo(directory:str):
    path_to_bundle = BUNDLE_DIR + "/" + directory + "/"
    audio = AudioFileClip(path_to_bundle + AUDIO_FILE).subclip(0,60)
    clip = VideoFileClip(path_to_bundle + VIDEO_FILE).subclip(0,60)
    clip = clip.set_audio(audio)

    generator = lambda txt: TextClip(txt, \
                                     font='Berlin-Sans-FB-Demi-Bold', \
                                     size=(350, None), \
                                     fontsize=40, \
                                     color='white', \
                                     stroke_color='black', \
                                     stroke_width=2.5, \
                                     method='caption', \
                                     )
    subtitles = SubtitlesClip(path_to_bundle + SRT_FILE, generator).subclip(0,60)
    result = CompositeVideoClip([clip, subtitles.set_pos('center')])

    result.write_videofile(path_to_bundle + "final.mp4", fps=clip.fps, temp_audiofile="temp-audio.m4a", remove_temp=True,
                           codec="libx264", audio_codec="aac")
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

