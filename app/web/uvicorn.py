# -*- coding: utf-8 -*-
from pydantic import BaseModel
import uvicorn
from app.log import log
from app.conf import Settings  # 假设Settings在一个单独的配置模块中

logger = log.set_log()

class UvicornCnf(BaseModel):
    host: str
    port: int
    workers: int
    reload: bool = False

    @staticmethod
    def load_settings():
        return UvicornCnf(
            host=Settings.host,
            port=Settings.port,
            workers=Settings.workers,
            reload=Settings.ENV != 'production'  # .env 增加 APP_ENV=production|development
        )

def run_server():
    settings = UvicornCnf.load_settings()
    logger.info(f"UvicornCnf {settings.model_dump()}")
    uvicorn.run("main:app", **settings.model_dump())  
