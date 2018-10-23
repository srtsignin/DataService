FROM bash

RUN ls
COPY . .
RUN ls

CMD ["/service"]