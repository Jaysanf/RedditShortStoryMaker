import csv
import os

from moviepy.editor import *
from moviepy.video.tools.subtitles import SubtitlesClip
from moviepy.video.fx import resize
from datetime import datetime
from const import *


def combineAudioAndVideo(directory_path:str):

    audio = AudioFileClip(directory_path + "\\" + AUDIO_FILE)
    clip = VideoFileClip(directory_path + "\\" + VIDEO_FILE)

    generator = lambda txt: TextClip(txt,
                                     font='Berlin-Sans-FB-Demi-Bold',
                                     size=(700, None),
                                     fontsize=60,
                                     color='white',
                                     stroke_color='black',
                                     stroke_width=2.5,
                                     method='caption',
                                     )
    subs = read_subs_csv(directory_path + "\\" + SUB_FILE)
    subtitles = SubtitlesClip(subs, generator)
    clip_duration = min(audio.duration, clip.duration, subtitles.duration,MAX_SECONDS)

    audio = audio.subclip(0, clip_duration)
    clip = clip.subclip(0, clip_duration)
    subtitles = subtitles.subclip(0, clip_duration)

    clip = clip.set_audio(audio)

    result = CompositeVideoClip([clip, subtitles.set_pos('center')])
    #result = resize.resize(result,newsize=(1080,1920))

    result.write_videofile(directory_path + "\\" + "final.mp4", fps=clip.fps, temp_audiofile="temp-audio.m4a",
                           remove_temp=True, codec="libx264", audio_codec="aac")
    return


def get_latest_dir() -> str:
    working_dir = get_working_dir_path() + "\\" + BUNDLE_DIR
    directories = os.listdir(working_dir)
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
    return working_dir + "\\" + latest_dir_str


def get_working_dir_path():
    working_path = os.getcwd()
    working_dir = os.path.dirname(working_path)
    if working_path[len(working_path) - len("Editing"):] == "Editing":
        return working_dir
    return working_path

def read_subs_csv(file_path):
    subs = []
    with open(file_path, newline='',encoding='utf-8') as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            start_time = float(row[0])
            end_time = float(row[1])
            subtitles = row[2]
            subs.append(((start_time, end_time), subtitles))
    return subs


# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    dir_path = get_latest_dir()
    combineAudioAndVideo(dir_path)
