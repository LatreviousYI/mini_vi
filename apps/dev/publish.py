"""
Author       : kerwin_lv 2509456238@qq.com
Date         : 2024-01-31 14:14:54
LastEditTime: 2026-01-15 10:55:27
"""

import os
import shutil
import pathlib
import subprocess

root_path = pathlib.Path(os.path.abspath(__file__)).parent
project_path = root_path.parent.parent

# 创建发布文件夹
publish_path = os.path.join(root_path, "publish")
if not os.path.exists(publish_path):
    os.makedirs(publish_path)
else:
    shutil.rmtree(publish_path)
    os.makedirs(publish_path)
# os.remove("main.zip")

# 需要复制
config_path = os.path.join(root_path, "config")
wwwroot_path = os.path.join(root_path, "www")
exe_path = os.path.join(root_path, "main")
# 需要创建
db_path = os.path.join(publish_path, "data")

publish_config_path = os.path.join(publish_path, "config")
publish_wwwroot_path = os.path.join(publish_path, "www")

result = subprocess.run(
    "set GOOS=linux&& set GOARCH=amd64&& go build -o main ./main.go",
    shell=True,
    cwd=root_path,
    capture_output=True,
    text=True,
)
print(result.stderr)
print("build")

shutil.copytree(config_path, publish_config_path)
shutil.copytree(wwwroot_path, publish_wwwroot_path)
shutil.move(exe_path, publish_path)
os.makedirs(db_path)

shutil.make_archive("main", "zip", publish_path)
print("发布成功")
