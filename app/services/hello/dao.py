# -*- coding: utf-8 -*-
from pydantic import BaseModel

class HelloRequest(BaseModel):
    name: str 