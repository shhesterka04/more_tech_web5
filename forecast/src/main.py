from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from router import forecast_router



def get_application() -> FastAPI:
    application = FastAPI()

    origins = [
        "*"
    ]

    application.add_middleware(
        CORSMiddleware,
        allow_origins=origins,
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )

    application.include_router(forecast_router)
    return application


app = get_application()