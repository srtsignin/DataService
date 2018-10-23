FROM bash

RUN ls
ADD service /
RUN ls
COPY . .
RUN ls

CMD ["/service"]