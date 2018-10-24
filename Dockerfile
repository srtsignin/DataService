FROM bash

RUN ls

ADD data-service.exe /

RUN ls

CMD ["/data-service.exe"]