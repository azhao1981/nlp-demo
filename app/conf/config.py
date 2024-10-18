
# -*- coding: utf-8 -*-
import os
from pathlib import Path
from dynaconf import Dynaconf
from pydantic import BaseModel
from typing import Union, List
from dotenv import load_dotenv, find_dotenv

# 如果在 .env 设置 app_env = "production", 需要手动加载
load_dotenv(find_dotenv(), override=True)

class Setting(BaseModel):
    envvar_prefix: str = "APP"
    settings_files: List[str] = [
        os.path.join(os.path.dirname(__file__), 'settings.yaml'), 
        os.path.join(os.path.dirname(__file__), '.secrets.yaml')
    ]
    base_dir: Union[str, Path] = Path(__file__).parent.parent.parent
    includes: list[str] = []
    merge_enabled: bool = True
    environments: bool = True # Enable multi-level configuration，eg: default, development, production
    load_dotenv: bool = True  # Enable load .env
    env_switcher:str ="APP_ENV" # 在 .env 中设置 APP_ENV = "production",得到ENV=production 并且production 覆盖 default

_settting = Setting()

# app.root + config / settings.yaml
_settting.includes.append(*[str(_settting.base_dir / "config" / file) for file in ["settings.yaml"]])

Settings = Dynaconf(**_settting.model_dump())

# 如果需要多个环境切换，手动修改
# Settings.setenv('production')
# 重新加载配置文件
# Settings.reload()
# 加载新的配置文件
# Settings.load_file("new_settings.yaml")

if __name__ == "__main__":
    print(Settings.to_dict())