import os, time, shutil

EXTENSIONS = (".jpg", ".png", ".webp", ".gif", ".mp4", ".mkv", ".jpeg", ".heic", ".mov")

origin_path = r'D:\OneDrive\图片\本机照片'
target_path = r'D:\OneDrive\图片\本机照片'
fs = os.listdir(origin_path)
for f in fs:
    fpath = os.path.join(origin_path, f)
    if os.path.isdir(fpath) or not f.lower().endswith(EXTENSIONS):
        continue
    print(f)
    timeStamp = os.path.getmtime(fpath)
    timeArray = time.localtime(timeStamp)
    y = timeArray.tm_year
    mo = timeArray.tm_mon
    mo = '%02d' % mo
    y_mo = '%d-%s' % (y, mo)
    folder = os.path.join(target_path, y_mo)
    if not os.path.exists(folder):
        os.mkdir(folder)
    shutil.move(fpath, os.path.join(folder, f))
