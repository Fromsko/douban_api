from contextlib import asynccontextmanager
from typing import AsyncGenerator

from fastapi import FastAPI
from douban_api.web.api.movies import load_movies_data

@asynccontextmanager
async def lifespan_setup(
    app: FastAPI,
) -> AsyncGenerator[None, None]:  # pragma: no cover
    """
    Actions to run on application startup.

    This function uses fastAPI app to store data
    in the state, such as db_engine.

    :param app: the fastAPI application.
    :return: function that actually performs actions.
    """

    await load_movies_data()
    app.middleware_stack = None
    app.middleware_stack = app.build_middleware_stack()

    yield
