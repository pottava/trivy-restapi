# Trivy v0.1
#   with initialized vulnerability database
#
# docker run -d --rm -p 9000:9000 pottava/trivy:0.1-db

FROM pottava/trivy:0.1
RUN trivy --debug --auto-refresh --clear-cache alpine:3.9
