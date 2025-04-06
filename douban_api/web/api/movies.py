import json
import random

import aiofiles
from fastapi import APIRouter, HTTPException

router = APIRouter()

movies_data = []


async def load_movies_data():
    global movies_data
    async with aiofiles.open("res/douban-top250.json", "r", encoding="utf-8") as afp:
        content = await afp.read()
        movies_data = json.loads(content)


@router.get("/movie/top/{movie_id}")
async def get_movie_by_id(movie_id: int):
    if movie_id < 0 or movie_id >= len(movies_data):
        raise HTTPException(status_code=404, detail="Movie not found")
    return movies_data[movie_id]


@router.get("/movie/all")
async def get_all_movies():
    return movies_data


@router.get("/movie/random")
async def get_random_movie():
    return random.choice(movies_data)


@router.get("/movie/random-image")
async def get_movie_image():
    movie = random.choice(movies_data)
    return movie["img"]
