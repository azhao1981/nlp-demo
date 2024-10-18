# -*- coding: utf-8 -*-
import requests
from app.services.hello.dao import HelloRequest

def get_ip(request: HelloRequest) -> str:
    response = requests.get('https://ifconfig.me')
    return response.text + " " +request.name