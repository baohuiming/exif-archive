# 媒体归档工具

本仓库包含两个版本的媒体归档工具：`exif-archive.go` 和 `archive.py`。
支持的文件格式：.jpg, .png, .webp, .gif, .mp4, .mkv, .jpeg, .heic, .mov

## Go 版本

`exif-archive.go` 文件包含了源码，可以利用 EXIF 拍摄日期、文件名中的时间或时间戳、文件创建日期来归档媒体文件。如将`<origin>/file`按日期归档到`<target>/yyyy-mm/file`。

支持的文件名格式：
1. 含有13位时间戳的文件或以mmexport开头的微信聊天文件，如：`wx_camera_1689324886317.jpg`或`mmexport1622020005757.jpg`
2. 含有14位时间或17位时间的文件，如`20210526085304575.jpg`或`20210526085304.jpg`
3. 含有14位时间的文件，但日期与时间之间有分隔符，如`Snapshot_20230626_113855_appname.mp4`
4. 含有格式为yyyy-mm-dd的文件，如`2021-05-26 08.53.04.jpg`

使用 Go 版本，请按照以下步骤操作：
1. 在您的系统上安装 Go。
2. 使用 Go 编译器编译 `exif-archive.go` 文件。
3. 运行编译后的可执行文件（`exif-archive.exe`）来开始归档您的媒体文件。
4. 使用方法：`exif-archive --origin <origin> --target <target>`。

## Python 版本

`archive.py` 文件包含了源码，只根据文件修改日期来归档媒体文件。

使用 Python 版本，请按照以下步骤操作：
1. 在您的系统上安装 Python。
2. 使用 Python 解释器运行 `archive.py` 文件。
3. 脚本将根据文件创建日期自动归档您的媒体文件。

**另外，您还可以直接下载预编译的 `exif-archive.exe` 文件，并直接使用它进行媒体归档。**

请注意，该工具按原样提供，根据您的系统设置可能需要额外的依赖项或配置。