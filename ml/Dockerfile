FROM python:3.14.7-alpine3.24

WORKDIR /usr/src/ml

COPY . .
RUN pip install uv
RUN rm -rf .venv
RUN uv venv --python 3.14.7
RUN source .venv/bin/activate
RUN uv sync

CMD ["uv", "run", "fastapi", "run", "main.py", "--host", "0.0.0.0", "--port", "8000"]