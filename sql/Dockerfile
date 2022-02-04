FROM mysql:8.0

ENV MYSQL_DATABASE products

ENV MYSQL_ROOT_PASSWORD root-password

RUN apt-get update && apt-get install -y locales

# Locale
RUN sed -i -e \
  's/# bg_BG.UTF-8 UTF-8/bg_BG.UTF-8 UTF-8/' /etc/locale.gen \
   && locale-gen

ENV LANG bg_BG.UTF-8
ENV LANGUAGE bg_BG:bg
ENV LC_LANG bg_BG.UTF-8
ENV LC_ALL bg_BG.UTF-8

COPY initialSetup /docker-entrypoint-initdb.d/