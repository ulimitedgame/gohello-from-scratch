gohello
==========

This is an example of putting a golang service into docker scratch image

It is almost the smallest possiable docker image

Prepare
----------
```
$ tar cv --files-from /dev/null | docker import - scratch
```

Docker Version < 17.05
----------
```
$ sh build.sh
```

Docker Version >= 17.05
----------
```
$ docker build -t fuxu/gohello -f Dockerfile.multi-stage .
```

Run
----------
```
$ docker run -it --rm -p 8080:8080 --name hello fuxu/gohello
```

Reference:
----------
- https://docs.docker.com/engine/userguide/eng-image/baseimages/#create-a-simple-parent-image-using-scratch
- https://medium.com/@adriaandejonge/simplify-the-smallest-possible-docker-image-62c0e0d342ef
- https://docs.docker.com/develop/develop-images/multistage-build/
