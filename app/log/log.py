# -*- coding: utf-8 -*-
from pydantic import BaseModel, field_validator
import logging
from typing import Union
from app.conf.config import Settings
import inspect

level_mapping = {
            'debug': logging.DEBUG,
            'info': logging.INFO,
            'warning': logging.WARNING,
            'error': logging.ERROR,
            'critical': logging.CRITICAL
        }

class LoggerCfg(BaseModel):
    level: Union[str, int]
    format: str = '%(asctime)s - %(name)s - %(levelname)s - %(message)s'
    filename: str

    @field_validator('level')
    def check_level(cls, value):
        if isinstance(value, str):
            return level_mapping.get(value.lower(), logging.INFO)
        return value
_default_cfg = LoggerCfg(**Settings.log.to_dict())

def set_log(name = None):
    logging.basicConfig(**_default_cfg.model_dump())
    
    if not name:
        caller_frame = inspect.stack()[1]
        caller_module = inspect.getmodule(caller_frame[0])
        name = caller_module.__name__ if caller_module else 'root'

    return logging.getLogger(name)

def set_record(tag, file_name):
        logger = logging.getLogger(tag)
        handler = logging.FileHandler(file_name)
        formatter = logging.Formatter('%(message)s')
        handler.setFormatter(formatter)
        logger.addHandler(handler)
        logger.setLevel(logging.INFO)
        logger.propagate = False
        return logger

if __name__ == "__main__":
    logger = set_log()
    logger.debug("debug")
    logger.info("info")
    logger.warning("warning")
    logger.error("error")
    logger.critical("critical")