from fastapi.routing import APIRouter

from douban_api.web.api import docs, echo, monitoring, movies

api_router = APIRouter()
api_router.include_router(monitoring.router)
api_router.include_router(docs.router)
api_router.include_router(movies.router, tags=["movies"])
api_router.include_router(echo.router, prefix="/echo", tags=["echo"])
