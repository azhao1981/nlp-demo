#!/bin/bash
set -x
exec_str="uv venv "

venv_root="$(pwd)"
if [ -f .env ] && grep -q "^VENV_ROOT=" .env; then
    venv_root=$(grep "^VENV_ROOT=" .env | cut -d '=' -f2)
fi

# 获得当前目录当做项目名, 用于创建虚拟环境，在 exec_str 后面加上 ". + 项目名"
project_name=$(basename "$(pwd)")
VENV_PATH="$venv_root/.$project_name"
if [ -f .env ] && grep -q "^VENV_PATH=" .env; then
    VENV_PATH=$(grep "^VENV_PATH=" .env | cut -d '=' -f2)
fi

exec_str+="$VENV_PATH"

# 如果.env中有指定 python 目录，则使用指定的 python 目录 -p
if [ -f .env ] && grep -q "^PYTHON_PATH=" .env; then
    python_path=$(grep "^PYTHON_PATH=" .env | cut -d '=' -f2)
    exec_str+=" -p $python_path"
fi

# 执行命令
eval $exec_str

activate_path=$(realpath "${VENV_PATH/#\~/$HOME}")
source "$activate_path/bin/activate"

# 安装依赖
uv pip install -r uv.poetry.txt

echo "" >> .env
echo "PATH=$VENV_PATH/bin:\$PATH" >> .env

echo "" >> .envrc
echo "source $VENV_PATH/bin/activate" >> .envrc
echo "source .env" >> .envrc

# 获取 Git 用户信息
GIT_NAME=$(git config user.name)
GIT_EMAIL=$(git config user.email)

# 获取当前目录名
DIR_NAME=$(basename "$PWD")

# 更新 pyproject.toml
sed -i "s/authors = .*/authors = [\"$GIT_NAME <$GIT_EMAIL>\"]/" pyproject.toml
sed -i "s/name = .*/name = \"$DIR_NAME\"/" pyproject.toml