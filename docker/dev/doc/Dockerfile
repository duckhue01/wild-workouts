FROM redocly/redoc

COPY ./main.yaml /usr/share/nginx/html/swagger.yaml

ENV SPEC_URL=swagger.yaml
ENV PORT=8080
ENV PAGE_TITLE="Tribe API documentation"
EXPOSE 8080
