# -*- coding: utf-8 -*-
from fastapi import FastAPI
from dotenv import load_dotenv, find_dotenv
from app.services import hello
from app.web import uvicorn
from app.log import log
from fastapi import Query
load_dotenv(find_dotenv(), override=True)
logger = log.set_log()

app = FastAPI()


@app.get("/")
def read_root(name: str = Query(...)):
    request = hello.dao.HelloRequest(name=name)
    return hello.do.get_ip(request)


if __name__ == "__main__":
    uvicorn.run_server()
