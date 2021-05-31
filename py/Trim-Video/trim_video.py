import os, sys
from moviepy.video.io.ffmpeg_tools import ffmpeg_extract_subclip
import moviepy.editor as editor

def to_seconds(f):
    f = float(f)
    return int(f)*60 + (f - round(int(f), 2))*100

src_file = sys.argv[1]
filename, ext = os.path.splitext(src_file)
dest_file = filename + " - Trimmed" + ext

start = to_seconds(sys.argv[2]) if sys.argv[2]!="_" else 0
end = to_seconds(sys.argv[3]) if sys.argv[3]!="_" else int(editor.VideoFileClip(src_file).duration)
ffmpeg_extract_subclip(src_file, start, end, targetname=dest_file)
