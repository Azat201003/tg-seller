FROM python:3.15.0a8-alpine3.23

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .

RUN python manage.py makemigrations --noinput
RUN python manage.py migrate --noinput
RUN python manage.py collectstatic --noinput
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
