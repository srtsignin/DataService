FROM bash

RUN ls
COPY . .
RUN ls

ADD data-service /

CMD ["/data-service"]