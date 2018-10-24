FROM bash

RUN ls
COPY . .
RUN ls

ADD data-service.exe /

CMD ["/data-service.exe"]