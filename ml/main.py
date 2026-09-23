from fastapi import FastAPI
from pydantic import BaseModel
from typing import Optional

app = FastAPI()


class Claim(BaseModel):
    id: int
    content: str
    verdict: str
    created_at: str


@app.post("/claim")
def read_root(payload: Claim):
    print(payload)
    return {"payload": payload}